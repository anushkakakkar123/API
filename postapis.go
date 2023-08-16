package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypass"
	dbname   = "form"
)

func leave_form(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the data from the form
		log.Println("checking1")
		name := c.PostForm("name")
		leave_type := c.PostForm("leave_type")
		start_date := c.PostForm("start_date")
		end_date := c.PostForm("end_date")
		teams := c.PostForm("teams")

		// Save the data to the database
		var err error

		_, err = db.Exec("INSERT INTO form1.leavestype2 (name, leave_type, start_date, end_date, teams) VALUES ($1, $2, $3, $4, $5)", name, leave_type, start_date, end_date, teams)
		if err != nil {
			log.Println(name, " ", leave_type, " ", start_date, " ", end_date, " ", teams)
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Success"})
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
	router.POST("/leave", leave_form(db))
	router.Run(":8086")

}
