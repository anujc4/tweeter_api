package response

import (
	"time"

	"github.com/anujc4/tweeter_api/model"
)

type UserResponse struct {
	ID        uint      `json:"id,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func TransformUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func MapUsersResponse(vs model.Users, f func(model.User) UserResponse) []UserResponse {
	vsm := make([]UserResponse, len(vs))
	for i := range vs {
		vsm[i] = f(*vs[i])
	}
	return vsm
}

type TweetResponse struct {
	ID          int       `json:"id,omitempty"`
	UserId      int       `json:"user_id,omitempty"`
	Content     string    `json:"content,omitempty"`
	ParentTweet int       `json:"parent_tweet,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func TransformTweetResponse(tweet model.Tweet) TweetResponse {
	return TweetResponse{
		ID:          tweet.ID,
		UserId:      tweet.UserId,
		Content:     tweet.Content,
		ParentTweet: tweet.ParentTweet,
		CreatedAt:   tweet.CreatedAt,
		UpdatedAt:   tweet.UpdatedAt,
	}
}

func MapTweetsResponse(vs model.Tweets, f func(model.Tweet) TweetResponse) []TweetResponse {
	vsm := make([]TweetResponse, len(vs))
	for i := range vs {
		vsm[i] = f(*vs[i])
	}
	return vsm
}
