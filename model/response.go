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
