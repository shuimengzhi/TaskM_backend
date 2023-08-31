package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taskm/enum"
	commonIoStruct "taskm/io_struct/common"
	userIoStruct "taskm/io_struct/user"
	"taskm/services"
	userservice "taskm/services/user"
)

func Register(c *gin.Context) {
	var params userIoStruct.RegisterRequest
	err := c.BindJSON(&params)
	if err != nil {
		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeParamErr, Msg: err.Error()})
		return
	}

	result := userservice.Register(params)
	if result.Code == services.FAIL {
		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeBad, Msg: result.Msg, Data: result.Data})
		return
	}
	c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeOk, Data: result.Data, Msg: result.Msg})
	return
}

//func Login(c *gin.Context) {
//	var params systemIoStruct.LoginRequest
//	err := c.BindJSON(&params)
//	if err != nil {
//		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeParamErr, Msg: err.Error()})
//		return
//	}
//	//  判断验证码是否正确
//	if os.Getenv("GIN_MODE") == enum.ModeProduct {
//		// 用redis作为存储
//		var store = captcha.NewStore(cache.Instance, context.Background())
//		//verify the captcha
//		if !store.Verify(params.CaptchaId, params.CaptchaValue, true) {
//			c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeBad, Msg: "验证码验证失败"})
//			return
//		}
//	}
//	result := service2.LoginService(params)
//	if result.Code == service.FAIL {
//		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeBad, Msg: result.Msg, Data: result.Data})
//		return
//	}
//	c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeOk, Data: result.Data, Msg: result.Msg})
//	return
//}
