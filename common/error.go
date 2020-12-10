package common

const (
	ErrNumSuccess       = 0
	ErrNumMenuCreateErr = 1000
)

func ErrMsg(errNum int32) string {
	translate := map[int32]string{
		ErrNumSuccess:       "ok",
		ErrNumMenuCreateErr: "创建菜单失败",
	}
	res := translate[errNum]
	return res
}
