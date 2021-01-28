package model

import (
	"context"

	"gorm.io/gorm"
)

//AppModel in model common.go
type AppModel struct {
	DB  *gorm.DB
	ctx context.Context
}

//NewAppModel in model common.go
func NewAppModel(ctx context.Context, db *gorm.DB) *AppModel {
	return &AppModel{
		ctx: ctx,
		DB:  db,
	}
}
