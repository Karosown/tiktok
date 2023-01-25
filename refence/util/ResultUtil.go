package util

type ResultUtil struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// OkData 成功的数据返回
func OkData(data interface{}) ResultUtil {
	return ResultUtil{Code: 0, Data: data}
}

// OkMsg 成功的消息返回
func OkMsg(msg string) ResultUtil {
	return ResultUtil{Code: 0, Msg: msg}
}

// Fail 失败的消息返回
func Fail(msg string) ResultUtil {
	return ResultUtil{Code: -1, Msg: msg}
}
