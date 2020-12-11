package model

// Page 分页
type Page struct {
	CurPage     int         // 当前页
	CountPage   int         // 总页数
	CountRecord int         // 总记录数
	Books       []*BookInfo // 当前页信息切片
	MinPrice    string
	MaxPrice    string
	Ses         *Session
	Car         *Car // 保存购物车
}

// IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.CurPage > 1
}

// IsHasNext 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.CurPage < p.CountPage
}

// GetPrevPage 获取下一页的页码
func (p *Page) GetPrevPage() int {
	if p.IsHasPrev() {
		return p.CurPage - 1
	}
	return 1
}

//GetNextPage 获取下一页页码
func (p *Page) GetNextPage() int {
	if p.IsHasNext() {
		return p.CurPage + 1
	}
	return p.CountPage
}

//IsHasLogin 判断是否已经登录
func (p *Page) IsHasLogin() bool {
	return p.Ses != nil
}
