package model

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/request"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Tweet struct {
	ID           int `gorm:"primarykey"`
	User_id      int
	Content      string
	Parent_tweet int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Tweets []*Tweet

func (appModel *AppModel) CreateTweet(request *request.CreateTweetRequest) (*Tweet, *app.Error) {
	tweet := Tweet{
		User_id:      request.User_id,
		Content:      request.Content,
		Parent_tweet: request.Parent_tweet,
	}
	result := appModel.DB.Create(&tweet)

	if result.Error != nil {
		me, ok := result.Error.(*mysql.MySQLError)
		if !ok {
			return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
		}
		fmt.Println(me)
		return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
	}
	return &tweet, nil
}

func (appModel *AppModel) GetTweets(request *request.GetTweetsRequest) (*Tweets, *app.Error) {
	var tweets Tweets
	var where *gorm.DB = appModel.DB
	var page, pageSize int

	if request.Page == 0 {
		page = 1
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

// GetTweetById
func (appModel *AppModel) GetTweetById(request *request.GetTweetsRequest, id string) (*Tweet, *app.Error) {
	var tweet Tweet
	var where *gorm.DB = appModel.DB
	int_id, err := strconv.ParseInt(id, 6, 12)
	if err != nil {
		fmt.Println(err)
	}
	result := where.
		Find(&tweet, int_id)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &tweet, nil
}

// UpdateTweet
func (appModel *AppModel) UpdateTweet(request *request.CreateTweetRequest, id string) (*Tweet, *app.Error) {
	tweet := Tweet{
		User_id:      request.User_id,
		Content:      request.Content,
		Parent_tweet: request.Parent_tweet,
	}
	int_id, err := strconv.ParseInt(id, 6, 12)
	if err != nil {
		fmt.Println(err)
	}
	result := appModel.DB.Where("id = ?", int_id).Updates(&tweet)

	if result.Error != nil {
		me, ok := result.Error.(*mysql.MySQLError)
		if !ok {
			return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
		}
		fmt.Println(me)
		return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
	}
	return &tweet, nil
}

func (appModel *AppModel) DeleteTweetByID(request *request.CreateTweetRequest, id string) (*Tweet, *app.Error) {
	var tweet Tweet
	int_id, err := strconv.ParseInt(id, 6, 12)
	if err != nil {
		fmt.Println(err)
	}
	result := appModel.DB.Delete(&Tweet{}, int_id)

	if result.Error != nil {
		me, ok := result.Error.(*mysql.MySQLError)
		if !ok {
			return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
		}
		fmt.Println(me)
		return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
	}
	return &tweet, nil
}
