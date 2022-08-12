package response

import (
	"EnjoyBlog/common/errorx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	StatusSuccess = 1000
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = int(err.(errorx.Status))
		body.Msg = err.Error()
	} else {
		body.Code = StatusSuccess
		body.Msg = "success"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
