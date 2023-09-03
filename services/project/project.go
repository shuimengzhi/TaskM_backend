package project_service

import (
	"taskm/io_struct/project"
	useriostruct "taskm/io_struct/user"
	"taskm/model"
	"taskm/services"
	"time"
)

func ProjectCreate(request project_io_struct.CreateRequest, userInfo *useriostruct.LoginResponse) services.ResultService {

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

func ProjectList(request project_io_struct.ListRequest, userInfo *useriostruct.LoginResponse) services.ResultService {
	var response []project_io_struct.ProjectListItem
	db := model.DB.Model(&model.Project{}).Joins("join tm_project_user on p_id = pu_project_id and pu_user_id = ?",
		userInfo.UID)
	offset := (request.Page - 1) * request.Size
	if request.ProjectName != "" {
		db = db.Where("p_name = ?", request.ProjectName)
	}
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return services.ResultService{Code: services.FAIL, Msg: err.Error()}
	}
	result := db.Select("p_id,p_name").
		Limit(request.Size).Offset(offset).Order("p_create_time desc").Scan(&response)
	if result.Error != nil {
		return services.ResultService{Code: services.FAIL, Msg: result.Error.Error()}
	}
	return services.ResultService{Code: services.SUCCESS,
		Data: project_io_struct.ProjectListResponse{List: response, Count: count}, Msg: "success"}
}
