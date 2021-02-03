package testData

import "userServer/model"

var UserToken1 = model.Token{Token: "admin-token"}
var UserToken2 = model.Token{Token: "editor-token"}

var UserInfo1 = model.UserInfo{Roles: []string{"admin"}, Introduction: "admin", Username: "admin"}
var UserInfo2 = model.UserInfo{Roles: []string{"editor"}, Introduction: "editor", Username: "editor"}

var TokenMap1 = model.TokenMap{"admin-token": UserInfo1}
var TokenMap2 = model.TokenMap{"editor-token": UserInfo2}

