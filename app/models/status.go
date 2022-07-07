package models

import (
	"fmt"
	"gloves/database"
	"gloves/pkg/logger"
	"strconv"
)

// Status 内容
type Status struct {
	BaseModel
	Content string `gorm:"column:context;type:text;not null"`
	UserID  uint   `gorm:"column:user_id;not null" sql:"index"` // 一对多，关联 User Model
}

// TableName 表名
func (Status) TableName() string {
	return "statuses"
}

// Create -
func (s *Status) Create() (err error) {
	if err = database.DB.Create(&s).Error; err != nil {
		logger.Warnf("创建失败: %v", err)
		return err
	}

	return nil
}

// StatusDelete Delete -
func StatusDelete(id int) (err error) {
	status := &Status{}
	status.BaseModel.ID = uint(id)

	if err = database.DB.Delete(&status).Error; err != nil {
		logger.Warnf("删除失败: %v", err)
		return err
	}

	return nil
}

// StatusGet Get -
func StatusGet(id int) (*Status, error) {
	s := &Status{}
	d := database.DB.First(&s, id)
	return s, d.Error
}

// GetByUsersStatusesCount 获取指定用户们的内容数量
func GetByUsersStatusesCount(ids []uint) (int, error) {
	sqlStr := "select count(*) from statuses where deleted_at is null and user_id in ("
	l := len(ids) - 1
	for i, v := range ids {
		sqlStr = sqlStr + strconv.Itoa(int(v))
		if i < l {
			sqlStr = sqlStr + ","
		}
	}
	sqlStr = sqlStr + ")"

	count := 0
	d := database.DB.Raw(sqlStr).Count(&count)
	return count, d.Error
}

// GetByUsersStatuses 获取指定用户们的内容
func GetByUsersStatuses(ids []uint, offset, limit int) ([]*Status, error) {
	status := make([]*Status, 0)

	sqlStr := "select * from statuses where deleted_at is null and user_id in ("
	l := len(ids) - 1
	for i, v := range ids {
		sqlStr = sqlStr + strconv.Itoa(int(v))
		if i < l {
			sqlStr = sqlStr + ","
		}
	}
	sqlStr = sqlStr + fmt.Sprintf(") order by `created_at` desc limit %d offset %d", limit, offset)

	d := database.DB.Raw(sqlStr).Scan(&status)
	return status, d.Error
}

// GetUser 通过 status_id 获取该内容的所有者
func GetUser(statusID int) (*User, error) {
	s, err := StatusGet(statusID)
	if err != nil {
		return nil, err
	}

	u, err := UserGet(int(s.UserID))
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetUserAllStatus 获取该用户的所有内容
func GetUserAllStatus(userID int) ([]*Status, error) {
	status := make([]*Status, 0)

	err := database.DB.Where("user_id = ?", userID).Order("id desc").Find(&status).Error

	if err != nil {
		return status, err
	}

	return status, nil
}

// GetUserStatus 获取该用户的内容 (分页)
func GetUserStatus(userID, offset, limit int) ([]*Status, error) {
	status := make([]*Status, 0)

	err := database.DB.Where("user_id = ?", userID).Offset(
		offset).Limit(limit).Order("id desc").Find(&status).Error

	if err != nil {
		return status, err
	}

	return status, nil
}

// GetUserAllStatusCount 获取该用户的所有内容 的 count
func GetUserAllStatusCount(userID int) (count int, err error) {
	err = database.DB.Model(&Status{}).Where("user_id = ?", userID).Count(&count).Error
	return
}
