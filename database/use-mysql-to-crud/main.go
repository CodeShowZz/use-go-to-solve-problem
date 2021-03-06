package main

//go对mysql执行crud
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/learning")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("huangjunlin", "finance", "2020-12-01")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	//更新
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("chenling", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows,err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for(rows.Next()) {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid,&username,&department,&created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt,err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res,err = stmt.Exec(id)
	checkErr(err)

	affect,err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
