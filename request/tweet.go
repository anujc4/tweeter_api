package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateTweetRequest struct {
	UserId      uint   `json:"user_id,omitempty"`
	Content     string `json:"content,omitempty"`
	ParentTweet uint   `json:"parent_tweet,omitempty"`
}

func (r CreateTweetRequest) ValidateCreateTweetRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserId, validation.Required, is.Digit),
		validation.Field(&r.Content, validation.Required, validation.Length(3, 20)),
		validation.Field(&r.ParentTweet, validation.Required, is.Digit),
	)
}

type GetTweetsRequest struct {
	PaginationRequest
	Content string `schema:"content"`
	UserId  string `schema:"user_id"`
}

// eg. Validation if either phone no or email was required
// func (r GetUsersRequest) ValidateGetUsersRequest() error {
// 	return validation.ValidateStruct(&r,
// 		validation.Field(&r.Email, validation.Required.When(r.FirstName == "").Error("Either phone or Email is required.")),
// 		validation.Field(&r.FirstName, validation.Required.When(r.Email == "").Error("Either phone or Email is required.")),
// 	)
// }
