package db

import (
	mydb "filestore-server/db/mysql"
	"fmt"
)

type User struct {
	Username string
	Email string
	Phone string
	SignUpAt string
	LastActiveAt string
	status int
}

//UserSignUp: 通过用户名及密码完成user表的注册操作
func UserSignUp(username, passwd string) bool {
	stmt, err := mydb.DBConn().Prepare("insert ignore into tbl_user (`user_name`,`user_pwd`) values (?,?)")
	if err != nil {
		fmt.Println("Failed to insert, err:", err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, passwd)
	if err != nil {
		fmt.Println("Failed to insert, err:", err.Error())
		return false
	}
	rowsAffected, err := ret.RowsAffected()
	if err != nil {
		return false
	}
	if rowsAffected <= 0 {
		return false
	}
	return true
}

//UserSignIn: 判断账号密码是否正确
func UserSignIn(username, encpwd string) bool {
	stmt, err :=mydb.DBConn().Prepare("select  *from  tbl_user where user_name = ? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}else if rows == nil {
		fmt.Printf("user %s not found\n", username)
		return false
	}

	pRows := mydb.ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd {
		return true
	}
	return false
}

//UpdateToken: 更新token
func UpdateToken(username, token string) bool {
	stmt, err := mydb.DBConn().Prepare("replace into tbl_user_token(`user_name`,`user_token`) values (?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

//GetUserInfo: 获取用户信息
func GetUserInfo(username string) (User, error) {
	user := User{}

	stmt, err := mydb.DBConn().Prepare("select user_name, signup_at from tbl_user where user_name = ? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignUpAt)
	if err != nil {
		return user, err
	}
	return user, nil
}