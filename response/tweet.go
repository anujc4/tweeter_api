package response

import (
	"time"

	"github.com/anujc4/tweeter_api/model"
)

// TweetResponse represents a tweet response
type TweetResponse struct {
	ID          uint      `json:"id,omitempty"`
	Content     string    `json:"first_name,omitempty"`
	UserID      uint      `json:"last_name"`
	ParentTweet uint      `json:"email,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

//TransformTweetResponse transforms a model.Tweet into TweetResponse
func TransformTweetResponse(tweet model.Tweet) TweetResponse {
	return TweetResponse{
		ID:          tweet.ID,
		Content:     tweet.Content,
		UserID:      tweet.UserID,
		ParentTweet: tweet.ParentTweet,
		CreatedAt:   tweet.CreatedAt,
		UpdatedAt:   tweet.UpdatedAt,
	}
}

//MapTweetResponse does what it says ig
// func MapTweetResponse(vs model.Users, f func(model.User) UserResponse) []UserResponse {
// 	vsm := make([]UserResponse, len(vs))
// 	for i := range vs {
// 		vsm[i] = f(*vs[i])
// 	}
// 	return vsm
// }
