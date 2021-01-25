package response

type LoginResponse struct {
	Email string `json:"email,omitempty"`
	Token string `json:"token,omitempty"`
}
