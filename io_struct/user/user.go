package user_io_struct

type RegisterRequest struct {
	Account  string `json:"account" binding:"required"`
	NickName string `json:"nick_name" binding:"required"`
	RoleType int8   `json:"role_type" binding:"required"`
	Avatar   string `json:"avatar"`
}

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResponse struct {
	UID      int32  `json:"u_id"`
	Account  string `json:"u_account"`
	NickName string `json:"u_nick_name"`
	Status   int8   `json:"u_status"`
	Avatar   string `json:"u_avatar"`
	RoleType int8   `json:"u_role_type"`
	Token    string `json:"token"`
}
