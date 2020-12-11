package dao

import (
	"fmt"
	"testing"
	"web/bookstore/model"
)

func TestBook(t *testing.T) {
	// t.Run("开始测试", testGetBooks)
	// t.Run("开始测试添加图书", testAddBooks)
	// t.Run("开始测试添加图书", testDeleteBook)
	// t.Run("开始测试分页", testPageGetBook)
	// t.Run("开始测试分页", testGetBookByParse)

}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()

	for k, v := range books {
		fmt.Printf("第%d图书信息%v\n", k+1, v)
	}

}

func testAddBooks(t *testing.T) {
	book := &model.BookInfo{
		Title:   "golang web",
		Author:  "aaa",
		Price:   100,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	AddBook(book)
}

func testDeleteBook(t *testing.T) {
	DeleteBook("34")
}

func testPageGetBook(t *testing.T) {
	page, _ := PageGetBook("2")
	books := page.Books
	for _, v := range books {
		fmt.Println(v)
	}
}

func testGetBookByParse(t *testing.T) {
	page, _ := GetBookByParse("50", "100", "1")
	books := page.Books
	for _, v := range books {
		fmt.Println(v)
	}
}
