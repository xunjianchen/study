package main

import (
	"encoding/json"
	"fmt"
)

//定义一个简单的结构体 Person
type Person struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sex      float32 `json:"sex"`
	Hobby    string  `json:"hobby"`
}

//写一个 testStruct()结构体的序列化方法
func testStruct() {
	person := Person{
		Name:     "小崽子",
		Age:      50,
		Birthday: "2019-09-27",
		Sex:      1000.01,
		Hobby:    "泡妞",
	}

	// 将Monster结构体序列化
	data, err := json.Marshal(&person)
	if err != nil {
		fmt.Printf("序列化错误 err is %v", err)
	}
	//输出序列化结果
	fmt.Printf("person序列化后 = %v\n", string(data))
	//反序列化
	person2 := Person{}
	json.Unmarshal(data,&person2)
	fmt.Println(person2)
}
func testMap()  {
	b := []byte(`{"age":1.3,"name":"5lmh.com","marry":false}`)

	// 声明接口
	var i map[string]interface{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n",i)
	fmt.Printf("%#v\n",i)
	fmt.Printf("%#v\n", i["age"])
	// 自动转到map
	fmt.Println(i)
	// 可以判断类型
	/*m := i.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case float64:
			fmt.Println(k, "是float64类型", vv)
		case string:
			fmt.Println(k, "是string类型", vv)
		default:
			fmt.Println("其他")
		}

	}*/
}
func main()  {
	testStruct()

}