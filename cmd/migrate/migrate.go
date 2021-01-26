package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"database/sql"

	"github.com/anujc4/tweeter_api/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4/database/mysql"
)

var out io.Writer = os.Stdout

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("migrate must be used with a migration command")
		fmt.Println("Usage: migrate down | drop | up | version | force number | step number | toVersion number [-configName string] [-configPath string]")
		os.Exit(1)
	}

	migrationPath := fmt.Sprintf("file://db/migration")
	conf := config.Initialize()
	connStr := fmt.Sprintf("%s:%s@/%s?parseTime=true", conf.MySql.Username, conf.MySql.Password, conf.MySql.Database)
	db, err := sql.Open("mysql", connStr)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationPath, "mysql", driver)
	if err != nil {
		log.Println(migrationError(args[0], err))
		os.Exit(1)
	}

	// Run the specified command.
	err = runMigration(m, args[0], args[1:])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("Success!")
}

func runMigration(m *migrate.Migrate, command string, args []string) error {
	fmt.Printf("Running migrate %s %s...\n", command, strings.Join(args, " "))

	var err error

	switch command {
	case "down":
		err = m.Down()
	case "drop":
		err = m.Drop()
	case "up":
		err = m.Up()
	case "version":
		version, dirty, err := m.Version()
		if err != nil {
			return migrationError(command, err)
		}
		fmt.Fprintf(out, "\tVersion: %d\n\tDirty: %t\n", version, dirty)
	default:
		return migrationError(command, errors.New("unknown command"))
	}

	if err != nil {
		return migrationError(command, err)
	}
	return nil
}

func migrationError(command string, err error) error {
	return fmt.Errorf("migrate %s: %v", command, err)
}
