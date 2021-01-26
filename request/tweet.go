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

/*
//UpdateTweetRequest describes an update user request
type UpdateTweetRequest struct {
	ID        string
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

//ValidateUpdateTweetRequest validates UpdateUserRequest object
func (r UpdateTweetRequest) ValidateUpdateTweetRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, is.Email),
		validation.Field(&r.FirstName, validation.Length(3, 20)),
		validation.Field(&r.LastName, validation.Length(3, 20)),
	)
}

// eg. Validation if either phone no or email was required
// func (r GetUsersRequest) ValidateGetUsersRequest() error {
// 	return validation.ValidateStruct(&r,
// 		validation.Field(&r.Email, validation.Required.When(r.FirstName == "").Error("Either phone or Email is required.")),
// 		validation.Field(&r.FirstName, validation.Required.When(r.Email == "").Error("Either phone or Email is required.")),
// 	)
// }
*/
