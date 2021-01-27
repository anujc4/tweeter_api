package handler

import "gorm.io/gorm"

type HttpApp struct {
	DB *gorm.DB
}

func NewHttpApp(db *gorm.DB) *HttpApp {
	return &HttpApp{
		DB: db,
	}
}
