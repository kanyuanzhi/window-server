package handle

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"userServer/model"
	"userServer/util"
)

func FetchRoles(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Token")
	log.Println(token)

	allRoles := util.GetAllRoles()
	var response interface{}
	response = model.FetchRolesResponse{Code: 20000, Data: allRoles}

	responseByte, _ := json.Marshal(response)
	w.Write(responseByte)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Token")
	log.Println(token)

	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
	}
	role := model.RoleDB{}
	err = json.Unmarshal(result, &role)
	if err != nil {
		log.Println(err.Error())
	}
	var response interface{}
	if ok:=util.UpdateRole(role); ok{
		response = model.UpdateRoleResponse{Code: 20000, Data: "更新权限说明成功"}
	}else {
		response = model.ErrorResponse{Code: 60204, Message: "更新权限说明失败"}
	}
	responseByte, _ := json.Marshal(response)
	w.Write(responseByte)
}
