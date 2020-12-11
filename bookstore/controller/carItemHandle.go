package controller

import (
	"net/http"
	"text/template"
	"web/bookstore/dao"
	"web/bookstore/model"
)

// ShowCarItem 添加图书到购物车
func ShowCarItem(w http.ResponseWriter, r *http.Request) {
	// 获取sessionID
	cookie, _ := r.Cookie("user")
	cookieValue := cookie.Value
	session, _ := dao.GetSession(cookieValue)
	page := &model.Page{}
	page.Ses = session
	car := &model.Car{}
	// 根据User_id查找当前购物车
	car, _ = dao.GetCarByUserID(session.UserID)
	// 根据car_id 查找当前用户购物车中选项
	var carItems []*model.CarItem
	carItems, _ = dao.GetCarItemByCarID(car.CarID)
	car.CarItems = carItems
	page.Car = car
	t, _ := template.ParseFiles("views/pages/cart/cart.html")
	t.Execute(w, page)
}

// DeleteCarItems 删除购物车的图书项
func DeleteCarItems(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookID")
	carID := r.FormValue("carID")
	// 查找当前要删除图书的数量和总价格
	carItem, _ := dao.GetCarItemByBookID(bookID, carID)
	// 删除
	dao.DeleteCarItem(bookID, carID)
	// 根据当前的用户的购物车
	car, _ := dao.GetCarByCarID(carID)
	// 更新购物车
	car.CarGodsCount -= carItem.GodsCount
	car.CarSumPrice -= carItem.SumPrice
	dao.UpdateCar(car)

	// 再查一遍
	ShowCarItem(w, r)

}
