package contract

import "time"

const DistributedKey = "wm:distributed"

type Distributed interface {
	// Select 分布式选择器，所有节点对某个服务进行抢占，只能选择其中一个节点
	// serviceName 服务名称
	// appID 当前的AppID
	// holdTime 分布式选择器hold时间
	// return
	// selectAppID 分布式选择器最终选择的AppID
	// err 异常返回错误，如果没有被选择，不返回err
	Select(serviceName string, appID string, holdTime time.Duration) (selectAppID string, err error)
}
