package model

import (
	"fmt"
	"time"
)

/*func init() {
	// 创建表时添加后缀
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
}*/


type User struct {
	ID        uint
	Name      string
	Email     string
	Age       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}


func (User) tableName() string {
	return "users"
}

func (u User) IsExists() bool {
	var count int64
	db.Model(&User{}).Where("id = ?", u.ID).Count(&count)
	return count > 0
}
func (u User) Insert() (int64,error) {
	// 要有效地插入大量记录，请将一个 slice 传递给 Create 方法。
	//var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	fmt.Printf("%+v\n",u)

	result := db.Create(&u)
	return result.RowsAffected,result.Error
}

func (u User) Update() {
	if u.IsExists() {
		// Save 会保存所有的字段，即使字段是零值
		db.Where("id = ?", u.ID).Save(&u)
	} else {
		// insert
		u.Insert()
	}
}