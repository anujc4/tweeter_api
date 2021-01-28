package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateTweetRequest struct {
	Content     string `json:"content,omitempty"`
	ParentTweet *uint  `json:"parent_tweet,omitempty"`
}

func (r CreateTweetRequest) ValidateCreateTweetRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Content, validation.Required),
	)
}
