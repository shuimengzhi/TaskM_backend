package common_controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"taskm/core"
	"taskm/extend"
	commoniostruct "taskm/io_struct/common"
	useriostruct "taskm/io_struct/user"
	router2 "taskm/router"
	"testing"
)

func setupLogin() string {
	envPath, _ := os.Getwd()
	envPath = envPath + "/../../.env"
	core.LoadCore(envPath)
	route := router2.NewRouter()
	w := httptest.NewRecorder()
	request := useriostruct.LoginRequest{UserName: "admin", Password: "123456"}
	infoJson, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(string(infoJson)))
	route.ServeHTTP(w, req)

	var resp commoniostruct.Response
	var userInfo useriostruct.LoginResponse
	json.Unmarshal([]byte(w.Body.String()), &resp)
	dataJson, _ := json.Marshal(resp.Data)
	json.Unmarshal(dataJson, &userInfo)

	return userInfo.Token
	//
	//// 格式化输出
	//respJson, _ := json.MarshalIndent(resp, "", "    ")
	//fmt.Println(string(respJson))
}
func TestUploadAvatar(t *testing.T) {
	token := setupLogin()
	route := router2.NewRouter()
	w := httptest.NewRecorder()
	filePath := "../../.env"
	// 创建上传文件的 multipart.FileHeader
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 获取文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		// 处理错误
		panic(err)
	}

	// 获取文件大小
	fileSize := fileInfo.Size()

	header := &multipart.FileHeader{
		Filename: fileInfo.Name(),
		Size:     fileSize,
	}

	// 模拟 multipart/form-data 请求
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", header.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}
	writer.Close()

	req, _ := http.NewRequest("POST", "/upload/avatar", body)
	req.Header.Set("token", token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	route.ServeHTTP(w, req)
	extend.PrintResponseJson(w.Body.String())
}

func TestUploadCommentFile(t *testing.T) {

}

func TestUploadTaskFile(t *testing.T) {

}
