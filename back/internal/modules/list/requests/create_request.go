package requests

type CreateListRequest struct {
	Name string `json:"name" binding:"required"`
}


