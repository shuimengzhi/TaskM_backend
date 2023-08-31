package user_controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"taskm/core"
	useriostruct "taskm/io_struct/user"
	"taskm/model"
	router2 "taskm/router"
	"testing"
)

func TestRegister(t *testing.T) {
	envPath, _ := os.Getwd()
	envPath = envPath + "/../../.env"
	core.LoadCore(envPath)

	router := router2.NewRouter()
	w := httptest.NewRecorder()
	request := useriostruct.RegisterRequest{Account: "admin", NickName: "admin", RoleType: model.USER_ROLE_TYPE_ADMIN, Avatar: ""}
	infoJson, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/user/register", strings.NewReader(string(infoJson)))
	router.ServeHTTP(w, req)
	fmt.Println(fmt.Sprintf("%#v", w.Body.String()))
}

func TestLogin(t *testing.T) {
	envPath, _ := os.Getwd()
	envPath = envPath + "/../../.env"
	core.LoadCore(envPath)

	router := router2.NewRouter()
	w := httptest.NewRecorder()
	request := useriostruct.LoginRequest{UserName: "admin", Password: "123456"}
	infoJson, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(string(infoJson)))
	router.ServeHTTP(w, req)
	fmt.Println(fmt.Sprintf("%#v", w.Body.String()))
}
