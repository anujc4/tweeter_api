package model

import (
	"context"

	"gorm.io/gorm"
)

type AppModel struct {
	DB  *gorm.DB
	ctx context.Context
}

func NewAppModel(ctx context.Context, db *gorm.DB) *AppModel {
	return &AppModel{
		ctx: ctx,
		DB:  db,
	}
}
