package user_service

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	userIoStruct "taskm/io_struct/user"
	"taskm/model"
	"taskm/services"
	"time"
)

// EncryptPWD 密码加密方法
func EncryptPWD(password string) string {
	key := "taskm-666"
	hashOne := md5.Sum([]byte(password))
	once := hex.EncodeToString(hashOne[:])
	hashTwo := md5.Sum([]byte(once + key))
	twice := hex.EncodeToString(hashTwo[:])

	twiceString := twice[8:24]
	hashThrice := md5.Sum([]byte(twiceString))
	thrice := hex.EncodeToString(hashThrice[:])
	return strings.ToLower(thrice)
}

func Register(params userIoStruct.RegisterRequest) services.ResultService {
	var userModel model.User
	var accountCount int64

	if err := model.DB.Model(&userModel).Where("u_account = ?", params.Account).Count(&accountCount).Error; err != nil {
		return services.ResultService{Code: services.FAIL, Msg: err.Error() + " Register:1"}
	}
	if accountCount > 0 {
		return services.FailResponse(params.Account+"已存在", "")
	}

	user := model.User{
		UAccount:    params.Account,
		UNickName:   params.NickName,
		UPassword:   EncryptPWD("123456"),
		UStatus:     params.RoleType,
		UAvatar:     params.Avatar,
		UCreateTime: int32(time.Now().Unix()),
		URoleType:   params.RoleType,
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return services.ResultService{Code: services.FAIL, Msg: err.Error() + " Register:2"}
	}
	return services.SuccessResponse("", params.Account+"注册成功")
}
