package controller

import (
	"fmt"
	"net/http"
	"web/bookstore/dao"
	"web/bookstore/model"
)

// AddBookToCar t添加图书到购物车
func AddBookToCar(w http.ResponseWriter, r *http.Request) {
	// 获取图书ID
	bookID := r.FormValue("bookID")
	// 获取图书信息
	book, _ := dao.GetBookByID(bookID)
	car := &model.Car{}
	carItem := &model.CarItem{}
	// 查询是否已经有这个购物车
	// 根据UserID查询
	// 要先登录的状态
	f, session := dao.IsLogin(r)
	// 如果没有登录转到登录页面
	if !f {
		// 因为是使用ajax请求所以没办法直接重定向
		// 将响应传入ajax中进行重定向
		w.Write([]byte("login"))
	} else {
		car, _ = dao.GetCarByUserID(session.UserID)
		if car.CarID != "" {
			// 修改本用户购物车的信息,同时还要更新catItem
			car.CarGodsCount++
			car.CarSumPrice += book.Price
			// update
			dao.UpdateCar(car)
			// 如果该书已在catItem,修改即可，否则插入
			// 同时判断该书是否属于该用户，该书是否已经添加进购物车
			carItem, _ = dao.GetCarItemByBookID(bookID, car.CarID)
			carItem.Book = book
			if carItem.ID > 0 {
				carItem.GodsCount++
				carItem.SumPrice += book.Price
				fmt.Println(carItem)
				dao.UpdateCarItem(carItem)
			} else {
				carItem.GodsCount = 1
				carItem.CarID = car.CarID
				dao.AddCarItem(carItem)
			}
		} else {
			// 创建新的购物车
			car.CarID = session.SessionID
			car.CarGodsCount = 1
			car.CarSumPrice = book.Price
			car.UserID = session.UserID
			dao.AddCar(car)
			// 并且添加新的carItem
			carItem.GodsCount = 1
			carItem.Book = book
			carItem.CarID = car.CarID
			dao.AddCarItem(carItem)
		}
	}
}
