package service

import (
	"context"
	"errors"
	"fmt"
	db "gocard/db"
	sqlc "gocard/db/sqlc"
	enum "gocard/enum"
	"gocard/util"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(username, password, email string) (string, error) {
	var err error
	arg := sqlc.PostUserParams{
		Name:          username,
		Password:      util.HashPassword(password),
		Email:         email,
		CreatedBy:     enum.Admin.AdminUuid(),
		LastUpdatedBy: enum.Admin.AdminUuid(),
	}

	if _, err = db.Queries.PostUser(context.Background(), arg); err != nil {
		log.Print("[Error] PostUser error: ", err.Error())
		if strings.Contains(err.Error(), "duplicate") {
			return "", errors.New("email is exist")
		}

		return "", errors.New("PostUser error")
	}

	return "Create user success", nil
}

func AuthHandler(email, password string) (string, error) {
	var err error
	userInfoFromDb, err := db.Queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		log.Println("[Error] GetUserByEmail error: ", err.Error())
		return "", errors.New("email or password is wrong")
	}

	// check password hash string
	if isMatch := util.CheckPasswordHash(userInfoFromDb.Password, password); !isMatch {
		return "", errors.New("email or password is wrong")
	}

	// sign JWT token and return to client
	token, err := util.CreateJWT("token", userInfoFromDb.ID, userInfoFromDb.Name, userInfoFromDb.Email, userInfoFromDb.IsAdmin)
	if err != nil {
		log.Println("createJWT error: ", err.Error())
		return "", errors.New("something went wrong")
	}

	return token, nil
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

		cm, err := util.ParseToken(part[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Token",
			})
			c.Abort()
			return
		}

		// Store Account info into Context
		// After that, we can get userInfo from c.GetStringMap("userInfo")
		c.Set("userInfo", cm["UserInfo"])

		c.Next()
	}
}

func AuthAdminMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		userInfo := c.GetStringMap("userInfo")
		fmt.Println(userInfo)
		isAdminFloat, ok := userInfo["IsAdmin"].(float64)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			c.Abort()
			return
		}

		if int(isAdminFloat) != 1 {
			c.JSON(http.StatusForbidden, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
