package services

const SUCCESS = 0
const FAIL = 100

type ResultService struct {
	Code int
	Data interface{}
	Msg  string
}

func FailResponse(msg string, data interface{}) ResultService {
	return ResultService{Code: FAIL, Data: data, Msg: msg}
}
func SuccessResponse(data interface{}, msg string) ResultService {
	return ResultService{Code: SUCCESS, Data: data, Msg: msg}
}
