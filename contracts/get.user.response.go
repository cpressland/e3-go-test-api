package contracts

// GetUserResponse is the response contract for the Get method, and the response contract for the List method
type GetUserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}
