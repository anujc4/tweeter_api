package model

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/request"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Tweet represents tweet object in our app
type Tweet struct {
	ID          uint `gorm:"primarykey"`
	Content     string
	ParentTweet uint `gorm:"default:null"`
	UserID      uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//CreateTweet creates a new tweet
func (appModel *AppModel) CreateTweet(request *request.CreateTweetRequest) (*Tweet, *app.Error) {
	tweet := Tweet{
		UserID:  request.UserID,
		Content: request.Content,
	}
	result := appModel.DB.Create(&tweet)

	if result.Error != nil {
		_, ok := result.Error.(*mysql.MySQLError)
		if !ok {
			return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
		}
		return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
	}

	// Hack to get this bloody thing working
	// Default vlaue of uint in parent_tweet was giving referencial error cause
	// tweet with 0 id does not exist
	var unmarshledData map[string]interface{}
	marshledData, _ := json.Marshal(request)
	json.Unmarshal(marshledData, &unmarshledData)

	if _, exist := unmarshledData["parent_tweet"]; exist {
		tweet.ParentTweet = request.ParentTweet
		appModel.DB.Save(&tweet)
	}

	return &tweet, nil
}

//Tweets is an alias to an array of tweet
type Tweets []*Tweet

// GetTweets return all tweets
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

// GetTweetByID returns a user given a userID
func (appModel AppModel) GetTweetByID(tweetID *string) (*Tweet, *app.Error) {
	var tweet Tweet
	tx := appModel.DB.First(&tweet, *tweetID)
	if tx.Error != nil {
		return nil, app.NewError(tx.Error).SetCode(http.StatusNotFound)
	}
	return &tweet, nil
}

// DeleteTweet deletes tweet form database given an ID
func (appModel AppModel) DeleteTweet(tweetID *string) *app.Error {
	tx := appModel.DB.Delete(&Tweet{}, *tweetID)
	if tx.RowsAffected == 0 {
		return app.NewError(errors.New("Object does not exist")).SetCode(404)
	}
	return nil
}
