package main

import (
	"html/template"
	"net/http"

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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parse HTML template file
		tmpl := template.Must(template.ParseFiles("page.html"))
		// Fetch users from the database
		users := fetch()
		// Create DataStructure instance with fetched users
		data := DataStructure{
			Info: users,
		}
		// Execute the template with the data and write to response writer
		tmpl.Execute(w, data)
	})
	// Start the HTTP server
	http.ListenAndServe(":5000", nil)
}
