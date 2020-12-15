package tools

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
)

// 完美解析Post请求json参数
// 支持：[x-www-from-urlencoded]、raw[json]
func ReqPostBody(c *context.Context, key string) (value string, err error) {
	err1 := c.Request.ParseForm()
	if err1 != nil {
		return "", err1
	}
	param := c.Request.Form[key]
	if param != nil && len(param) != 0 {
		fmt.Printf("param::%v\n", param)
		return param[0], nil
	} else {
		decoder := json.NewDecoder(c.Request.Body)
		// 用于存放参数key=value数据
		var params map[string]string
		// 解析参数 存入map
		err2 := decoder.Decode(&params)
		if err2 != nil {
			return "", err2
		}
		return params[key], nil
	}
}
