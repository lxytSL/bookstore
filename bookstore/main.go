package main

import (
	"net/http"
	"web/bookstore/controller"
)

func main() {
	// 设置处理静态资源
	// 当浏览器请求index.html中的静态资源时，会从views/static/目录下寻找
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	http.HandleFunc("/index", controller.MainHandle)
	http.HandleFunc("/login", controller.LoginHandle)
	http.HandleFunc("/loginOut", controller.LoginOut)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	// http.HandleFunc("/getBooks", controller.GetBooks)
	http.HandleFunc("/PageGetBook", controller.PageGetBook)
	http.HandleFunc("/addBook", controller.AddBook)
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	// 更新图书分两步
	http.HandleFunc("/getBookById", controller.GetBookByID)
	http.HandleFunc("/updateBook", controller.UpdateBook)
	// 价格区间查询处理器
	http.HandleFunc("GetBookByParse", controller.GetBookByParse)

	// 购物车handle
	http.HandleFunc("/addBookToCar", controller.AddBookToCar)
	http.HandleFunc("/showCarItem", controller.ShowCarItem)
	http.HandleFunc("/deleteCarItems", controller.DeleteCarItems)
	http.ListenAndServe(":8080", nil)
}
