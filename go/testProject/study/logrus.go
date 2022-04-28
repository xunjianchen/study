package main

import (
	"fmt"
)


func main()  {
	/*logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})*/
	ss := [5]int{1,2,3,4}
	/*fmt.Printf("%v\n",ss)
	fmt.Printf("%+v\n",ss)
	fmt.Printf("%#v\n",ss)*/
	fmt.Printf("%T\n",ss)
	/*logrus.SetReportCaller(true)    //是否显示行号
	logrus.Error("info"," abc")
	logrus.Debug()
	fmt.Println(utils.DateTime())*/
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			//runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
	}


}
