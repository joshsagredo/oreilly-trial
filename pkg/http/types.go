package http

type SuccessResponse struct {
	UserID string `json:"user_id"`
}

type FailureResponse struct {
	Email []string `json:"email"`
}