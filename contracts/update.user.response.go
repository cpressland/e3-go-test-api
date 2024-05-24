package contracts

// UpdateResponse is the response contract for the Update method
type UpdateUserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}
