package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func dbConn() {
	connStr := fmt.Sprintf("%s:%s@/%s?parseTime=true", "root", "", "tweeter_db")
	database, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db = database
}

// User type
type User struct {
	ID        int64     `json:"id,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func createUser(firstName, lastName, email, password string) int64 {
	return 0
}

func deleteUser(id int64) int64 {
	return 0
}

func updateUserName(id int64, firstName, lastName string) int64 {
	return 0
}

func getUserByID(id int64) User {
	return User{}
}

func getAllUsers() []User {
	return nil
}

func main() {
	dbConn()
	// Close the DB connection once the program completes execution
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.Panic("Unable to connect to db", err)
	}
	fmt.Print("Established contact with database")

	fmt.Println("sql demo application.")
	for true {
		fmt.Print("1. Insert\n2. Update\n3. Get by ID\n4. Delete User by ID.\n5. Print All.\n6. Exit.\nEnter your choice: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var firstName, lastName, email, password string
			fmt.Print("Insert first name: ")
			fmt.Scanln(&firstName)
			fmt.Print("Insert last name: ")
			fmt.Scanln(&lastName)
			fmt.Print("Insert email: ")
			fmt.Scanln(&email)
			fmt.Print("Insert password: ")
			fmt.Scanln(&password)
			id := createUser(firstName, lastName, email, password)
			fmt.Printf("User created with id: %d\n", id)
		case 2:
			var firstName, lastName string
			var id int64
			fmt.Print("Input the id of the user to update: ")
			fmt.Scanln(&id)
			fmt.Print("Insert first name: ")
			fmt.Scanln(&firstName)
			fmt.Print("Insert last name: ")
			fmt.Scanln(&lastName)
			rows := updateUserName(id, firstName, lastName)
			if rows > 0 {
				fmt.Println("User has been updated.")
			} else {
				fmt.Println("No user was updated.")
			}
		case 3:
			var id int64
			fmt.Print("Input the id of the user to find: ")
			fmt.Scanln(&id)
			user := getUserByID(id)
			enc, _ := json.Marshal(user)
			fmt.Println(string(enc))
		case 4:
			var id int64
			fmt.Print("Input the id of the user to delete: ")
			fmt.Scanln(&id)
			rows := deleteUser(id)
			if rows > 0 {
				fmt.Println("User has been deleted.")
			}
		case 5:
			users := getAllUsers()
			for _, u := range users {
				enc, _ := json.Marshal(u)
				fmt.Println(string(enc))
			}
		case 6:
			fmt.Println("Exiting....")
			os.Exit(0)
		default:
		}
	}
}
