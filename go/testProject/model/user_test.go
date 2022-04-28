package model

import (
	"testing"
	"time"
)



func TestCon(t *testing.T) {
	InitDB()

	//r := strconv.Itoa(rand.Intn(1000))
	userModel := User{Name: "Jinzhu2" , Email: "test@test.com", Age: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	db.Create(&userModel)
	//Insert(userModel)
	/*user := User{Name: "Jinzhu", Email:"test@test.com",Age: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	result := db.Create(&user) // 通过数据的指针来创建*/
	//var user User
	/*var users []User
	db.Where("name like ?","%j%").Find(&users)
	fmt.Println(users)*/
}

func TestInsert(t *testing.T) {
	InitDB()
	u := User{Name: "Jinzhu" , Email: "test@test.com", Age: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	count,_ := u.Insert()
	println(count)
}
