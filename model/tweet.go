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
	ID        int `gorm:"primarykey"`
	UserId  int
	Content  string
	ParentTweet  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tweets []*Tweet

func (appModel *AppModel) CreateTweet(request *request.CreateTweetRequest) (*Tweet, *app.Error) {
	tweet := Tweet{
		UserId: request.UserId,
		Content:  request.Content,
		ParentTweet: request.ParentTweet,
	}
	result := appModel.DB.Create(&tweet)

	if result.Error != nil {
		me, ok := result.Error.(*mysql.MySQLError)
		if !ok {
			return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
		}
		if me.Number == 1062 {
			return nil, app.
				NewError(result.Error).
				SetMessage("already").
				SetCode(http.StatusBadRequest)
		}
		return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
	}
	return &tweet, nil
}

func (appModel *AppModel) GetTweets(request *request.GetTweetsRequest) (*Tweets, *app.Error) {
	var tweets Tweets
	var where *gorm.DB = appModel.DB
	var page, pageSize int

	if request.ID != 0 {
		where = appModel.DB.Where("ID = ?", request.ID)
	} else if request.ParentTweet != 0 {
		where = appModel.DB.Where("parent_tweet = ?", request.ParentTweet)
	}else if request.UserId != 0 {
		where = appModel.DB.Where("user_id = ?",request.UserId)
	}

	if request.Page == 0 {
		page = 1
	} else {
		page = request.Page
	}

	switch {
	case request.PageSize > 100:
		pageSize = 100
	case request.PageSize <= 0:
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



func(appModel *AppModel) GetTweetByID(id int)(*Tweet, *app.Error)  {
	var tweet Tweet
	var where *gorm.DB = appModel.DB
	if id != 0 {
		where = appModel.DB.Where("id = ?", id)
	}

  result := where.Find(&tweet)
  if tweet.Content == ""  && tweet.UserId == 0 && tweet.ParentTweet == 0{
    return nil, app.
      NewError(result.Error).
      SetMessage("not found").
      SetCode(http.StatusBadRequest)
    }
	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}
	return &tweet,nil
}





func (appModel *AppModel) UpdateTweet(request *request.CreateTweetRequest,id int) (*app.Error) {

	tweet := Tweet{
		Content:  request.Content,
	}
  var tweet1 Tweet

	var result *gorm.DB = appModel.DB.Model(&tweet)

	if id != 0 {

    result.Where("id = ?",id).First(&tweet1)
    if result.Error != nil {
      return app.
				NewError(result.Error).
				SetMessage("not found").
				SetCode(http.StatusBadRequest)
    }



	if request.Content !="" {
		result.Where("id = ?",id).Update("content", tweet.Content)
	if result.Error != nil {
		me, ok := result.Error.(*mysql.MySQLError)
		if !ok {
			return nil
		}
		if me.Number == 1062 {
			return  app.
				NewError(result.Error).
				SetMessage("already taken").
				SetCode(http.StatusBadRequest)
		}
		return nil
	}
}
}
	return nil
}




func(appModel *AppModel) DeleteTweet(id int)(*app.Error)  {
	var tweet Tweet
	var where *gorm.DB = appModel.DB
	if id != 0 {
		where = appModel.DB.Delete(&tweet,id)
	}
	if where.Error != nil {
		return app.NewError(where.Error).SetCode(http.StatusNotFound)
	}

	return nil
}
