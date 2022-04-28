package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

var db *gorm.DB
var err error

type User2 struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


func main() {
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("db connect error")
	}
	defer db.Close() //延时调用函数
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User2{})

	r := gin.Default()
	r.GET("/users", index)          //获取所有用户
	r.GET("/users/:id", detail)       //根据id获取用户
	r.POST("/users", insert)        //保存新用户
	r.PUT("/users/:id", update)     //根据id更新用户
	r.DELETE("/users/:id", delete) //根据id删除用户
	_ = r.Run()
}

func index(c *gin.Context) {
	var users []User2
	db.Find(&users)
	c.JSON(200, users)
}

func detail(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User2
	db.First(&user, id)
	if user.ID == 0 {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	}
	c.JSON(200, user)
}

func insert(c *gin.Context) {
	var user User2
	_ = c.BindJSON(&user) //绑定一个请求主体到一个类型
	db.Create(&user)
	c.JSON(200, user)
}

//此更新依赖body中的数据，id值不能错
func update(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User2
	db.First(&user, id)
	if user.ID == 0 {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	} else {
		_ = c.BindJSON(&user)
		intId,_ := strconv.Atoi(id)
		user.ID = uint(intId)
		//db.Save(&user)
		db.Model(&user).UpdateColumns(User2{Name: user.Name, Email: user.Email})
		c.JSON(200, user)
	}
}

func delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User2
	db.First(&user, id)
	if user.ID == 0 {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	} else {
		_ = c.BindJSON(&user)
		//db.Delete(&user)
		db.Where("id = ?",id).Delete(User2{})
		c.JSON(200, gin.H{"message": "delete success"})
	}
}
