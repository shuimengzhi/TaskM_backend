package task_io_struct

type TaskTagCreateRequest struct {
	ProjectId int32  `json:"project_id" binding:"required"`
	TagName   string `json:"tag_name" binding:"required"`
	Sort      int32  `json:"sort"`
	Color     string `json:"color" binding:"required"`
}
