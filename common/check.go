package common

import (
	"sort"
	"strings"
	"wx_api/util"
)

type CheckSignatureOpt struct {
	Signature string
	Timestamp string
	Nonce     string
}

func CheckSignature(opt *CheckSignatureOpt) bool {
	params := []string{WxCheckServerToken, opt.Timestamp, opt.Nonce}
	sort.Sort(sort.StringSlice(params))
	sha1Sum := util.Sha1Sum(strings.Join(params, ""))
	if sha1Sum != opt.Signature {
		return false
	}
	return true
}
