package main

import (
	"fmt"
	"testProject/utils"
)

type PersonEntity struct {
	Id   string
	Name string
	Key  string
	Addr string
}

func main() {
	entity := PersonEntity{
		Id:   "11111",
		Name: "11111",
		Key:  "11111",
		Addr: "11111",
	}

	j := utils.StructToMap(entity)
	fmt.Printf("%#v",j)
}
