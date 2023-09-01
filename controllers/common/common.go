package common_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	usercontroller "taskm/controllers/user"
	"taskm/enum"
	"taskm/extend"
	commonIoStruct "taskm/io_struct/common"
)

func UploadAvatar(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")

	userInfo := usercontroller.GetUserInfo(c)
	ext := filepath.Ext(file.Filename)
	fileName := extend.GetMD5(strconv.Itoa(int(userInfo.UID))) + ext
	dst := "upload/avatar/" + fileName
	// 上传文件至指定的完整文件路径
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeBad, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, commonIoStruct.Response{Code: enum.CodeOk, Data: dst})
	return
}
func UploadCommentFile(c *gin.Context) {

}
func UploadTaskFile(c *gin.Context) {

}
