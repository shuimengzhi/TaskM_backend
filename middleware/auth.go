package middleware

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"runtime"
	"strconv"
	"taskm/cache"
	"taskm/enum"
	commoniostruct "taskm/io_struct/common"
	useriostruct "taskm/io_struct/user"
	"time"
)

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		userInfo, err := cache.Instance.Get(context.Background(), token).Result()
		if err != nil {
			switch {
			case err == redis.Nil:
				c.JSON(200, commoniostruct.Response{Code: enum.CodeParamErr, Msg: "登陆失败"})
			case err != nil:
				c.JSON(200, commoniostruct.Response{Code: enum.CodeParamErr, Msg: err.Error()})
			case userInfo == "":
				c.JSON(200, commoniostruct.Response{Code: enum.CodeParamErr, Msg: "登陆失败"})
			}
			c.Abort()
			return
		}
		var resp useriostruct.LoginResponse
		err = json.Unmarshal([]byte(userInfo), &resp)
		if err != nil {
			c.JSON(200, commoniostruct.Response{Code: enum.CodeParamErr, Msg: err.Error()})
			c.Abort()
			return
		}

		//请求就重新计算过期时间
		setErr := cache.Instance.Set(context.Background(), token, userInfo, 2*time.Hour).Err()
		if setErr != nil {
			_, file, line, ok := runtime.Caller(0)
			if ok {
				c.JSON(200, commoniostruct.Response{Code: enum.CodeParamErr, Msg: setErr.Error() + "  " + file + ":" + strconv.Itoa(line)})
				c.Abort()
				return
			}
			c.JSON(200, commoniostruct.Response{Code: enum.CodeParamErr, Msg: setErr.Error()})
			c.Abort()
			return
		}

		c.Set("user_info", &resp)
		c.Next()
		return
	}
}
