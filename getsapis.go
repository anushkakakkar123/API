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
	Name      string `json:"name"`
	LeaveType string `json:"leave_type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Teams     string `json:"teams"`
}

func GetEmployeeByName(db *sql.DB, name string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var employee Employee

		row := db.QueryRow("SELECT name, leave_type, start_date, end_date, teams FROM form1.leavestype2 WHERE name = $1", name)
		if err := row.Scan(&employee.Name, &employee.LeaveType, &employee.StartDate, &employee.EndDate, &employee.Teams); err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, employee)
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
	router.GET("/employees/:name", GetEmployeeByName(db, "Kanta"))
	fmt.Println("Successfully connected!")
	router.Run(":8087")
}
