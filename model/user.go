package model

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/request"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tweet struct {
	ID          int `gorm:"primarykey"`
	UserId      int
	Content     string
	ParentTweet int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Users []*User

func (appModel *AppModel) CreateUser(request *request.CreateUserRequest) (*User, *app.Error) {
	user := User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}
	result := appModel.DB.Create(&user)

	if result.Error != nil {
		me, ok := result.Error.(*mysql.MySQLError)
		if !ok {
			return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
		}
		if me.Number == 1062 {
			return nil, app.
				NewError(result.Error).
				SetMessage("Email " + request.Email + " is already taken").
				SetCode(http.StatusBadRequest)
		}
		return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
	}
	return &user, nil
}

func (appModel *AppModel) GetUsers(request *request.GetUsersRequest) (*Users, *app.Error) {
	var users Users
	var where *gorm.DB = appModel.DB
	var page, pageSize int

	if request.Email != "" {
		where = appModel.DB.Where("email = ?", request.Email)
	} else if request.FirstName != "" {
		where = appModel.DB.Where("first_name LIKE ?", "%"+request.FirstName+"%")
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

	result := where.
		Offset(offset).
		Limit(pageSize).
		Find(&users)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &users, nil
}

func (appModel *AppModel) GetUserByID(ID int64, request *request.GetUserByIDRequest) (*Users, *app.Error) {
	var users Users

	result := appModel.DB.Find(&users, ID)
	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &users, nil
}

func (appModel *AppModel) UpdateUser(ID int64, request *request.CreateUserRequest) (*User, *app.Error) {
	var user User
	result := appModel.DB.First(&user, ID)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	fmt.Println(request.FirstName)
	fmt.Println("Hello-3")
	if request.FirstName != "" {
		user.FirstName = request.FirstName
	}
	if request.LastName != "" {
		user.LastName = request.LastName
	}
	if request.Email != "" {
		user.Email = request.Email
	}

	result = appModel.DB.Save(&user)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &user, nil

}

func (appModel *AppModel) DeleteUserByID(ID int64, request *request.GetUserByIDRequest) *app.Error {

	result := appModel.DB.Delete(&User{}, ID)

	if result.Error != nil {
		return app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return nil
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
		where = appModel.DB.Where("first_name = ?", "%", request.ParentTweet)
	} else if request.UserId != 0 {
		where = appModel.DB.Where("user_id = ?", request.UserId)
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

	result := where.
		Offset(offset).
		Limit(pageSize).
		Find(&tweets)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &tweets, nil
}

func (appModel *AppModel) GetTweetByID(ID int64, request *request.GetTweetByIDRequest) (*Tweets, *app.Error) {
	var tweets Tweets

	result := appModel.DB.Find(&tweets, ID)
	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &tweets, nil
}

func (appModel *AppModel) UpdateTweet(ID int64, request *request.CreateTweetRequest) (*Tweet, *app.Error) {
	var tweet Tweet
	result := appModel.DB.First(&tweet, ID)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	if request.ParentTweet != 0 {
		tweet.ParentTweet = request.ParentTweet
	}
	if request.UserId != 0 {
		tweet.UserId = request.UserId
	}

	result = appModel.DB.Save(&tweet)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &tweet, nil

}

func (appModel *AppModel) DeleteTweetByID(ID int64, request *request.GetTweetByIDRequest) *app.Error {

	result := appModel.DB.Delete(&Tweet{}, ID)

	if result.Error != nil {
		return app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return nil
}
