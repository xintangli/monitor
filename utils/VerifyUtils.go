package utils

import (
	"crypto/md5"
	"strings"
	derrors "github.com/xintangli/monitor/error"
)

func LoginVerify(memId string, timestamp string, sign string, key string) (bool, error)  {
	str := memId + timestamp + key
	localSign := MD5(str)
	if strings.EqualFold(string(localSign), sign) {
		return true, nil
	}else{
		return false, derrors.New("1099","sign verify fail!")
	}
}

func MD5(key string) string {
	md5 := md5.New()
	md5.Write([]byte(key))
	return string(md5.Sum(nil))
}