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
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LeaveType string `json:"leave_type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Teams     string `json:"teams"`
}

func GetAllEmployees(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		rows, err := db.Query("SELECT id, name, leave_type, start_date, end_date, teams FROM form1.leavestype3")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer rows.Close()

		var employees []Employee

		for rows.Next() {
			var employee Employee

			if err := rows.Scan(&employee.ID, &employee.Name, &employee.LeaveType, &employee.StartDate, &employee.EndDate, &employee.Teams); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			employees = append(employees, employee)
		}
		fmt.Println("employees", employees)

		c.JSON(http.StatusOK, employees)
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
