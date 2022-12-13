package service

import (
	db "gocard/db"
	"log"

	"github.com/gin-gonic/gin"
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

// TODO: replace db.Queries.GetUser with db.Queries.GetUserById

// func GetUser(c *gin.Context) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println(r)
// 			c.JSON(400, gin.H{
// 				"result": "UUID format error",
// 			})
// 		}
// 	}()
// 	id := uuid.MustParse(c.Param("id"))

// 	rows, err := db.Queries.GetUser(c, id)

// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"result": "something went wrong...",
// 		})
// 	} else {
// 		c.JSON(200, gin.H{
// 			"result": rows,
// 		})
// 	}
// }
