package model

import (
	"gorm.io/gorm"
)

// Tweet represents tweet object in our app
type Tweet struct {
	gorm.Model  //gorm.Model includes fields like ID, CreatedAt and UpdatedAt
	Content     string
	ParentTweet uint
	UserID      uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
