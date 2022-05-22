package utils

import (
	"math/rand"
	"strconv"
	"time"
)

// GetRandomString 获取随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandomInt 获取随机数
func GetRandomInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

// GetIntervalString 获取min到max之间的随机字符串
func GetIntervalString(min, max int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return GetRandomString(r.Intn(max-min) + min)
}

// StringToInt 字符串转int
func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// StringToInt64 字符串转int64
func StringToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}
