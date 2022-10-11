package cronJob

import (
	"EnjoyBlog/app/comment/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"time"
)

func CronTimer() {
	// 同步评论点赞的数量
	threading.GoSafe(SyncMultiCmtSupport)
}

func SyncMultiCmtSupport() {
	if idx, err := model.GlobalComment.SyncMultiCmtSupport(context.Background()); err != nil {
		logx.Errorf("同步评论点赞第%d条失败, err%s: ", idx+1, err)
	} else {
		logx.Infof("同步%d条评论点赞成功", idx)
	}
	time.AfterFunc(time.Minute*5, SyncMultiCmtSupport)
}
