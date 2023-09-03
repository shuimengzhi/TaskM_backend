package project_service

import (
	"taskm/io_struct/project"
	useriostruct "taskm/io_struct/user"
	"taskm/model"
	"taskm/services"
	"time"
)

func ProjectCreate(request project_io_struct.CreateRequest, userInfo *useriostruct.LoginResponse) services.ResultService {
	tx := model.DB.Begin()
	//创建项目
	projectModel := model.Project{
		PName:           request.ProjectName,
		PStatus:         model.PROJECT_STATUS_NORMAL,
		PBelongUserID:   userInfo.UID,
		PBelongUserName: userInfo.NickName,
		PCreateTime:     int32(time.Now().Unix()),
		PUpdateTime:     int32(time.Now().Unix()),
	}

	if err := tx.Create(&projectModel).Error; err != nil {
		tx.Rollback()
		return services.ResultService{Code: services.FAIL, Msg: err.Error()}
	}
	// 创建项目个人关系表
	projectUserModel := model.ProjectUser{PuProjectID: projectModel.PID, PuUserID: userInfo.UID, PuCreateTime: int32(time.Now().Unix())}
	if err := tx.Create(&projectUserModel).Error; err != nil {
		tx.Rollback()
		return services.ResultService{Code: services.FAIL, Msg: err.Error()}
	}
	tx.Commit()
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

func ProjectUpdate(request project_io_struct.UpdateRequest, userInfo *useriostruct.LoginResponse) services.ResultService {
	var projectModel model.Project
	var newBelongUser model.User
	if err := model.DB.Where("p_id = ? and p_belong_user_id = ?", request.PID, userInfo.UID).First(&projectModel).Error; err != nil {
		return services.ResultService{Code: services.FAIL, Msg: err.Error() + " ProjectUpdate:1"}
	}
	if projectModel.PID == 0 {
		return services.ResultService{Code: services.FAIL, Msg: "找不到要更改的项目"}
	}
	if request.PBelongUserID != 0 {
		if err := model.DB.Model(model.User{}).Where("u_id = ? and u_status = ?", request.PBelongUserID, model.USER_STATUS_NORMAL).
			Select("u_id,u_nick_name").First(&newBelongUser).Error; err != nil {
			return services.ResultService{Code: services.FAIL, Msg: err.Error() + " ProjectUpdate:2"}
		}
		if newBelongUser.UID == 0 {
			return services.ResultService{Code: services.FAIL, Msg: "找不到新的负责人"}
		}
	}
	// 更新模型的

	if err := model.DB.Model(model.Project{}).Where("p_id = ?", request.PID).
		Updates(model.Project{PName: request.PName, PStatus: request.PStatus, PBelongUserID: newBelongUser.UID,
			PBelongUserName: newBelongUser.UNickName, PUpdateTime: int32(time.Now().Unix())}).Error; err != nil {
		return services.ResultService{Code: services.FAIL, Msg: err.Error() + " ProjectUpdate:3"}
	}
	return services.ResultService{Code: services.SUCCESS}
}
