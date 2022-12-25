package util

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	// Secret key
	mySigningKey = []byte("mySigningKey")
)

type userInfo struct {
	ID      uuid.UUID
	Email   string
	IsAdmin int16
}

// 如果需要包含客製化使用者資訊的Claim可使用以下的struct
type JWTClaim struct {
	*jwt.RegisteredClaims
	UserInfo userInfo
}

func CreateJWT(sub string, userId uuid.UUID, username, email string, isAdmin int16) (string, error) {
	// Create the Claims
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	claim := JWTClaim{
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "kemp",
			Subject:   sub,
			Audience:  []string{username},
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			// ID:        uuid.UUID.String(userId),
		},
		UserInfo: userInfo{
			ID:      userId,
			Email:   email,
			IsAdmin: isAdmin,
		},
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
		return claim, nil
	}

	return nil, errors.New("invalid token")
}
