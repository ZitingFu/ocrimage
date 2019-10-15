package string_tools

import (
	"math/rand"
	"time"
)

/**
* 得到随机字符串
*
* @param strType 要生成的随机字符串的类型 支持alnum numeric nozero alpha ，分别代表纯字母、字母数字、纯数字包含0、纯数字不包含0
* @param length 生成字符串的长度
 */
func RandomStr(strType string, length int) string {
	alphaStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alnumStr := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numericStr := "0123456789"
	nozeroStr := "123456789"

	pool := ""
	switch strType {
	case "alpha":
		pool = alphaStr
	case "alnum":
		pool = alnumStr
	case "numeric":
		pool = numericStr
	case "nozero":
		pool = nozeroStr
	}

	resultStr := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= length; i++ {
		//字符串随机键值
		index := r.Intn(len(pool) - 1)
		if index == (len(pool) - 1) {
			resultStr += pool[index:]
		} else {
			resultStr += pool[index : index+1]
		}

	}
	return resultStr
}
