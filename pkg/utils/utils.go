package utils

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"reflect"
)

// 结构体转映射
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// GetUUID 获取uuid
func GetUUID() string {
	return uuid.NewV4().String()
}

// GetMD5 获取md5
func GetMD5(str string) string {
	h := md5.Sum([]byte(str))
	m := fmt.Sprintf("%x", h)
	return m
}

// Md5Compare 比较md5
func Md5Compare(str string, md5str string) bool {
	return GetMD5(str) == md5str
}

// PasswordEncrypt 密码加密, 加密后的密码为md5(md5(password)+salt)
func PasswordEncrypt(password string) string {
	return GetMD5(GetMD5(password) + "yueqing2617")
}

// PasswordCompare 密码比较
func PasswordCompare(password string, md5str string) bool {
	return PasswordEncrypt(password) == md5str
}

// ReadNetFile 读取网络文件
func ReadNetFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// ParseTxt 将 ReadNetFile 读取的文件转换为txt
func ParseTxt(body []byte) string {
	return string(body)
}
