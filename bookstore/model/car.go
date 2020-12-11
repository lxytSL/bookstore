package model

// Car 购物车
type Car struct {
	CarID        string
	CarItems     []*CarItem
	CarGodsCount int     // 购物车商品总数总数量
	CarSumPrice  float32 // 购物车总金额
	UserID       int     // 购物车所属用户
}

// GetCarGodsCount 获取购物车商品总数量
func (c *Car) GetCarGodsCount() (int, float32) {
	var (
		price float32
		count int
	)
	for _, carItem := range c.CarItems {
		count += carItem.GodsCount
		price += carItem.SumPrice
	}
	return count, price
}
