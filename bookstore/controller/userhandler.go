package controller

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
	"web/bookstore/dao"
	"web/bookstore/model"
	"web/bookstore/utils"
)

//LoginHandle 用户登录处理
func LoginHandle(w http.ResponseWriter, r *http.Request) {
	// 获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	// 用方法校验
	user, _ := dao.CheckUser(username, password)
	// 如果第一次登录
	if user.ID > 0 {
		// 判断是否已经登录过
		// 不用set sessiom
		if f, _ := dao.IsLogin(r); !f {
			// 创建session
			uuid := utils.CreateUUID()
			session := &model.Session{
				SessionID: uuid,
				UserName:  user.UserName,
				UserID:    user.ID,
			}
			dao.AddSession(session)
			// 设置cookie
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
		}
		// 用户名密码正确,解析首页
		t, _ := template.ParseFiles("views/pages/user/login_success.html")
		t.Execute(w, user)
	} else {
		t, _ := template.ParseFiles("views/pages/user/login.html")
		t.Execute(w, "用户名或密码不正确")
	}
}

//Register 用户注册处理
func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	fmt.Println(username, password, email)
	// 判断是否存在用户名
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		// 说明存在重复，重新注册
		t, _ := template.ParseFiles("views/pages/user/regist.html")
		t.Execute(w, "用户名已存在!")
	} else {
		// 可以注册，信息进行保存
		dao.SaveUser(username, password, email)

		// 保存成功，定位登录界面
		t, _ := template.ParseFiles("views/pages/user/login.html")
		t.Execute(w, nil)
	}
}

// CheckUserName 处理ajax传过来的url，检查用户名是否存在
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		// 说明存在重复
		w.Write([]byte("用户名已存在！"))
	} else {
		// 传入样式，绿色字体显示
		// 前端必须使用html
		fmt.Fprintf(w, "<font style='color:green'>用户名可用！</font>")

	}
}

// LoginOut 注销
func LoginOut(w http.ResponseWriter, r *http.Request) {
	// 注销，删除cookie
	cookie, _ := r.Cookie("user")
	cookieValue := cookie.Value
	// 删除session
	err := dao.DeleteSession(cookieValue)
	if err != nil {
		log.Fatal(err)
	}
	// 删除cookie
	rc := http.Cookie{
		Name:    "user",
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
	}
	http.SetCookie(w, &rc)
	// 注销之后去首页
	MainHandle(w, r)
}
