package model

type Token struct {
	Token string `json:"token"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserAuthenticationDB struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserRegister struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Token           string `json:"token"`
}

type UserUpdatePassword struct {
	Token              string `json:"token"`
	OldPassword        string `json:"old_password"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}

type UserInfoDB struct {
	Token        string   `json:"token"`
	Roles        []string `json:"roles"` // guest(default),operator,admin,super
	Introduction string   `json:"introduction"`
	Username     string   `json:"username"`
	Name         string   `json:"name"`
	Avatar       string   `json:"avatar"`
}

type UserBriefInfoDB struct {
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Username     string   `json:"username"`
	Name         string   `json:"name"`
}