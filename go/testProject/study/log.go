package main

import (
	"fmt"
	"log"
	"os"
)

func logTest(data interface{},remark string,prefix string)  {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("["+prefix+"]")
	log.Println(remark,data)
}


func main()  {

	ss := map[string]string{"aa":"cc"}
	logTest(ss,"这是说明","神仙")


}
