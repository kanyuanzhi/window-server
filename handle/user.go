package handle

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"userServer/model"
	"userServer/util"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
	}
	userLogin := model.UserLogin{}
	err = json.Unmarshal(result, &userLogin)
	if err != nil {
		log.Println(err.Error())
	}

	var response interface{}
	if ok, token := util.CheckPasswordWithUsername(userLogin.Username, userLogin.Password); ok {
		response = model.LoginResponse{Code: 20000, Data: model.Token{Token: token}}
	} else {
		response = model.ErrorResponse{Code: 60204, Message: "Account and password are incorrect."}
	}
	responseByte, _ := json.Marshal(response)
	w.Write(responseByte)
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	var response interface{}
	if ok, userInfo := util.GetUserInfo(token); ok {
		response = model.UserInfoResponse{Code: 20000, Data: userInfo}
	} else {
		response = model.ErrorResponse{Code: 50008, Message: "Login failed, unable to get user details."}
	}
	responseByte, _ := json.Marshal(response)
	w.Write(responseByte)
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	response := model.LogoutResponse{Code: 20000, Data: "success"}
	responseByte, _ := json.Marshal(response)
	w.Write(responseByte)
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	//todo:服务端验证该token是否有权限注册新用户
	token := r.URL.Query().Get("token")
	log.Println(token)

	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
	}
	userRegister := model.UserRegister{}
	err = json.Unmarshal(result, &userRegister)
	if err != nil {
		log.Println(err.Error())
	}

	var response interface{}

	if ok := util.CheckUsername(userRegister.Username); !ok {
		response = model.ErrorResponse{Code: 60300, Message: "该用户名已被注册"}
	} else {
		if ok := util.Register(userRegister); ok {
			response = model.RegisterResponse{Code: 20000, Data: "注册成功"}
		} else {
			response = model.ErrorResponse{Code: 60204, Message: "添加用户失败"}
		}
	}
	responseByte, _ := json.Marshal(response)
	w.Write(responseByte)
}

func UserUpdatePassword(w http.ResponseWriter, r *http.Request) {
	//todo:服务端验证
	token := r.URL.Query().Get("token")
	log.Println(token)

	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
	}
	userUpdatePassword := model.UserUpdatePassword{}
	err = json.Unmarshal(result, &userUpdatePassword)
	if err != nil {
		log.Println(err.Error())
	}
	userUpdatePassword.Token = token

	var response interface{}

	if ok := util.CheckPasswordWithToken(token, userUpdatePassword.NewPassword); ok {
		if ok := util.UpdatePassword(userUpdatePassword); ok {
			response = model.ChangePasswordResponse{Code: 20000, Data: "修改密码成功"}
		} else {
			response = model.ErrorResponse{Code: 60204, Message: "修改密码失败"}
		}
	} else {
		response = model.ErrorResponse{Code: 60204, Message: "原密码输入错误"}
	}
	responseByte, _ := json.Marshal(response)
	w.Write(responseByte)
}

func FetchUsers(w http.ResponseWriter, r *http.Request) {
	usersBriefInfo := util.GetUsersBriefInfo()
	var response interface{}
	response = model.UsersInfoResponse{Code: 20000, Data: usersBriefInfo}
	responseByte, _ := json.Marshal(response)
	w.Write(responseByte)
}

func UpdateUserRoles(w http.ResponseWriter, r *http.Request){
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
	}
	userBriefInfo := model.UserBriefInfoDB{}
	err = json.Unmarshal(result, &userBriefInfo)
	if err != nil {
		log.Println(err.Error())
	}

	var response interface{}
	if ok := util.UpdateUserRoles(userBriefInfo); ok {
		response = model.UpdateUserRolesResponse{Code: 20000, Data: "修改权限成功"}
	} else {
		response = model.ErrorResponse{Code: 60204, Message: "修改权限失败"}
	}
	responseByte, _ := json.Marshal(response)
	w.Write(responseByte)}
