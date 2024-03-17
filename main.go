package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User struct represents a user model
type User struct {
	gorm.Model // Embedding gorm.Model will automatically add ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name       string
}

// DataStructure is a struct to hold data for the HTML template
type DataStructure struct {
	Info []User // Changed 'info' to 'Info' for proper JSON marshalling in template
}

// fetch function retrieves all users from the database
func fetch() []User {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&User{})
	var users []User
	db.Find(&users)

	return users
}

func main() {
	data := fetch()
	component := hello(data)
	http.Handle("/",templ.Handler(component))
	// Start the HTTP server
	fmt.Println("Listening on :5000")
	http.ListenAndServe(":5000", nil)
}
