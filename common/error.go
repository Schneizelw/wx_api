package common

const (
	ErrNumSuccess = 0
)

func ErrMsg(errNum int32) string {
	translate := map[int32]string{
		ErrNumSuccess: "ok",
	}
	res := translate[errNum]
	return res
}
