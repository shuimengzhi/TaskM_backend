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

// Register godoc
// @Summary      注册
// @Description  注册
// @Tags         user
// @Accept       json
// @Param request body userIoStruct.RegisterRequest true "params"
// @Success      200  {object}  commonIoStruct.Response
// @Router       /user/register [post]
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

// Login godoc
// @Summary      登陆
// @Description  登陆
// @Tags         user
// @Accept       json
// @Param request body userIoStruct.LoginRequest true "params"
// @Success      200  {object}  commonIoStruct.Response
// @Router       /user/login [post]
func Login(c *gin.Context) {
	var params userIoStruct.LoginRequest
	err := c.BindJSON(&params)
	if err != nil {
		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeParamErr, Msg: err.Error()})
		return
	}

	result := userservice.Login(params)
	if result.Code == services.FAIL {
		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeBad, Msg: result.Msg, Data: result.Data})
		return
	}
	c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeOk, Data: result.Data, Msg: result.Msg})
	return
}

func GetUserInfo(c *gin.Context) *userIoStruct.LoginResponse {
	if user, _ := c.Get("user_info"); user != nil {
		if u, ok := user.(*userIoStruct.LoginResponse); ok {
			return u
		}
	}
	return nil
}
