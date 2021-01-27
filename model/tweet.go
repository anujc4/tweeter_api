package model

import (
	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/request"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Tweet struct {
	ID        int `gorm:"primarykey"`
	UserId  int    //TODO:Add association to user table
	Content  string
	ParentTweet  int //TODO:Add association
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tweets []*Tweet

func (appModel *AppModel) CreateTweet(request *request.CreateTweetRequest)  (*Tweet, *app.Error){
	tweet := Tweet{
		UserId: request.UserId,
		Content: request.Content,
		ParentTweet: request.ParentTweet,
	}
	result := appModel.DB.Create(&tweet)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
	}
	return &tweet, nil
}

func (appModel *AppModel) GetTweets(request *request.GetTweetsRequest)  (*Tweets, *app.Error){
	var tweets Tweets
	var tweetModel *gorm.DB = appModel.DB
	var page, pageSize int

	if request.ID != 0 {
		tweetModel = appModel.DB.Where("id=?", request.ID)
	}else if request.UserID != 0 {
		tweetModel = appModel.DB.Where("user_id=?", request.UserID)
	}

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

	result := tweetModel.
		Offset(offset).
		Limit(pageSize).
		Find(&tweets)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &tweets, nil
}

func (appModel AppModel) GetTweetById(id int) (*Tweet, *app.Error) {
	if id != 0{
		var tweet Tweet
		var tweetModel *gorm.DB = appModel.DB
		tweetModel = tweetModel.Where("id=?", id)
		result := tweetModel.Find(&tweet)
		if result.Error != nil{
			return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
		}
		return &tweet, nil
	}
	return nil, nil
}

func (appModel AppModel) UpdateTweet(request *request.UpdateTweetRequest, id int)  (*Tweet, *app.Error){
	if id != 0{
		var tweet Tweet
		var tweetModel *gorm.DB = appModel.DB
		result := tweetModel.Where("id=?", id).UpdateColumns(Tweet{Content: request.Content})
		count := result.RowsAffected
		if result.Error != nil || count == 0{
			return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
		}
		return &tweet, nil
	}
	return nil, nil
}

func (appModel AppModel) DeleteTweet(id int) (*Tweet, *app.Error){
	if id != 0{
		tweet, err := appModel.GetTweetById(id)
		if err != nil{
			return nil, app.NewError(err).SetCode(http.StatusBadRequest)
		}
		result := appModel.DB.Where("id=?", id).Delete(&tweet)
		count := result.RowsAffected
		if result.Error != nil || count==0{
			println("Error Count", count)
			return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
		}
		return tweet, nil
	}
	return nil, nil
}
