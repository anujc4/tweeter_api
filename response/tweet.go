package response

import (
	"time"

	"github.com/anujc4/tweeter_api/model"
)

type TweetResponse struct {
	ID          uint      `json:"id,omitempty"`
	UserID      uint      `json:"user_id,omitempty"`
	Content     string    `json:"content,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	ParentTweet uint      `json:"parent_tweet,omitempty"`
}

func TransformTweetResponse(tweet model.Tweet) TweetResponse {
	return TweetResponse{
		ID:          tweet.ID,
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
