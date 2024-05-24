package contracts

// CreateUserRequest is the request contract for the Create method
type CreateUserRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}
