package model

// CarItem 购物车单本图书信息
type CarItem struct {
	ID        int
	GodsCount int // 图书数量
	Book      *BookInfo
	SumPrice  float32 // 图书总价
	CarID     string  // 属于谁个购物车
}

// GetSumPrice 获取单价
func (c *CarItem) GetSumPrice(GodsCount int) float32 {
	return float32(GodsCount) * c.Book.Price
}
