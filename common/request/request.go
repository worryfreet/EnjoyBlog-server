package request

type PageLimitReq struct {
	Page     int    `json:"page"`               // 页面
	PageSize int    `json:"pageSize"`           // 页面大小
	Desc     int    `json:"desc,omitempty"`     // 1降序, 0默认升序
	OrderKey string `json:"orderKey,omitempty"` // 排序关键字
}
