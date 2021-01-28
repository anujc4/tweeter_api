package model

import (
	"net/http"
	"time"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/request"
	"gorm.io/gorm"
)

type Tweet struct {
	ID          uint `gorm:"primarykey"`
	UserID      uint
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User `gorm:"foreignKey:UserID"`
	ParentTweet uint
}

type Tweets []*Tweet

func (appModel *AppModel) CreateTweet(request *request.CreateTweetRequest) (*Tweet, *app.Error) {
	tweet := Tweet{
		UserID:      request.UserID,
		ParentTweet: request.ParentTweet,
		Content:     request.Content,
	}
	result := appModel.DB.Create(&tweet)
	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
	}
	return &tweet, nil
}

func (appModel *AppModel) GetTweets(request *request.GetTweetsRequest) (*Tweets, *app.Error) {
	var tweets Tweets
	var where *gorm.DB = appModel.DB
	var page, pageSize int

	if request.ParentTweet != "" {
		where = appModel.DB.Where("parent_tweet = ?", request.ParentTweet)
	} else if request.Content != "" {
		where = appModel.DB.Where("content LIKE ?", "%"+request.Content+"%")
	} else if request.UserID != "" {
		where = appModel.DB.Where("user_id = ?", request.UserID)
	}

	if request.Page == 0 {
		page = 1
	} else {
		page = request.Page
	}

	if request.PageSize <= 0 {
		pageSize = 10
	} else {
		pageSize = request.PageSize
	}

	offset := (page - 1) * pageSize

	result := where.
		Offset(offset).
		Limit(pageSize).
		Find(&tweets)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &tweets, nil
}
