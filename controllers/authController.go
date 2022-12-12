package controller

import (
	sqlc "gocard/db/sqlc"
	"gocard/service"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type signupReqBody struct {
	Name          string
	Email         string
	Password      string
	CheckPassword string
}

// RegisterHandler godoc
// @Summary				 User register
// @Schemes
// @Description		 User register
// @Tags					 system
// @Accept				 json
// @Produce				 json
// @Param       	 request body signupReqBody true "signupReqBody"
// @Success	  		 201			{string}	json		"{"result":"Create user success"}"
// @Router				 /signup [post]
func RegisterHandler(c *gin.Context) {
	var userInfoReqBody signupReqBody
	err := c.BindJSON(&userInfoReqBody)
	if err != nil {
		log.Println("BindJSON error: ", err.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	username := strings.Trim(userInfoReqBody.Name, " ")
	password := strings.Trim(userInfoReqBody.Password, " ")
	email := strings.Trim(userInfoReqBody.Email, " ")
	checkPassword := strings.Trim(userInfoReqBody.CheckPassword, " ")

	if username == "" || password == "" || email == "" || checkPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "field can't be empty",
		})
		return
	}

	if password != checkPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password and checkPassword is not equal",
		})
		return
	}

	result, err := service.RegisterHandler(username, password, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"result": result,
	})
}

func AuthHandler(c *gin.Context) {
	var userInfo sqlc.User
	err := c.BindJSON(&userInfo)
	if err != nil {
		log.Println("BindJSON error: ", err.Error())
		c.JSON(400, nil)
		return
	}

	username := strings.Trim(userInfo.Name, " ")
	password := strings.Trim(userInfo.Password, " ")
	if username == "" || password == "" {
		c.JSON(400, gin.H{
			"message": "field can't be empty",
		})
		return
	}
	// service.AuthHandler(c, userInfo)
}
