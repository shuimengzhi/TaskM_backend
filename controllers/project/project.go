package project_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	usercontroller "taskm/controllers/user"
	"taskm/enum"
	commonIoStruct "taskm/io_struct/common"
	"taskm/io_struct/project"
	"taskm/services"
	projectservice "taskm/services/project"
)

// ProjectCreate godoc
// @Summary      创建项目
// @Description  创建项目
// @Tags         project
// @Accept       json
// @Param request body project.CreateRequest true "params"
// @Success      200  {object}  commonIoStruct.Response
// @Router       /task/project/create [post]
func ProjectCreate(c *gin.Context) {
	userInfo := usercontroller.GetUserInfo(c)
	if userInfo == nil {
		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeBad, Msg: "获取不到用户信息"})
		return
	}

	var params project.CreateRequest
	err := c.BindJSON(&params)
	if err != nil {
		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeParamErr, Msg: err.Error()})
		return
	}
	result := projectservice.ProjectCreate(params, userInfo)

	if result.Code == services.FAIL {
		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeBad, Msg: result.Msg, Data: result.Data})
		return
	}
	c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeOk, Data: result.Data, Msg: result.Msg})
	return
}
