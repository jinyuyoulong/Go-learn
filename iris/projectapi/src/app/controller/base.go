package controller

type ApiJson struct {
	Status bool        `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func ApiResult(status bool, object interface{}, msg string) (apijson *ApiJson) {
	// apijson 已经在返回值处 声明了，不用重复声明。
	apijson = &ApiJson{Status: status, Data: object, Msg: msg}
	return apijson
}
