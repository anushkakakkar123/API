package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypass"
	dbname   = "form"
)

type Employee struct {
	Name string `json:"name"`
}

func GetAllEmployees(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		rows, err := db.Query("SELECT name FROM form1.leavestype2")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer rows.Close()

		var names []string

		for rows.Next() {
			var name string

			if err := rows.Scan(&name); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			names = append(names, name)
		}

		c.JSON(http.StatusOK, names)
	}
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	router := gin.Default()
	router.GET("/employees", GetAllEmployees(db))
	fmt.Println("Successfully connected!")
	router.Run(":8087")
}
