package contracts

// CreateUserResponse is the response contract for the Create method
type CreateUserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}
