package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"testing"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "myuser"
	Password = "mypass"
	Dbname   = "form"
)

func TestLeaveForm(t *testing.T) {
	// Create a new database connection.
	db, err := sql.Open("postgres", "host=localhost port=5432 user=myuser password=mypass dbname=form")
	if err != nil {
		t.Errorf("got %q, wanted %q", "nil", err)
	}
	defer db.Close()

	// Create a new leave form.
	leaveForm := LeaveForm{
		Name:      "John Doe",
		LeaveType: "Sick Leave",
		StartDate: "2023-01-01",
		EndDate:   "2023-01-05",
		Teams:     "Engineering",
	}

	// Save the leave form to the database.
	_, err = db.Exec("INSERT INTO form1.leavestype2 (name, leave_type, start_date, end_date, teams) VALUES ($1, $2, $3, $4, $5)", leaveForm.Name, leaveForm.LeaveType, leaveForm.StartDate, leaveForm.EndDate, leaveForm.Teams)
	if err != nil {
		t.Errorf("Error saving leave form: %v", err)
	}
}

// func main() {
// 	// Do something with the main package
// }

type LeaveForm struct {
	Name      string `json:"name"`
	LeaveType string `json:"leave_type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Teams     string `json:"teams"`
}
