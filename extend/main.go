package extend

import (
	"encoding/json"
	"fmt"
	commoniostruct "taskm/io_struct/common"
)

// PrintResponseJson 打印响应的json内容
func PrintResponseJson(data string) {
	var resp commoniostruct.Response
	json.Unmarshal([]byte(data), &resp)

	// 格式化输出
	respJson, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(respJson))
}
