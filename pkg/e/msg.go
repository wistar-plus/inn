package e

import "errors"

var Msgs = map[int]string{
	SUCCESS:        "ok",
	Fail:           "fail",
	INVALID_PARAMS: "无效参数",
}

var (
	ERROR_EXIST         = errors.New("资源已存在")
	ERROR_NOTEXIST      = errors.New("资源不存在")
	ERROR_WRONGPASSWORD = errors.New("错误密码")
	ERROR_DBERROR       = errors.New("查询数据库出错")
)

func GetMsg(code int) string {
	msg, ok := Msgs[code]
	if ok {
		return msg
	}

	return Msgs[Fail]
}
