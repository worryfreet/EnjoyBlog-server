package request

type PageLimitReq struct {
	Page     int    `form:"page"`               // 页面
	PageSize int    `form:"pageSize"`           // 页面大小
	Desc     int    `form:"desc,omitempty"`     // 1降序, 0默认升序
	OrderKey string `form:"orderKey,omitempty"` // 排序关键字
}
