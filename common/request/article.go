package request

type (
	// 公开的文章列表
	ArticleListReq struct {
		Page     int    `form:"page"`     // 页面
		PageSize int    `form:"pageSize"` // 页面大小
		Desc     int    `form:"desc"`     // 1降序, 0默认升序
		OrderKey string `form:"orderKey"` // 排序关键字
		UserId   string `form:"userId"`
	}

	// 查询列表
	SearchArticleListReq struct {
		ArticleListReq
		SearchKey string `form:"searchKey"`
	}

	// 根据目录分类查询文章列表
	ArticleListByGroupIdReq struct {
		ArticleListReq
		ArticleGroupId string `form:"articleGroupId"`
	}

	// 根据标签查询文章列表
	ArticleListByLabelReq struct {
		ArticleListReq
		Label string `form:"label"` // 标签
	}

	// 新增文章, 更新文章入参
	EditArticleInfoReq struct {
		ArticleTitle   string `json:"articleTitle"`
		ArticleCttHtml string `json:"articleCttHtml"`
		ArticleCttMd   string `json:"articleCttMd"`
		ArticleGroupId string `json:"articleGroupId"`
		Avatar         string `json:"avatar"`
		Label          string `json:"label"`
		IsTop          int    `json:"isTop"`
		IsPub          int    `json:"isPub"`
	}
)
