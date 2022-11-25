package api

import (
	"fmt"
	db "gocard/db"
	"log"

	"github.com/gin-gonic/gin"
)

type PostTodoRequestBody struct {
	Title   string
	Content string
}

func GetTodoLists(c *gin.Context) {
	query := "SELECT title, content FROM todo"
	rows, err := db.SqlDB.QueryContext(c, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	todos := make([]PostTodoRequestBody, 0)

	for rows.Next() {
		var todo PostTodoRequestBody
		if err := rows.Scan(&todo.Title, &todo.Content); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, todos)
}

func GetTodoList(c *gin.Context) {
	id := c.Param("id")
	query := "SELECT title, content FROM todo WHERE id = ?"
	row := db.SqlDB.QueryRow(query, id)

	var todo PostTodoRequestBody
	err := row.Scan(&todo.Title, &todo.Content)
	if err != nil {
		fmt.Println("get todo error", err.Error())
	}

	c.JSON(200, todo)
}

func PostTodo(c *gin.Context) {
	var requestBody PostTodoRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		log.Fatal("error", err)
	}

	query := "INSERT INTO todo (title, content, is_complete) VALUES (?, ?, 0)"
	result, err := db.SqlDB.Exec(query, requestBody.Title, requestBody.Content)
	if err != nil {
		log.Fatal("insert todo err", err.Error())
	}
	if row, _ := result.RowsAffected(); row != 1 {
		log.Fatal("insert todo data unmatched.")
	}

	c.JSON(201, nil)
}

func PutTodo(c *gin.Context) {
	id := c.Param("id")
	var requestBody PostTodoRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		log.Fatal("error", err)
	}

	// TODO: 修改更update邏輯, `title or content無值時會是空白字串`
	query := "UPDATE todo SET title = IFNULL(?, title), content = IFNULL(?, content) WHERE id = ?"
	_, err := db.SqlDB.Exec(query, requestBody.Title, requestBody.Content, id)
	if err != nil {
		fmt.Println("update todo err:", err.Error())
	}

	c.JSON(204, nil)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	query := "DELETE FROM todo WHERE id = ?"
	_, err := db.SqlDB.ExecContext(c, query, id)
	if err != nil {
		log.Fatal(err)
	}
}
