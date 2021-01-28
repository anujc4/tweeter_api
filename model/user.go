package model

import (
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

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
	Tweets    []Tweet
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users []*User

func (appModel *AppModel) CreateUser(request *request.CreateUserRequest) (*User, *app.Error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, app.NewError(err)
	}

	user := User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  string(hashedPassword),
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

func (appModel *AppModel) VerifyUserCredential(request *request.LoginRequest) (*User, *app.Error) {
	var user User

	result := appModel.DB.
		Where("email = ?", request.Email).
		First(&user)

	if result.Error != nil {
		return nil, app.NewError(result.Error)
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

func (appModel *AppModel) GetUser(w map[string]interface{}) (*User, *app.Error) {
	var user User
	result := appModel.DB.Where(w).First(&user)
	if result.Error != nil {
		return nil, app.NewError(result.Error)
	}

	return &user, nil
}

func (appModel *AppModel) UpdateUser(id uint, update *request.UpdateUserRequest) (*User, *app.Error) {
	user := User{
		FirstName: update.FirstName,
		LastName:  update.LastName,
		Password:  update.Password,
	}

	// Passing a struct will work as GORM will only update non-zero fields
	// https://gorm.io/docs/update.html#Updates-multiple-columns
	result := appModel.DB.First(&user, id).Updates(user)
	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &user, nil
}

func (appModel *AppModel) DeleteUser(id uint) (int64, error) {
	result := appModel.DB.Delete(&User{}, id)
	return result.RowsAffected, result.Error
}

func (appModel *AppModel) FindAllTweetsByUser(userID uint) (*User, *app.Error) {
	var user User
	result := appModel.DB.Preload("Tweets").First(&user, userID)
	if result.Error != nil {
		return nil, app.NewError(result.Error)
	}
	return &user, nil
}
