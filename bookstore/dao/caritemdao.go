package dao

import (
	"web/bookstore/model"
	"web/bookstore/utils"
)

// AddCarItem 添加进购物车
func AddCarItem(carItem *model.CarItem) error {
	sql := "insert into caritem(godsCount,sumPrice,book_id,car_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sql, carItem.GodsCount, carItem.Book.Price, carItem.Book.ID, carItem.CarID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCarItem 删除购物车中图书
func DeleteCarItem(bookID, carID string) error {
	sql := "delete from caritem where book_id=? and car_id=?"
	_, err := utils.Db.Exec(sql, bookID, carID)
	if err != nil {
		return err
	}
	return nil
}

//GetCarItemByBookID 通过bookid获取购物项
func GetCarItemByBookID(bookID, carID string) (*model.CarItem, error) {
	sql := "select id,godsCount,sumPrice from caritem where book_id=? and car_id=?"
	carItem := &model.CarItem{}
	utils.Db.QueryRow(sql, bookID, carID).Scan(&carItem.ID, &carItem.GodsCount, &carItem.SumPrice)
	return carItem, nil
}

// GetCarItemByCarID 通过carID获取所有的购物项
func GetCarItemByCarID(carID string) ([]*model.CarItem, error) {
	sql := "select id,godsCount,sumPrice,book_id,car_id from caritem where car_id=?"
	rows, err := utils.Db.Query(sql, carID)
	if err != nil {
		return nil, err
	}
	var carItems []*model.CarItem
	for rows.Next() {
		carItem := &model.CarItem{}
		var bookID string
		var carID string
		rows.Scan(&carItem.ID, &carItem.GodsCount, &carItem.SumPrice, &bookID, &carID)
		// 通过bookID获取图书信息
		book, _ := GetBookByID(bookID)
		carItem.Book = book
		carItem.CarID = carID
		carItems = append(carItems, carItem)
	}
	return carItems, nil
}

//UpdateCarItem 更新caritem
func UpdateCarItem(cItem *model.CarItem) error {
	sql := "update caritem set godsCount=?,sumPrice=? where book_id=?"
	_, err := utils.Db.Exec(sql, cItem.GodsCount, cItem.SumPrice, cItem.Book.ID)
	if err != nil {
		return err
	}
	return nil
}
