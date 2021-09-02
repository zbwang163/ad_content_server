package client

import (
	logs "github.com/zbwang163/ad_common/logger"
	"github.com/zbwang163/ad_common/rpc"
	"github.com/zbwang163/ad_content_server/common/consts"
)

func Init() {
	rpc.InitRpc(consts.AccountServerPSM)
	logs.InitLogger()
}
