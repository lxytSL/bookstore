package controller

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"web/bookstore/dao"
	"web/bookstore/model"
)

// MainHandle 解析首页
func MainHandle(w http.ResponseWriter, r *http.Request) {
	// 获取当前页
	curPage := r.FormValue("curPage")
	minPrice := r.FormValue("minPrice")
	maxPrice := r.FormValue("maxPrice")
	if curPage == "" {
		curPage = "1"
	}
	var page *model.Page
	if minPrice != "" || maxPrice != "" {
		page, _ = dao.GetBookByParse(minPrice, maxPrice, curPage)
		page.MaxPrice = maxPrice
		page.MinPrice = minPrice
	} else {
		page, _ = dao.PageGetBook(curPage)
	}
	// GETcookie
	var session *model.Session
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		// 获取cookie的value
		cookieValue := cookie.Value
		// 查询session结构
		session, _ = dao.GetSession(cookieValue)
		page.Ses = session
	}
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, page)
}

// PageGetBook 获取分页图书
func PageGetBook(w http.ResponseWriter, r *http.Request) {
	// 获取当前页
	curPage := r.FormValue("curPage")
	page, _ := dao.PageGetBook(curPage)
	t, _ := template.ParseFiles("views/pages/manager/book_manager.html")
	t.Execute(w, page)
}

// AddBook 添加图书
func AddBook(w http.ResponseWriter, r *http.Request) {
	// 获取提交的图书信息
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")

	fPrice, _ := strconv.ParseFloat(price, 64)
	fsales, _ := strconv.Atoi(sales)
	fstock, _ := strconv.Atoi(stock)
	book := &model.BookInfo{
		Title:   title,
		Author:  author,
		Price:   float32(fPrice),
		Sales:   fsales,
		Stock:   fstock,
		ImgPath: "/static/img/default.jpg",
	}
	err := dao.AddBook(book)
	if err != nil {
		log.Fatal(err)
	}
	// 调用GetBooks查一遍
	PageGetBook(w, r)
}

// DeleteBook 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// 获取BookID
	bookID := r.FormValue("bookId")
	err := dao.DeleteBook(bookID)
	if err != nil {
		panic(err)
	}
	// 调用GetBooks查一遍
	PageGetBook(w, r)
}

// GetBookByID 通过ID获取所有信息
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	// 获取bookid
	bookID := r.FormValue("bookId")
	// 获取要修改的图书信息
	book, _ := dao.GetBookByID(bookID)
	// 转到更新页面
	t, _ := template.ParseFiles("views/pages/manager/book_update.html")
	t.Execute(w, book)

}

// UpdateBook 更新图书
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := &model.BookInfo{}
	book.ID, _ = strconv.Atoi(r.PostFormValue("bookID"))

	book.Title = r.PostFormValue("title")
	book.Author = r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	fprice, _ := strconv.ParseFloat(price, 64)
	book.Price = float32(fprice)
	book.Sales, _ = strconv.Atoi(sales)
	book.Stock, _ = strconv.Atoi(stock)
	err := dao.UpdateBook(book)
	if err != nil {
		log.Fatal(err)
	}
	// 调用GetBooks查一遍
	PageGetBook(w, r)
}

// GetBookByParse 通过价格区间查询
func GetBookByParse(w http.ResponseWriter, r *http.Request) {
	// 获取当前页，获取价格区间
	curPage := r.FormValue("curPage")
	minPrice := r.FormValue("minPrice")
	maxPrice := r.FormValue("maxPrice")
	page, _ := dao.GetBookByParse(minPrice, maxPrice, curPage)
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, page)
}
