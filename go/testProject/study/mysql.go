package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person2 struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}


var Db *sqlx.DB

func init() {
	//defer Db.Close()  // 注意这行代码要写在上面err判断的下面
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database

}

func insertTest()  {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}

func updateTest()  {
	res, err := Db.Exec("update person set username=? where user_id=?", "stu00034444", 5)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ",err)
	}
	fmt.Println("update succ:",row)
}

func selectTest()  {
	var person []Person2
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 5)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Printf("%+v\n",person)
	fmt.Printf("%T\n",person)

	fmt.Println("select succ:", person)
}

func deleteTest() {
	res, err := Db.Exec("delete from person where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	row,err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ",err)
	}

	fmt.Println("delete succ: ",row)
}

func main() {
	//defer Db.Close()
	deleteTest()

}