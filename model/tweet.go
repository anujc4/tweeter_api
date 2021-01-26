package model

import (
	"net/http"
	"time"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/request"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Tweet struct {
	ID          uint `gorm:"primarykey"`
	UserId      uint
	Content     string
	ParentTweet uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tweets []*Tweet

func (appModel *AppModel) CreateTweet(request *request.CreateTweetRequest) (*Tweet, *app.Error) {
	tweet := Tweet{
		UserId:      request.UserId,
		Content:     request.Content,
		ParentTweet: request.ParentTweet,
	}

	result := appModel.DB.Create(&tweet)

	if result.Error != nil {
		_, ok := result.Error.(*mysql.MySQLError)
		if !ok {
			return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
		}
		return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
	}
	return &tweet, nil
}

func (appModel *AppModel) GetTweets(request *request.GetTweetsRequest) (*Tweets, *app.Error) {
	var tweets Tweets
	var where *gorm.DB = appModel.DB
	var page, pageSize int

	if request.UserId != "" {
		where = appModel.DB.Where("user_id = ?", request.UserId)
	} else if request.Content != "" {
		where = appModel.DB.Where("content LIKE ?", "%"+request.Content+"%")
	}
	page = 1
	pageSize = 10

	if request.Page != 0 {
		page = request.Page
	}

	if request.PageSize != 0 {
		pageSize = request.PageSize
	}

	switch {
	case request.PageSize > 100:
		pageSize = 100
	case request.PageSize <= 0:
		pageSize = 10
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

func (appModel *AppModel) GetTweetById(id int) (*Tweet, *app.Error) {
	var tweet Tweet
	var where *gorm.DB = appModel.DB

	where = appModel.DB.Where("ID = ?", id)

	result := where.Find(&tweet)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &tweet, nil
}

func (appModel *AppModel) UpdateTweetById(request *request.CreateTweetRequest, id int) (*Tweet, *app.Error) {
	var tweet Tweet
	var where *gorm.DB = appModel.DB

	result := where.First(&tweet, id)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	if request.Content != "" {
		tweet.Content = request.Content
	}

	result = where.Save(&tweet)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &tweet, nil
}

func (appModel *AppModel) DeleteTweetById(id int) *app.Error {

	result := appModel.DB.Delete(&Tweet{}, id)

	if result.Error != nil {
		return app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return nil
}
