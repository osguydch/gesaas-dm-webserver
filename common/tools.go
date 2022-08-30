package common

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"time"
)

var CstSh, _ = time.LoadLocation("Asia/Shanghai")

const (
	LocateTimeFormat  = "2006-01-02 15:04:05"
	LocateMilliFormat = "2006-01-02 15:04:05.9999"
)

func Now() string {
	return time.Now().In(CstSh).Format(LocateTimeFormat)
}

func NowMilli() string {
	return time.Now().In(CstSh).Format(LocateMilliFormat)
}

func OrderBy(c *gin.Context) int32 {
	var o int32 = 1
	if order, isExist := c.GetQuery("orderBy"); isExist == true {
		o = cast.ToInt32(order)
		if o < 0 || o > 3 {
			o = 1
		}
	}
	return o
}

func Pagination(c *gin.Context) (int32, int32) {
	var p, s int32 = 1, 10
	if page, isExist := c.GetQuery("pageNum"); isExist == true {
		p = cast.ToInt32(page)
		if p < 1 {
			p = 1
		}
	}
	if pageSize, isExist := c.GetQuery("pageSize"); isExist == true {
		s = cast.ToInt32(pageSize)
		if s < 10 {
			s = 10
		}
	}
	return p, s
}

type ResList struct {
	List  interface{} `json:"list"`
	Total int32       `json:"total"`
}

func NewResList(list interface{}, total int32) *ResList {
	if total == 0 {
		return &ResList{[]string{}, 0}
	}
	return &ResList{list, total}
}
