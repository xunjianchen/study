package main

import (
	"math/rand"
	"strconv"
	"testProject/model"
	"time"
)

func main() {
	model.InitDB()
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	r := strconv.Itoa(rand.Intn(1000))
	u := model.User{Name: "Jinzhu"+ r, Email: "test" + r + "@test.com", Age: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	count,_ := u.Insert()
	println(count)
}
