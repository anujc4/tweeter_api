package model

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/request"
	"github.com/go-sql-driver/mysql"
)

type Tweet struct {
	ID           int `gorm:"primarykey"`
	user_id      int
	content      string
	parent_tweet int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Tweets []*Tweet

func (appModel *AppModel) CreateTweet(request *request.CreateTweetRequest) (*Tweet, *app.Error) {
	tweet := Tweet{
		user_id:      request.User_id,
		content:      request.Content,
		parent_tweet: request.Parent_tweet,
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
