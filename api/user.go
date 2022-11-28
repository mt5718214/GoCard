package api

import (
	"fmt"
	db "gocard/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListUsers(c *gin.Context) {
	rows, err := db.Queries.ListUsers(c)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"result": rows,
	})
}

func GetUser(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			c.JSON(400, gin.H{
				"result": "UUID format error",
			})
		}
	}()
	id := uuid.MustParse(c.Param("id"))

	rows, err := db.Queries.GetUser(c, id)

	if err != nil {
		c.JSON(400, gin.H{
			"result": "something went wrong...",
		})
	} else {
		c.JSON(200, gin.H{
			"result": rows,
		})
	}
}
