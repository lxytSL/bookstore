package dao

import (
	"fmt"
	"testing"
	"web/bookstore/model"
)

func TestCarItem(t *testing.T) {
	// t.Run("开始测试GetCarByUserID", testAddCarItem)
	t.Run("开始测试GetCarItemByCarID", testGetCarItemByCarID)
}

func testAddCarItem(t *testing.T) {
	book := &model.BookInfo{
		ID:      1,
		Title:   "golang web",
		Author:  "aaa",
		Price:   100,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	carItem := &model.CarItem{
		GodsCount: 1,
		Book:      book,
		CarID:     "d2775574-a35c-44b1-4fcf-12af5e65fdf2",
	}
	fmt.Println(carItem.Book.Price)
	AddCarItem(carItem)
}

func testGetCarItemByCarID(t *testing.T) {
	car, _ := GetCarItemByCarID("975d7c6f-1142-4da2-689e-5670a3e19a56")
	fmt.Println(car)
}
