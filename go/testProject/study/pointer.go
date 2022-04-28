package main

import "fmt"


type OrderInfo struct {
	Id string
	Price float64
	Status int
}

// GetOne 形参 orderInfo 存储的是一个地址p，而 *orderInfo 表示访问p指向的内存中的值，改变这个值就改变了实参的实际内容，达到了赋值的效果。
func GetOne(orderInfo *OrderInfo) {
	/*
	错误写法
	orderInfo = &OrderInfo{
		Id: "aaaa",
		Price: 100.00,
		Status: 1,
	}
	*/

	*orderInfo = OrderInfo{
		Id: "aaaa",
		Price: 100.00,
		Status: 1,
	}
	fmt.Printf("%p\n",orderInfo)
	//fmt.Println("orderInfo: ",orderInfo)
}

func main() {
	var oi OrderInfo
	fmt.Printf("&oi:%p\n",&oi)
	GetOne(&oi)
	fmt.Println(oi)
}



