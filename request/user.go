package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateUserRequest struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

func (r CreateUserRequest) ValidateCreateUserRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.FirstName, validation.Required, validation.Length(3, 20)),
		validation.Field(&r.LastName, validation.Required, validation.Length(3, 20)),
		validation.Field(&r.Password, validation.Required, validation.Length(5, 50)),
	)
}

type GetUsersRequest struct {
	PaginationRequest
	FirstName string `schema:"first_name"`
	Email     string `schema:"email"`
}
// eg. Validation if either phone no or email was required
// func (r GetUsersRequest) ValidateGetUsersRequest() error {
// 	return validation.ValidateStruct(&r,
// 		validation.Field(&r.Email, validation.Required.When(r.FirstName == "").Error("Either phone or Email is required.")),
// 		validation.Field(&r.FirstName, validation.Required.When(r.Email == "").Error("Either phone or Email is required.")),
// 	)
// }


type GetUserByIDRequest struct{
	ID uint `schema:"id"`
}

func (r GetUserByIDRequest) ValidateGetUserByIDRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, validation.Match(regexp.MustCompile("[0-9]"))),
}

type UpdateUserRequest struct {
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

//Password & User_ID is required to update/delete a user
func (r UpdateUserRequest) UpdateUserRequestIDRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, validation.Match(regexp.MustCompile("[0-9]"))),
		validation.Field(&r.FirstName, validation.Length(3, 20)),
		validation.Field(&r.LastName, validation.Length(3, 20)),
		validation.Field(&r.Password, validation.Required, validation.Length(5, 50)),
}

