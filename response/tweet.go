package response

import (
	"time"

	"github.com/anujc4/tweeter_api/model"
)

type TweetResponse struct {
	ID int `json:"id,omitempty"`
	UserId int `json:"user_id,omitempty"`
	Content string `json:"content,omitempty"`
	ParentTweet int `json:"parent_tweet,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func TransformTweetResponse(tweet model.Tweet) TweetResponse {
	return TweetResponse{
		ID: tweet.ID,
		UserId: tweet.UserId,
		Content: tweet.Content,
		ParentTweet: tweet.ParentTweet,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
	}
}

func MapTweetResponse(vs model.Tweets, f func(tweet model.Tweet) TweetResponse) []TweetResponse {
	vsm := make([]TweetResponse, len(vs))
	for i := range vs {
		vsm[i] = f(*vs[i])
	}
	return vsm
}
