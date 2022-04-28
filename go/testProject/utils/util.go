package utils

import (
	"reflect"
	"time"
)

func DateTime() string  {
	return time.Now().Format("2006-01-02 15:04:05")
}

// StructToMap 结构体传map
func StructToMap(s interface{}) interface{}  {
	//如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	newMap := make(map[string]interface{})  // map 一定要初始化才能使用
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	for i := 0; i < t.NumField(); i++ {
		newMap[t.Field(i).Name] = v.Field(i).Interface()
	}
	return newMap
}