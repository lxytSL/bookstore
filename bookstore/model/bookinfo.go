package model

//BookInfo 图书信息
type BookInfo struct {
	ID      int
	Title   string
	Author  string
	Price   float32
	Sales   int
	Stock   int
	ImgPath string
}
