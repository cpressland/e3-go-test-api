package contracts

// DeleteResponse is the response contract for the Delete method
type DeleteUserResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Deleted bool   `json:"result"`
}
