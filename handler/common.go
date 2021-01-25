package handler

import (
	"crypto/rsa"

	"github.com/anujc4/tweeter_api/internal/env"
	"gorm.io/gorm"
)

type HttpApp struct {
	DB         *gorm.DB
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewHttpApp(env *env.Env) *HttpApp {
	return &HttpApp{
		DB:         env.DB,
		PrivateKey: env.PrivateKey,
	}
}
