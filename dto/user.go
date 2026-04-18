package dto

type User struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type Userinfo struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type UserCreateReq struct {
	Userinfo
	Password string `json:"password,omitempty" binding:"required"`
}

type UserUpdateReq struct {
	ID
	Userinfo
}

type UserDeleteReq struct {
	ID
}
