package project_service

import (
	"taskm/io_struct/project"
	useriostruct "taskm/io_struct/user"
	"taskm/model"
	"taskm/services"
	"time"
)

func ProjectCreate(request project.CreateRequest, userInfo *useriostruct.LoginResponse) services.ResultService {

	projectModel := model.Project{
		PName:           request.ProjectName,
		PStatus:         model.PROJECT_STATUS_NORMAL,
		PCreateUserID:   userInfo.UID,
		PCreateUserName: userInfo.NickName,
		PCreateTime:     int32(time.Now().Unix()),
		PUpdateTime:     int32(time.Now().Unix()),
	}
	result := model.DB.Create(&projectModel)
	if result.Error != nil {
		return services.ResultService{Code: services.FAIL, Msg: result.Error.Error()}
	}
	return services.ResultService{Code: services.SUCCESS}
}
