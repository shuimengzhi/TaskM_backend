package extend

import (
	"crypto/md5"
	"encoding/hex"
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

// VarDump 打印数据
func VarDump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}

func GetMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
