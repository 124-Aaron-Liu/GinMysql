package Model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	Username   string
	Department string
	Created    string
}

const (
	UserName     string = "root"
	Password     string = "123456"
	Addr         string = "ginmysql_mysql_1"
	Port         int    = 3306
	Database     string = "test"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func Create(param UserInfo) error {
	//組合sql連線字串
	// "root:123456@tcp(ginmysql_mysql_1)/test"
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
	db, err := sql.Open("mysql", conn)
	defer db.Close()
	checkErr(err)

	//插入資料
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec(param.Username, param.Department, param.Created)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	return nil
}

func GetUserInfo(id int) (UserInfo, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(ginmysql_mysql_1)/test?charset=utf8")
	defer db.Close()

	checkErr(err)
	userInfo := UserInfo{}
	row := db.QueryRow("select username, department, created FROM userinfo WHERE uid=?", id)

	//Scan對應的欄位與select語法的欄位順序一致
	if err := row.Scan(&userInfo.Username, &userInfo.Department, &userInfo.Created); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return UserInfo{}, err
	}

	return userInfo, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
