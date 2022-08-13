package request

type ArticleListReq struct {
	PageLimitReq
	Tag string `form:"tag"`
}

type SearchArticleListReq struct {
	ArticleListReq
	UserId    string `form:"userId"` // 用户id, 若无置位0
	SearchKey string `form:"searchKey"`
}

type MyArticleListReq struct {
	ArticleListReq
	UserId         string `form:"userId"`
	ArticleGroupId string `form:"articleGroupId"`
}
