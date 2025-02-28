package initial

import (
	"github.com/RichFastTrade/rft_shared/logger"
	"github.com/go-kratos/kratos/v2/log"
)

func initLog() {
	// 初始化日志
	log.SetLogger(logger.Get())
}
