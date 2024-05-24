package contracts

// UpdateRequest is the request contract for the Update method
type UpdateUserRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}
