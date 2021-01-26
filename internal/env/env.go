package env

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/anujc4/tweeter_api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Env struct {
	DB *gorm.DB
}

func Init() *Env {
	conf := config.Initialize()

	connStr := fmt.Sprintf("%s:%s@/%s?parseTime=true", conf.MySql.Username, conf.MySql.Password, conf.MySql.Database)
	db, err := sql.Open("mysql", connStr)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	return &Env{
		DB: gormDB,
	}
}
