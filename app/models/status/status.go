package status

import (
	"gloves/app/models"
)

const tableName = "statuses"

// Status 内容
type Status struct {
	models.BaseModel
	Content string `gorm:"column:context;type:text;not null"`
	UserID  uint   `gorm:"column:user_id;not null" sql:"index"` // 一对多，关联 User Model
}

// TableName 表名
func (Status) TableName() string {
	return tableName
}
