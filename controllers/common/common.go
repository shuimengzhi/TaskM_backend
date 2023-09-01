package common_controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"taskm/enum"
	commonIoStruct "taskm/io_struct/common"
)

func UploadAvatar(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := "upload/avatar/" + file.Filename
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
