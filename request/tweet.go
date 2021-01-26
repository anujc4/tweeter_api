package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

//CreateTweetRequest describes a create users request
type CreateTweetRequest struct {
	UserID      uint   `json:"user_id,omitempty"`
	Content     string `json:"content,omitempty"`
	ParentTweet uint   `json:"parent,omitempty"`
}

//ValidateCreateTweetRequest validates CreateUserRequest object
func (r CreateTweetRequest) ValidateCreateTweetRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserID, validation.Required),
		validation.Field(&r.Content, validation.Required, validation.Length(1, 100)),
	)
}

//GetTweetsRequest describes a get users request
type GetTweetsRequest struct {
	PaginationRequest
	UserID      string `schema:"user_id"`
	ParentTweet string `schema:"parent_tweet"`
	Content     string `schema:"content"`
}

//UpdateTweetRequest describes an update user request
type UpdateTweetRequest struct {
	ID      string
	Content string `json:"content,omitempty"`
}

//ValidateUpdateTweetRequest validates UpdateUserRequest object
func (r UpdateTweetRequest) ValidateUpdateTweetRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Content, validation.Required, validation.Length(1, 100)),
		validation.Field(&r.ID, validation.Required),
	)
}
