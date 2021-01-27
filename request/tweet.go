package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateTweetRequest struct {
	UserId int `json:"user_id,omitempty"`
	Content string `json:"content,omitempty"`
	ParentTweet int `json:"parent_tweet,omitempty"`
}

func (request CreateTweetRequest) ValidateCreateTweetRequest()  error{
	return validation.ValidateStruct(&request,
		validation.Field(&request.Content, validation.Required),
		validation.Field(&request.UserId, validation.Required, is.Int),
		)
}

type GetTweetsRequest struct {
	PaginationRequest
	ID int `schema:"id"`
	UserID int `schema:"user_id"`
}

type UpdateTweetRequest struct {
	Content string `json:"content"`
}

func (request *UpdateTweetRequest) ValidateUpdateTweet()  error{
	return validation.ValidateStruct(&request,
		validation.Field(&request.Content, validation.Required),
		)
}