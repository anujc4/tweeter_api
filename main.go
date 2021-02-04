package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

func createUser(firstName, lastName, email, password string) (int64, error) {
	query := "INSERT INTO users (first_name, last_name, email, password) VALUES(?, ?, ?, ?);"
	statement, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	resp, err := statement.Exec(firstName, lastName, email, password)
	if err != nil {
		return 0, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func deleteUser(id int64) (int64, error) {
	query := "DELETE FROM users WHERE id = ?;"
	statement, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	resp, err := statement.Exec(id)
	if err != nil {
		return 0, err
	}

	id, err = resp.RowsAffected()
	if err != nil {
		return 0, err
	}
	return id, err
}

func updateUserName(id int64, firstName, lastName string) (int64, error) {
	query := "UPDATE users SET first_name = ?, last_name = ? WHERE id = ?;"
	statement, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	resp, err := statement.Exec(firstName, lastName, id)
	if err != nil {
		return 0, err
	}

	id, err = resp.RowsAffected()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func getUserByID(id int64) (User, error) {
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
		return user, err
	}
	return user, nil
}

func getAllUsers() ([]User, error) {
	var users []User
	query := "SELECT id, first_name, last_name, email, created_at, updated_at FROM users;"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func healthCheck(c *gin.Context) {
	c.String(200, "UP")
}

func create(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := createUser(user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(200, fmt.Sprint(id))
}

func read(c *gin.Context) {
	users, err := getAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rows, err := updateUserName(int64(id), user.FirstName, user.LastName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rows > 0 {
		user, err := getUserByID(int64(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}
}

func delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rows, err := deleteUser(int64(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rows > 0 {
		c.Status(http.StatusNoContent)
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Unable to find user with id %d", id)})
		return
	}
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

	r := gin.Default()
	r.GET("/simple_health", healthCheck)

	v1 := r.Group("/v1")
	v1.POST("/user", create)
	v1.GET("/users", read)
	v1.PUT("/user/:id", update)
	v1.DELETE("/user/:id", delete)

	r.Run(":3000")
}
