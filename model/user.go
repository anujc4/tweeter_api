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

	if request.Page == 0 {
		page = 1
	}

	switch {
	case request.PageSize > 100:
		pageSize = 100
	case request.PageSize <= 0:
		pageSize = 10
	case request.PageSize > 0:
		pageSize = request.PageSize
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


func (appModel *AppModel) GetUserByID(id int) (*User, *app.Error) {
	var users User
	var where *gorm.DB = appModel.DB

	if id > 0 {
		where = appModel.DB.Where("id = ?", id)
	} 

	
	result := where.
		Find(&users)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}
 
	return &users, nil
}


func (appModel *AppModel) UpdateUser(request *request.CreateUserRequest,id int) (*app.Error) {
	var user1 User
	user := User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}
	var result *gorm.DB = appModel.DB.Model(&user)
	if id > 0 {
		result.Where("id = ?",id).First(&user1)
		if result.Error != nil {
			return app.
						NewError(result.Error).
						SetMessage("user not found").
						SetCode(http.StatusBadRequest)
		}


		if user.FirstName !="" {
			result.Where("id = ?",id).Update("first_name", user.FirstName)
			if result.Error != nil {
				me, ok := result.Error.(*mysql.MySQLError)
				if !ok {
					return nil
				}
				if me.Number > 0 {
					return nil
				}
				return nil
			}
		}
		if request.LastName !="" {
			result.Where("id = ?",id).Update("last_name", user.LastName)
			if result.Error != nil {
				me, ok := result.Error.(*mysql.MySQLError)
				if !ok {
					return nil
				}
				if me.Number > 0 {
					return nil
				}
				return nil
			}

		}
		if request.Email !="" {
			result.Where("id = ?",id).Update("email", user.Email)
			if result.Error != nil {
				me, ok := result.Error.(*mysql.MySQLError)
				if !ok {
					return nil
				}
				if me.Number == 1062 {
					return  app.
						NewError(result.Error).
						SetMessage("Email " + request.Email + " is already taken").
						SetCode(http.StatusBadRequest)
				}
				return nil
			}
		}
	}
	return nil
}

func(appModel *AppModel) DeleteUser(id int)(*app.Error)  {
	var users User
	var where *gorm.DB = appModel.DB
	if id > 0 {
		where = appModel.DB.Delete(&users,id)
	}
	if where.Error != nil {
		return app.
			NewError(where.Error).
			SetMessage("user not found").
			SetCode(http.StatusBadRequest)
	}
	return nil
}