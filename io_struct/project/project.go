package project

type CreateRequest struct {
	ProjectName string `json:"p_name" binding:"required"`
}
