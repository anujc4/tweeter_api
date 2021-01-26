package model

import (
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
		Find(&users)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &users, nil
}

func (appModel *AppModel) GetUserById(id int) (*User, *app.Error) {
	var user User
	var where *gorm.DB = appModel.DB

	where = appModel.DB.Where("ID = ?", id)

	result := where.Find(&user)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &user, nil
}

func (appModel *AppModel) UpdateUserById(request *request.CreateUserRequest, id int) (*User, *app.Error) {
	var user User
	var where *gorm.DB = appModel.DB

	result := where.First(&user, id)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	if request.FirstName != "" {
		user.FirstName = request.FirstName
	}
	if request.LastName != "" {
		user.LastName = request.LastName
	}
	if request.Email != "" {
		user.Email = request.Email
	}

	result = where.Save(&user)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &user, nil
}

func (appModel *AppModel) DeleteUserById(id int) *app.Error {

	result := appModel.DB.Delete(&User{}, id)

	if result.Error != nil {
		return app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return nil
}
