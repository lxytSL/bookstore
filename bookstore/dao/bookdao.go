package dao

import (
	"strconv"
	"web/bookstore/model"
	"web/bookstore/utils"
)

// GetBooks 查询图书
func GetBooks() ([]*model.BookInfo, error) {
	sql := "select * from books"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var books []*model.BookInfo

	for rows.Next() {
		book := &model.BookInfo{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}

// AddBook 添加图书
func AddBook(b *model.BookInfo) error {
	sql := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBook 删除BOOK
func DeleteBook(bookID string) error {
	sql := "delete from books where id=?"
	_, err := utils.Db.Exec(sql, bookID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateBook 更新数据
func UpdateBook(book *model.BookInfo) error {
	sql := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	_, err := utils.Db.Exec(sql, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetBookByID 通过BookID得到要修改的信息
func GetBookByID(bookID string) (*model.BookInfo, error) {
	book := &model.BookInfo{}
	sql := "select id,title,author,price,sales,stock from books where id=?"
	err := utils.Db.QueryRow(sql, bookID).Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock)
	if err != nil {
		return nil, err
	}
	return book, nil
}

// PageGetBook 分页查询
func PageGetBook(curPageStr string) (*model.Page, error) {
	var countRecord int
	// 查询得到总页数
	sql := "select count(*) from books"
	utils.Db.QueryRow(sql).Scan(&countRecord)
	var page *model.Page
	sql = "select * from books limit ?,?"
	page, _ = PageGetHelper(countRecord, curPageStr, sql)
	return page, nil
}

// MinPrice 默认为0
// MaxPrice 默认为999
var (
	MinPrice string = "0"
	MaxPrice string = "999"
	Flag     bool
)

// GetBookByParse 查询价格区间[minPrice, maxPrice]图书
func GetBookByParse(minPrice, maxPrice, curPageStr string) (*model.Page, error) {
	var countRecord int
	if minPrice != "" {
		MinPrice = minPrice
	}
	if maxPrice != "" {
		MaxPrice = maxPrice
	}
	Flag = true
	sql := "select count(*) from books where price > ? and price < ?"
	utils.Db.QueryRow(sql, minPrice, maxPrice).Scan(&countRecord)
	var page *model.Page
	sql = "select * from books where price > ? and price < ? limit ?,?"
	page, _ = PageGetHelper(countRecord, curPageStr, sql)
	return page, nil
}

// PageGetHelper 查询分页复用函数
func PageGetHelper(countRecord int, curPageStr, sql string) (*model.Page, error) {
	// countRecord 总记录数
	// 设置每页4条数据
	var pageSize int = 4
	// 计算多少页
	var countPage int
	if countRecord%pageSize == 0 {
		countPage = countRecord / pageSize
	} else {
		countPage = countRecord/pageSize + 1
	}
	// 查询当前页的数据
	curPage, _ := strconv.Atoi(curPageStr)
	// Flag true 价格区间查询 false普通查询

	rows := utils.Rows
	var err error
	if Flag {
		rows, err = utils.Db.Query(sql, MinPrice, MaxPrice, (curPage-1)*pageSize, pageSize)
	} else {
		rows, err = utils.Db.Query(sql, (curPage-1)*pageSize, pageSize)
	}
	if err != nil {
		return nil, err
	}
	var books []*model.BookInfo
	for rows.Next() {
		book := &model.BookInfo{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	page := &model.Page{
		CurPage:     curPage,
		CountPage:   countPage,
		CountRecord: countRecord,
		Books:       books,
	}
	return page, nil
}
