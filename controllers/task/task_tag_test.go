package task_controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"taskm/core"
	commoniostruct "taskm/io_struct/common"
	taskiostruct "taskm/io_struct/task"
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
func TestTaskTagCreate(t *testing.T) {
	token := setupLogin()
	envPath, _ := os.Getwd()
	envPath = envPath + "/../../.env"
	core.LoadCore(envPath)

	router := router2.NewRouter()
	w := httptest.NewRecorder()
	request := taskiostruct.TaskTagCreateRequest{ProjectId: 1, TagName: "待处理", Color: "#E0457B"}
	infoJson, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/task_tag/create", strings.NewReader(string(infoJson)))
	req.Header.Set("token", token)
	router.ServeHTTP(w, req)
	fmt.Println(fmt.Sprintf("%#v", w.Body.String()))
}
