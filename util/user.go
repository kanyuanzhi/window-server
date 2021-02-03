package util

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
	"userServer/model"
)

// 登陆时检验用户名与密码是否一致
func CheckPasswordWithUsername(username string, password string) (bool, string) {
	filter := bson.D{{"username", username}}
	singleResult := FindOne("UserAuthentication", filter)
	if singleResult != nil {
		userAuthenticationDB := model.UserAuthenticationDB{}
		err := singleResult.Decode(&userAuthenticationDB)
		if err != nil {
			log.Println(err.Error())
			return false, ""
		}
		encodePassword := userAuthenticationDB.Password
		err = bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(password)) //验证（对比）
		if err != nil {
			return false, ""
		}
		return true, userAuthenticationDB.Token
	}
	return false, ""
}

// 修改密码时检验token与原密码是否一致
func CheckPasswordWithToken(token string, password string) bool {
	filter := bson.D{{"token", token}}
	singleResult := FindOne("UserAuthentication", filter)
	if singleResult != nil {
		userAuthenticationDB := model.UserAuthenticationDB{}
		err := singleResult.Decode(&userAuthenticationDB)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		encodePassword := userAuthenticationDB.Password
		err = bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(password)) //验证（对比）
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func UpdatePassword(userUpdatePassword model.UserUpdatePassword) bool {
	hash, err := bcrypt.GenerateFromPassword([]byte(userUpdatePassword.NewPassword), bcrypt.DefaultCost) //加密处理
	if err != nil {
		log.Println(err)
		return false
	}
	encodeNewPassword := string(hash)
	filter := bson.D{{"token", userUpdatePassword.Token}}
	update := bson.D{{"$set", bson.D{{"password", encodeNewPassword}}}}
	UpdateOne("UserAuthentication", filter, update)
	return true
}

func Register(userRegister model.UserRegister) bool {
	// 添加密码
	hash, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		log.Println(err)
		return false
	}
	encodePassword := string(hash)

	// 生成token：用户名+注册时间
	tokenStr := userRegister.Username + strconv.FormatInt(time.Now().Unix(),10)
	token := MD5Str(tokenStr)

	userAuthenticationDB := model.UserAuthenticationDB{Username: userRegister.Username, Password: encodePassword,Token: token}
	InsertOne("UserAuthentication", userAuthenticationDB)
	return true
}

func CheckUsername(username string) bool {
	filter := bson.D{{"username", username}}
	singleResult := FindOne("UserAuthentication", filter)
	userAuthenticationDB := model.UserAuthenticationDB{}
	err := singleResult.Decode(&userAuthenticationDB)
	if err != nil {
		// 解码失败，用户名不存在，可以使用
		log.Println(err.Error())
		return true
	}
	return false
}
