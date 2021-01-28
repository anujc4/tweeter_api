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



type GetTweetsRequest struct {
	PaginationRequest
	ID  int `schema:"ID"`
	UserId     int `schema:"user_id"`
	ParentTweet int `schema:"parent_tweet"`
}

type CreateTweetRequest struct {
	UserId int `json:"user_id,omitempty"`
	Content  string `json:"content,omitempty"`
	ParentTweet int `json:"parent_tweet,omitempty"`
}
func (r CreateTweetRequest) ValidateCreateTweetRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserId, validation.Required),
		validation.Field(&r.Content, validation.Required),
	)
}


// eg. Validation if either phone no or email was required
// func (r GetUsersRequest) ValidateGetUsersRequest() error {
// 	return validation.ValidateStruct(&r,
// 		validation.Field(&r.Email, validation.Required.When(r.FirstName == "").Error("Either phone or Email is required.")),
// 		validation.Field(&r.FirstName, validation.Required.When(r.Email == "").Error("Either phone or Email is required.")),
// 	)
// }
