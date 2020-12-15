package tools

import (
	"encoding/json"
	"net/http"
)

// 完美解析Post请求json参数
// 支持：[x-www-from-urlencoded]、raw[json]
// get请求获取参数：controller.GetString("key")
func ReqPostBody(c *http.Request, key string) (value string, err error) {
	err1 := c.ParseForm()
	if err1 != nil {
		return "", err1
	}
	param := c.Form.Get(key)
	if param != "" && len(param) != 0 {
		return param, nil
	} else {
		decoder := json.NewDecoder(c.Body)
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
