package po

import "time"

type Testdata struct {
}

type HelloWorld struct {
	Id             uint64    `gorm:"column:id;primaryKey" json:"id"`                // ID
	CreatedTime    time.Time `gorm:"column:created_time" json:"created_time"`       // 创建时间
	UpdatedTime    time.Time `gorm:"column:updated_time" json:"updated_time"`       // 更新时间
	RequestMessage string    `gorm:"column:request_message" json:"request_message"` //
}
