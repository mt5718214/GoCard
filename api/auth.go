package api

import (
	"errors"
	"fmt"
	db "gocard/db"
	"gocard/util"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

var (
	// Secret key
	mySigningKey = []byte("mySigningKey")
)

type userInfoReqBody struct {
	Username, Password, CheckPassword string
}

type userInfo struct {
	Id                 int
	Username, Password string
}

func RegisterHandler(c *gin.Context) {
	var userInfoReqBody userInfoReqBody
	err := c.BindJSON(&userInfoReqBody)
	if err != nil {
		fmt.Println("BindJSON error: ", err.Error())
		c.JSON(400, nil)
		return
	}

	username := strings.Trim(userInfoReqBody.Username, " ")
	password := strings.Trim(userInfoReqBody.Password, " ")
	checkPassword := strings.Trim(userInfoReqBody.CheckPassword, " ")

	if username == "" || password == "" || checkPassword == "" {
		c.JSON(400, gin.H{
			"message": "field can't be empty",
		})
		return
	}

	if password != checkPassword {
		c.JSON(400, gin.H{
			"message": "password and checkPassword is not equal",
		})
		return
	}

	row := db.SqlDB.QueryRow("SELECT username FROM users WHERE username = ?", username)
	var userInfo userInfo
	err = row.Scan(&userInfo.Username)
	if err != nil {
		// hash password before insert
		hashString := util.HashPassword(password)

		result, insertErr := db.SqlDB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashString)
		if insertErr != nil {
			fmt.Println("Data insert error: ", insertErr.Error())
			c.JSON(400, gin.H{
				"messages": "Data insert error",
			})
			return
		}

		if affect, _ := result.RowsAffected(); affect != 1 {
			fmt.Println("Data insert failed")
			c.JSON(400, gin.H{
				"messages": "Data insert failed",
			})
			return
		}
		c.JSON(200, gin.H{
			"messages": "Create user success",
		})
		return
	}

	c.JSON(400, gin.H{
		"message": "Username already exist",
	})
}

func AuthHandler(c *gin.Context) {
	var userInfo userInfo
	err := c.BindJSON(&userInfo)
	if err != nil {
		fmt.Println("BindJSON error: ", err.Error())
		c.JSON(400, nil)
		return
	}

	username := strings.Trim(userInfo.Username, " ")
	password := strings.Trim(userInfo.Password, " ")
	if username == "" || password == "" {
		c.JSON(400, gin.H{
			"message": "field can't be empty",
		})
		return
	}

	row := db.SqlDB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	err = row.Scan(&userInfo.Id, &userInfo.Username, &userInfo.Password)
	if err != nil {
		fmt.Println("QueryRow error: ", err.Error())
		c.JSON(400, gin.H{
			"message": "Username or password is wrong.",
		})
		return
	}

	// check password hash string
	isMatch := util.CheckPasswordHash(userInfo.Password, password)
	if !isMatch {
		c.JSON(400, gin.H{
			"message": "Username or password is wrong.",
		})
		return
	}

	// sign JWT token and return to client
	token, err := createJWT("token", userInfo.Id, userInfo.Username)
	if err != nil {
		fmt.Println("createJWT error: ", err.Error())
		c.JSON(400, gin.H{
			"token": "",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get token from Header.Authorization field.
		// same as c.Request.Header.Get("Authorization")
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization is null in Header",
			})
			c.Abort()
			return
		}

		part := strings.SplitN(authHeader, " ", 2)
		if !(len(part) == 2 && part[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Format of Authorization is wrong",
			})
			c.Abort()
			return
		}

		cm, err := ParseToken(part[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Token",
			})
			c.Abort()
			return
		}

		// Store Account info into Context
		// After that, we can get userId from c.Get("userId")
		c.Set("userId", cm["jti"])

		c.Next()
	}
}

// 如果需要包含客製化使用者資訊的Claim可使用以下的struct
// type JWTClaim struct {
// 	*jwt.RegisteredClaims
// 	UserInfo interface{}
// }

func createJWT(sub string, userId int, username string) (string, error) {
	// Create the Claims
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	claim := jwt.RegisteredClaims{
		Issuer:    "kemp",
		Subject:   sub,
		Audience:  []string{username},
		ExpiresAt: jwt.NewNumericDate(expireTime),
		IssuedAt:  jwt.NewNumericDate(nowTime),
		ID:        strconv.Itoa(userId),
	}

	// token instance
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("generate token fail: ", err.Error())
		return "", err
	}

	return ss, err
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	if strings.Trim(tokenString, " ") == "" {
		return nil, errors.New("invalid token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		/**
		* Comma-ok 斷言
		* 可以直接判斷是否是該型別的變數： value, ok = element.(T)
		* value 就是變數的值，ok 是一個 bool 型別，element 是 interface 變數，T 是斷言的型別。
		* 如果 element 裡面確實儲存了 T 型別的數值，那麼 ok 回傳 true，否則回傳 false。
		 */
		// 驗證 alg 是否為預期的HMAC演算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println("parse token error: ", err.Error())
		return nil, err
	}

	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("claim", claim)
		return claim, nil
	}

	return nil, errors.New("invalid token")
}
