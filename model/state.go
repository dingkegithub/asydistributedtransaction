package model

const (
	CREATE  = 1 << iota // 服务创建状态
	RUNNING             // 服务运行
	CLOSED              // 服务关闭
	FAILED              // 服务失败
)
