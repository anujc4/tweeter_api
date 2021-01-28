package model

import (
	"net/http"
	"time"

	"github.com/anujc4/tweeter_api/internal/app"
)

type Tweet struct {
	ID          uint `gorm:"primarykey"`
	UserID      uint
	Content     string
	ParentTweet *uint `gorm:"foreignkey:ID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
type Tweets []Tweet

func (appModel *AppModel) CreateTweet(userID uint, content string, parentID *uint) (*Tweet, *app.Error) {
	tweet := Tweet{
		UserID:      userID,
		Content:     content,
		ParentTweet: parentID,
	}

	result := appModel.DB.Create(&tweet)
	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
	}
	return &tweet, nil
}
