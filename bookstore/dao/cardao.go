package dao

import (
	"web/bookstore/model"
	"web/bookstore/utils"
)

//AddCar 添加购物车信息
func AddCar(car *model.Car) error {
	sql := "insert into cars values(?,?,?,?)"
	_, err := utils.Db.Exec(sql, car.CarID, car.CarGodsCount, car.CarSumPrice, car.UserID)
	if err != nil {
		return err
	}
	// 将每本书插进数据库
	for _, carItem := range car.CarItems {
		AddCarItem(carItem)
	}
	return nil
}

// GetCarByUserID 查看当前用户是否有购物车
func GetCarByUserID(UserID int) (*model.Car, error) {
	sql := "select cars_id,carGodsCount,carSumPrice,user_id from cars where user_id = ?"
	car := &model.Car{}
	utils.Db.QueryRow(sql, UserID).Scan(&car.CarID, &car.CarGodsCount, &car.CarSumPrice, &car.UserID)
	return car, nil
}

//GetCarByCarID 查看购物车信息
func GetCarByCarID(carID string) (*model.Car, error) {
	sql := "select cars_id,carGodsCount,carSumPrice,user_id from cars where cars_id = ?"
	car := &model.Car{}
	utils.Db.QueryRow(sql, carID).Scan(&car.CarID, &car.CarGodsCount, &car.CarSumPrice, &car.UserID)
	return car, nil
}

// UpdateCar 更新购物车
func UpdateCar(c *model.Car) error {
	sql := "update cars set carGodsCount=?,carSumPrice=? where user_id=?"
	_, err := utils.Db.Exec(sql, c.CarGodsCount, c.CarSumPrice, c.UserID)
	if err != nil {
		return err
	}
	return nil
}
