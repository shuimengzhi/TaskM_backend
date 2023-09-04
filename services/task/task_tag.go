package task_service

import (
	taskiostruct "taskm/io_struct/task"
	"taskm/model"
	"taskm/services"
)

func TaskTagCreate(request taskiostruct.TaskTagCreateRequest) services.ResultService {

	//创建任务标签
	taskTagModel := model.TaskTag{
		TtName:      request.TagName,
		TtSort:      request.Sort,
		TtColor:     request.Color,
		TtProjectID: request.ProjectId,
	}

	if err := model.DB.Create(&taskTagModel).Error; err != nil {
		return services.ResultService{Code: services.FAIL, Msg: err.Error() + "  TaskTagCreate:1"}
	}

	return services.ResultService{Code: services.SUCCESS}
}
