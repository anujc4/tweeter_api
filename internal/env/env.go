package env

import (
	"crypto/rsa"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/anujc4/tweeter_api/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Env struct {
	DB         *gorm.DB
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	Redis      *redis.Client
}

func Init() *Env {
	conf := config.Initialize()
	var env Env

	InitAuth(conf, &env)
	InitDB(conf, &env)
	InitRedis(conf, &env)

	return &env
}

func InitDB(conf *config.Configuration, env *Env) {
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
	env.DB = gormDB
}

func InitAuth(conf *config.Configuration, env *Env) {
	private_key_file := filepath.Join(conf.Auth.Path, conf.Auth.PrivateKeyFileName)
	public_key_file := filepath.Join(conf.Auth.Path, conf.Auth.PublicKeyFileName)

	// Parse Private Key
	pem_key, err := ioutil.ReadFile(private_key_file)
	if err != nil {
		log.Fatal("unable to open private key file", err)
	}
	var private_key *rsa.PrivateKey
	if private_key, err = jwt.ParseRSAPrivateKeyFromPEM(pem_key); err != nil {
		log.Fatal("unable to parse private key", err)
	}
	env.PrivateKey = private_key

	// Parse Public Key
	pub_key, err := ioutil.ReadFile(public_key_file)
	if err != nil {
		log.Fatal("unable to open public key file", err)
	}
	var public_key *rsa.PublicKey
	if public_key, err = jwt.ParseRSAPublicKeyFromPEM(pub_key); err != nil {
		log.Fatal("unable to parse public key", err)
	}
	env.PublicKey = public_key
}

func InitRedis(conf *config.Configuration, env *Env) {
	redisURL := fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
	env.Redis = rdb
}
