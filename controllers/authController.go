package controller

import (
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

type loginReqBody struct {
	Email    string
	Password string
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
		"message": result,
	})
}

// AuthHandler godoc
// @Summary			verify user information and issue token
// @Schemes
// @Description	verify user information and issue token
// @Tags				system
// @Accept			json
// @Produce			json
// @Param       request body loginReqBody true "loginReqBody"
// @Success	  	200			{string}	json		"{"result": "JWT token"}"
// @Router			/login [post]
func AuthHandler(c *gin.Context) {
	var (
		userInfo loginReqBody
		err      error
	)
	if err = c.BindJSON(&userInfo); err != nil {
		log.Println("BindJSON error: ", err.Error())
		c.JSON(400, nil)
		return
	}

	email := strings.Trim(userInfo.Email, " ")
	password := strings.Trim(userInfo.Password, " ")
	if email == "" || password == "" {
		c.JSON(400, gin.H{
			"message": "field can't be empty",
		})
		return
	}

	token, err := service.AuthHandler(email, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"data": token,
	})
}
