package model

import (
	"net/http"
	"time"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/request"
)

type Tweet struct {
	ID          uint `gorm:"primarykey"`
	UserID      uint
	Content     string
	ParentTweet uint
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tweets []*Tweet

func (appModel *AppModel) CreateTweet(request *request.CreateTweetRequest) (*Tweet, *app.Error) {
	tweet := Tweet{
		UserID:      request.UserID,
		Content:     request.Content,
		ParentTweet: request.ParentTweet,
	}
	result := appModel.DB.Create(&tweet)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
	}
	return &tweet, nil
}
