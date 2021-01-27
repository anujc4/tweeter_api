package response

import (
	"time"

	"github.com/anujc4/tweeter_api/model"
)

type TweetResponse struct {
	ID          int       `json:"id,omitempty"`
	Content     string    `json:"content,omitempty"`
	UserId      int       `json:"user_id,omitempty"`
	ParentTweet int       `json:"parent_tweet,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func TransformTweetResponse(tweet model.Tweet) TweetResponse {
	return TweetResponse{
		ID:          tweet.ID,
		Content:     tweet.Content,
		ParentTweet: tweet.Parent_tweet,
		UserId:      tweet.User_id,
		CreatedAt:   tweet.CreatedAt,
		UpdatedAt:   tweet.UpdatedAt,
	}
}

func MapTweetResponse(vs model.Tweets, f func(model.Tweet) TweetResponse) []TweetResponse {
	vsm := make([]TweetResponse, len(vs))
	for i := range vs {
		vsm[i] = f(*vs[i])
	}
	return vsm
}
