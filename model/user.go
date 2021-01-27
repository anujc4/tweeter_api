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
	}else if request.ID != 0 {
		where = appModel.DB.Where("id = ?", request.ID)
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

func (appModel *AppModel) GetUserById(id int) (*User, *app.Error){
	var user User
	var userModel *gorm.DB = appModel.DB
	userModel = appModel.DB.Where("id = ?", id)
	result := userModel.Find(&user)
	fmt.Println(result)
	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &user, nil
}

func (appModel *AppModel) UpdateUser(request *request.UpdateUserRequest,id int) (*User, *app.Error){
	if id != 0{
		var user User
		var userModel *gorm.DB = appModel.DB.Model(&user)
		result:= userModel.Where("id = ?", id).UpdateColumns(User{FirstName: request.FirstName, LastName: request.LastName, Email: request.Email})
		count := result.RowsAffected
		if result.Error != nil || count == 0{
			fmt.Println("Checkout")
			return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
		}
		return appModel.GetUserById(id)
	}
	return nil, nil
}

func (appModel *AppModel) DeleteUser(id int)  (*User, *app.Error){
	if id != 0{
		user, err := appModel.GetUserById(id)
		if err != nil{
			return nil, app.NewError(err).SetCode(http.StatusBadRequest)
		}
		result := appModel.DB.Where("id=?", id).Delete(&user)
		count := result.RowsAffected
		if result.Error != nil || count==0{
			println("Error Count", count)
			return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
		}
		return user, nil
	}
	return nil, nil
}
