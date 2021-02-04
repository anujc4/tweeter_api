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
	query := "INSERT INTO users (first_name, last_name, email, password) VALUES(?, ?, ?, ?);"
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := statement.Exec(firstName, lastName, email, password)
	if err != nil {
		log.Fatal(err)
	}

	id, err := resp.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func deleteUser(id int64) int64 {
	query := "DELETE FROM users WHERE id = ?;"
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := statement.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	id, err = resp.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func updateUserName(id int64, firstName, lastName string) int64 {
	query := "UPDATE users SET first_name = ?, last_name = ? WHERE id = ?;"
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := statement.Exec(firstName, lastName, id)
	if err != nil {
		log.Fatal(err)
	}

	id, err = resp.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func getUserByID(id int64) User {
	var user User
	query := "SELECT id, first_name, last_name, email, created_at, updated_at FROM users WHERE id = ?;"

	row := db.QueryRow(query, id)
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt); err != nil {
		log.Panic(err)
	}
	return user
}

func getAllUsers() []User {
	var users []User
	query := "SELECT id, first_name, last_name, email, created_at, updated_at FROM users;"

	rows, err := db.Query(query)
	if err != nil {
		log.Panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
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
