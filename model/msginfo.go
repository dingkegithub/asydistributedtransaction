
import (
	"fmt"
	"time"
)

type MsgInfo struct {
	// 主键
	ID uint64 `gorm:"primary_key"`

	// 事务消息
	Content string `json:"content"`

	// MQ 主题，消息放到哪个主题
	Topic string `json:"status"`

	// 二级topic (rocketmq)
	Tag string `json:"status"`

	// 延迟时间
	Delay uint64 `json:"delay"` // 延迟时间，单位 s

	// 状态：1-等待  2-发送
	Status int `json:"status"`

	// 创建时间
	CreatedAt time.Time
}

func (m Msg) String() string {
	return fmt.Sprintf("MsgInfo [id=%d, content=%s, topic=%s, tag=%s, status=%s, createTime=%v, delay=%d]", m.ID, m.Content, m.Topic, m.Tag, m.Status, m.CreatedAt, m.Delay)
}
