package request

type CreateTweetRequest struct {
	User_id      int    `json:"user_id,omitempty"`
	Content      string `json:"content,omitempty"`
	Parent_tweet int    `json:"parent_tweet,omitempty"`
}
