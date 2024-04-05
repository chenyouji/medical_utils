package utils

import (
	"github.com/alibaba/sentinel-golang/core/base"
	"log"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
)

func InitQps() {
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("初始化sentinel 异常: %v", err)
	}
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some-test", // 资源名称
			TokenCalculateStrategy: flow.Direct, // 令牌计算策略
			ControlBehavior:        flow.Reject, // 控制行为：直接拒绝
			Threshold:              10000,       // 每秒钟响应10000次
			StatIntervalInMs:       1000,        // 统计间隔时间：每秒钟统计一次
		},
	})
	if err != nil {
		log.Fatalf("加载规则失败: %v", err)
	}
}

func Entry() *base.BlockError {
	_, b := sentinel.Entry("some-test", sentinel.WithTrafficType(base.Inbound))
	return b
}
