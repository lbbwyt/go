//如何连接mysql以及常用的增删改查操作
package webserver

import (
	"database/sql"
	"fmt"
	"strconv"
	//不应该直接使用驱动所提供的方法, 而是应该使用 sql.DB, 因此在导入 mysql 驱动时, 这里使用了匿名导入的方式(在包路径前添加 _), 当导入了一个数据库驱动后,
	//此驱动会自行初始化并注册自己到Golang的database/sql上下文中, 因此我们就可以通过 database/sql 包提供的方法访问数据库了.
	_ "github.com/go-sql-driver/mysql"
)

var connStr string = "root:123456@tcp(127.0.0.1:3306)/world"

//获取mysql数据库连接
func GetConn(str string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", connStr)
	return
}

//异常处理
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//mysql增删改查
func Ddl() {
	db, err := GetConn(connStr)
	checkErr(err)
	defer db.Close()
	//查询
	{
		rows, err := db.Query("select id, name from city order by id limit 5 ")
		checkErr(err)
		for rows.Next() {
			var id int
			var cityName string
			err = rows.Scan(&id, &cityName)
			checkErr(err)
			fmt.Println(strconv.Itoa(id) + ":" + cityName)
		}
		fmt.Println("**********************************")
	}

	//修改
	{
		//Prepare函数返回与当前连接相关的执行Sql语句的准备状态，可以进行查询、删除等操作。
		stmt, err := db.Prepare("update city set name=? where id=?")
		checkErr(err)
		//Stmt是一种准备好的状态，和Conn相关联，而且只能应用于一个goroutine中，不能应用于多个goroutine
		res, err := stmt.Exec("beijing", 1)
		affect, err := res.RowsAffected()
		checkErr(err)
		fmt.Println(affect)
		fmt.Println("**********************************")
	}

	//插入 和 删除
	{
		stmt, err := db.Prepare("INSERT city SET name=?,countrycode=?")
		checkErr(err)

		res, err := stmt.Exec("shanghai", "AFG")
		checkErr(err)
		//LastInsertId函数返回由数据库执行插入操作得到的自增ID号。
		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)
		fmt.Println("**********************************")

		//删除
		stmt, err = db.Prepare("delete from city where id=?")
		checkErr(err)
		res, err = stmt.Exec(id)
		checkErr(err)
		//		RowsAffected函数返回query等操作影响的数据条目数。
		affect, err := res.RowsAffected()
		checkErr(err)
		fmt.Println("删除行数为：" + strconv.FormatInt(affect, 10) + ",id 为" + strconv.FormatInt(id, 10))

		//#string到int
		//int,err:=strconv.Atoi(string)
		//#string到int64
		//int64, err := strconv.ParseInt(string, 10, 64)
		//#int到string
		//string:=strconv.Itoa(int)
		//#int64到string
		//string:=strconv.FormatInt(int64,10)

	}

}
