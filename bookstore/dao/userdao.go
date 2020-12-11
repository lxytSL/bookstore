package dao

import (
	"net/http"
	"web/bookstore/model"
	"web/bookstore/utils"
)

// CheckUser 处理登录
func CheckUser(username, password string) (*model.User, error) {
	sql := "select * from users where username=? and password=?"
	row := utils.Db.QueryRow(sql, username, password)
	u := &model.User{}
	row.Scan(&u.ID, &u.UserName, &u.PassWord, &u.Email)
	return u, nil
}

// CheckUserName 处理注册
func CheckUserName(username string) (*model.User, error) {
	sql := "select * from users where username=?"
	row := utils.Db.QueryRow(sql, username)
	u := &model.User{}
	row.Scan(&u.ID, &u.UserName, &u.PassWord, &u.Email)
	return u, nil
}

// SaveUser 保存user
func SaveUser(username, password, email string) error {
	sql := "insert into users(username,password,email) values(?,?,?)"
	_, err := utils.Db.Exec(sql, username, password, email)
	if err != nil {
		return err
	}
	return nil
}

// IsLogin 判断是否已经登录
func IsLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user")
	if cookie == nil {
		return false, nil
	}
	cookieValue := cookie.Value
	// 得到Session
	session, _ := GetSession(cookieValue)
	if session.UserID > 0 {
		return true, session
	}
	return false, nil
}
