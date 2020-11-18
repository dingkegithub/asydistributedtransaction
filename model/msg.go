package model

import (
	"fmt"
	"time"
)

// manager msginfo
type Msg struct {
	// 主键
	ID uint64 `gorm:"primary_key"`

	// db url, 数据源
	Url string `gorm:"size:255" json:"url"`

	// 已处理的次数
	HaveDealedTimes int `json:"head_dealed_times"`

	// 超时时间
	NextExpireTime uint64 `json:"next_expire_time"`

	// 创建时间
	CreatedAt time.Time
}

func (m Msg) String() string {
	return fmt.Sprintf("Msg [id=%d, url=%s, haveDealedTimes=%d, createTime=%v, nextExpireTime=%d]", m.ID, m.Url, m.HaveDealedTimes, m.CreatedAt, m.NextExpireTime)
}
