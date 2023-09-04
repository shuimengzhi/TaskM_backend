package task_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	usercontroller "taskm/controllers/user"
	"taskm/enum"
	commonIoStruct "taskm/io_struct/common"
	taskiostruct "taskm/io_struct/task"
	"taskm/services"
	projectservice "taskm/services/project"
	taskservice "taskm/services/task"
)

// TaskTagCreate godoc
// @Summary      创建任务标签
// @Description  创建任务标签
// @Tags         project
// @Accept       json
// @Param request body taskiostruct.TaskTagCreateRequest true "params"
// @Success      200  {object}  commonIoStruct.Response
// @Router       /task_tag/create [post]
func TaskTagCreate(c *gin.Context) {
	userInfo := usercontroller.GetUserInfo(c)
	if userInfo == nil {
		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeBad, Msg: "获取不到用户信息"})
		return
	}

	var params taskiostruct.TaskTagCreateRequest
	err := c.BindJSON(&params)
	if err != nil {
		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeParamErr, Msg: err.Error()})
		return
	}
	// 检查当前用户是否参与项目
	if authResult := projectservice.ProjectAuthCheck(userInfo.UID, params.ProjectId); authResult.Code == services.FAIL {
		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeBad, Msg: authResult.Msg, Data: authResult.Data})
		return
	}
	result := taskservice.TaskTagCreate(params)

	if result.Code == services.FAIL {
		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeBad, Msg: result.Msg, Data: result.Data})
		return
	}
	c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeOk, Data: result.Data, Msg: result.Msg})
	return
}
