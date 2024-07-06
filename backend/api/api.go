package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type ListItem struct {
	Id   string `json:"id"`
	Item string `json:"item"`
	Done bool   `json:"done"`
}

var db *sql.DB
var err error

func SetupPostgres() {
	// db, err = sql.Open("postgres", "postgres://postgres:password@postgres/todo?sslmode=disable")

	// when running in docker
	db, err = sql.Open("postgres", "postgres://postgres:password@postgres/todo?sslmode=disable")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
		return
	}

	log.Println("connected to postgres")
}

// CRUD: Create Read Update Delete API Format

// List all todo items
func TodoItems(c *gin.Context) {
	// Use SELECT Query to obtain all rows
	rows, err := db.Query("SELECT * FROM list")
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
		return
	}
	defer rows.Close()

	// Get all rows and add into items
	items := make([]ListItem, 0)
	for rows.Next() {
		item := ListItem{}
		if err := rows.Scan(&item.Id, &item.Item, &item.Done); err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
			return
		}
		item.Item = strings.TrimSpace(item.Item)
		items = append(items, item)
	}

	// Return JSON object of all rows
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{"items": items})
}

// Create todo item and add to DB
func CreateTodoItem(c *gin.Context) {
	var TodoItem ListItem
	if err := c.BindJSON(&TodoItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	if len(TodoItem.Item) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter an item"})
		return
	}

	TodoItem.Done = false

	// Insert item to DB and return the inserted id
	var id int
	err := db.QueryRow("INSERT INTO list(item, done) VALUES($1, $2) RETURNING id;", TodoItem.Item, TodoItem.Done).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
		return
	}

	TodoItem.Id = fmt.Sprintf("%d", id)

	log.Println("created todo item", TodoItem.Item)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusCreated, gin.H{"items": &TodoItem})
}

// Update todo item
func UpdateTodoItem(c *gin.Context) {
	id := c.Param("id")
	done := c.Param("done")

	if len(id) == 0 || len(done) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter both id and done state"})
		return
	}

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM list WHERE id=$1);", id).Scan(&exists)
	if err != nil || !exists {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	_, err = db.Exec("UPDATE list SET done=$1 WHERE id=$2;", done == "true", id)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
		return
	}

	log.Println("updated todo item", id, done)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{"message": "successfully updated todo item", "todo": id})
}

// Delete todo item
func DeleteTodoItem(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter an id"})
		return
	}

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM list WHERE id=$1);", id).Scan(&exists)
	if err != nil || !exists {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	_, err = db.Exec("DELETE FROM list WHERE id=$1;", id)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
		return
	}

	log.Println("deleted todo item", id)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted todo item", "todo": id})
}

// Filter todo items by their done status
func FilterTodoItems(c *gin.Context) {
	done := c.Param("done")

	if done != "true" && done != "false" {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter a valid done state (true/false)"})
		return
	}

	rows, err := db.Query("SELECT * FROM list WHERE done=$1", done == "true")
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
		return
	}
	defer rows.Close()

	items := make([]ListItem, 0)
	for rows.Next() {
		item := ListItem{}
		if err := rows.Scan(&item.Id, &item.Item, &item.Done); err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
			return
		}
		item.Item = strings.TrimSpace(item.Item)
		items = append(items, item)
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{"items": items})
}
