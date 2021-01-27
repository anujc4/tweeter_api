package env

import (
	"database/sql"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Env struct {
	DB *gorm.DB
}

func Init() *Env {
	// conf := config.Initialize()

	// connStr := fmt.Sprintf("%s:%s@/%s?parseTime=true", conf.MySql.Username, conf.MySql.Password, conf.MySql.Database)
	dsn := "root:sid@tcp(127.0.0.1:3307)/tweeter_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
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
