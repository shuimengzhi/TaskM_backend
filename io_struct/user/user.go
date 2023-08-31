package user_io_struct

type RegisterRequest struct {
	Account  string `json:"account"`
	NickName string `json:"nick_name"`
	RoleType int8   `json:"role_type"`
	Avatar   string `json:"avatar"`
}
