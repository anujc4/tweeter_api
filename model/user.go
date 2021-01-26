package model

import (
	"net/http"
	"time"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/request"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// User represents a user object in our app
type User struct {
	ID        uint `gorm:"primarykey"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Users alias list of 'User' type
type Users []*User

//CreateUser creates user object and adds it to database
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

// GetUsers returns all users from database (paginated)
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
		Find(&users)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &users, nil
}

// GetUserByID returns a user given a userID
func (appModel AppModel) GetUserByID(userID string) (User, *app.Error) {
	var user User
	tx := appModel.DB.First(&user, userID)
	if tx.Error != nil {
		return User{}, app.NewError(tx.Error).SetCode(http.StatusNotFound)
	}
	return user, nil
}

// DeleteUser deletes user form database given an ID
func (appModel AppModel) DeleteUser(userID string) *app.Error {
	tx := appModel.DB.Delete(&User{}, userID)
	if tx.RowsAffected == 0 {
		return app.NewError(errors.New("Object does not exist")).SetCode(404)
	}
	return nil
}

// UpdateUser updates a user in database
func (appModel AppModel) UpdateUser(user request.UpdateUserRequest) (User, *app.Error) {
	var userObject User
	tx := appModel.DB.Find(&userObject, user.ID).Updates(
		User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
	)

	if tx.RowsAffected == 0 {
		return User{}, app.NewError(errors.New("Object does not exist")).SetCode(404)
	}
	return userObject, nil
}
