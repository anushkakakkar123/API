package main

import (
	"database/sql"
	"errors"
	"testing"

	_ "github.com/lib/pq"
)

type Employee struct {
	Name      string `json:"name"`
	LeaveType string `json:"leave_type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Teams     string `json:"teams"`
}

func TestGetAllEmployees(t *testing.T) {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=myuser password=mypass dbname=form")
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db.Close()

	//employees, err := GetAllEmployees(db)

	if err != nil && err != errors.New("Database connection is nil") {
		t.Errorf("Expected error to be not nil, got nil")
	}
}

func GetAllEmployees(db *sql.DB) ([]Employee, error) {

	if db == nil {
		return nil, errors.New("Database connection is nil")
	}

	rows, err := db.Query("SELECT name, leave_type, start_date, end_date, teams FROM form1.leavestype2 LIMIT 2")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees []Employee

	for rows.Next() {
		var employee Employee

		if err := rows.Scan(&employee.Name, &employee.LeaveType, &employee.StartDate, &employee.EndDate, &employee.Teams); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
