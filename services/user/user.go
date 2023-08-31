package user_service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	_ "github.com/redis/go-redis/v9"
	"math/rand"
	"strconv"
	"strings"
	"taskm/cache"
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

func Token(UserId int32) string {
	number := strconv.Itoa(rand.Intn(100))
	seed := strconv.Itoa(int(UserId)) + "_" + strconv.FormatInt(time.Now().Unix(), 10) + "+" + number
	hashOne := md5.Sum([]byte(seed))
	return hex.EncodeToString(hashOne[:])
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

func Login(params userIoStruct.LoginRequest) services.ResultService {
	var userInfo model.User
	encryptPwd := EncryptPWD(params.Password)
	if err := model.DB.Select("u_id,u_account,u_nick_name,u_status,u_avatar,u_role_type,u_token").
		Where("u_account = ? and u_password = ?", params.UserName, encryptPwd).First(&userInfo).Error; err != nil {
		return services.ResultService{Code: services.FAIL, Msg: err.Error() + " Login:1"}
	}

	if userInfo == (model.User{}) {
		return services.FailResponse("账号或者密码错误", "")
	}
	if userInfo.UStatus == model.USER_STATUS_BAN {
		return services.FailResponse("账户被禁不能使用", "")
	}

	ctx := context.Background()
	token := Token(userInfo.UID)
	tokenInfo, err := cache.Instance.Get(ctx, token).Result()
	//如果有报错信息并且不是空的那种报错信息
	if err != nil && err != redis.Nil {
		return services.FailResponse(err.Error(), " Login:2")
	}
	//如果已经存在了就报错,重试
	if tokenInfo != "" {
		return services.FailResponse("服务繁忙,请重新尝试", "")
	}
	//删除数据库存的token
	cache.Instance.Del(ctx, userInfo.UToken)
	data := userIoStruct.LoginResponse{
		UID:      userInfo.UID,
		Account:  userInfo.UAccount,
		NickName: userInfo.UNickName,
		Status:   userInfo.UStatus,
		Avatar:   userInfo.UAvatar,
		RoleType: userInfo.URoleType,
		Token:    token,
	}

	dataJson, err := json.Marshal(data)
	if err != nil {
		return services.FailResponse(err.Error(), " Login:3")
	}
	//数据库更新token
	if err = model.DB.Model(&model.User{}).Where("u_id = ?", userInfo.UID).Updates(model.User{UToken: token}).Error; err != nil {
		return services.FailResponse(err.Error(), " Login:4")
	}
	setErr := cache.Instance.Set(ctx, token, dataJson, 2*time.Hour).Err()
	if setErr != nil {
		return services.FailResponse(setErr.Error(), " Login:4")
	}
	return services.SuccessResponse(data, "success")
}
