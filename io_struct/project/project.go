package project_io_struct

type CreateRequest struct {
	ProjectName string `json:"p_name" binding:"required"`
}

// ListRequest  表示项目列表请求结构体
type ListRequest struct {
	Page        int    `json:"page" binding:"required"` // 页码
	Size        int    `json:"size" binding:"required"` // 每页数量
	ProjectName string `json:"project_name"`            // 项目名称
}

// ProjectListItem 表示项目列表响应中的单个项目项结构体
type ProjectListItem struct {
	PID   int32  `json:"p_id"`   // 项目 ID
	PName string `json:"p_name"` // 项目名称
}

// ProjectListResponse 表示项目列表响应结构体
type ProjectListResponse struct {
	List  []ProjectListItem `json:"list"`  // 项目列表
	Count int64             `json:"count"` // 总数量
}
