package model

type ErrorResponse struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Code uint32 `json:"code"`
	Data Token  `json:"data"`
}

type LogoutResponse struct {
	Code uint32 `json:"code"`
	Data string `json:"data"`
}

type RegisterResponse struct {
	Code uint32 `json:"code"`
	Data string `json:"data"`
}

type ChangePasswordResponse struct {
	Code uint32 `json:"code"`
	Data string `json:"data"`
}

type UserInfoResponse struct {
	Code uint32     `json:"code"`
	Data UserInfoDB `json:"data"`
}

type UsersInfoResponse struct {
	Code uint32            `json:"code"`
	Data []UserBriefInfoDB `json:"data"`
}

type FetchRolesResponse struct {
	Code uint32   `json:"code"`
	Data []RoleDB `json:"data"`
}

type UpdateUserRolesResponse struct {
	Code uint32 `json:"code"`
	Data string `json:"data"`
}

type UpdateRoleResponse struct {
	Code uint32 `json:"code"`
	Data string `json:"data"`
}
