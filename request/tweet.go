package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateTweetRequest struct {
	UserID      uint   `json:"user_id,omitempty"`
	Content     string `json:"content,omitempty"`
	ParentTweet uint   `json:"parent_tweet,omitempty"`
}

func (r CreateTweetRequest) ValidateCreateTweetRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserID, validation.Required),
		validation.Field(&r.Content, validation.Required, validation.Length(1, 240)),
		validation.Field(&r.ParentTweet, validation.Required),
	)
}
