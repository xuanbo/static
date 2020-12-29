package api

// Result API统一返回
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 成功
func Success(data interface{}) *Result {
	return &Result{Code: 0, Data: data}
}

// Fail 失败
func Fail(msg string) *Result {
	return &Result{Code: 1, Msg: msg}
}
