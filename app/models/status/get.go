package status

import (
	"fmt"
	userModel "gloves/app/models/user"
	"gloves/database"
	"strconv"
)

// Get -
func Get(id int) (*Status, error) {
	s := &Status{}
	d := database.DB.First(&s, id)
	return s, d.Error
}

// GetByUsersStatusesCount 获取指定用户们的内容数量
func GetByUsersStatusesCount(ids []uint) (int, error) {
	sqlStr := fmt.Sprintf("select count(*) from %s where deleted_at is null and user_id in (", tableName)
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

	sqlStr := fmt.Sprintf("select * from %s where deleted_at is null and user_id in (", tableName)
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
func GetUser(statusID int) (*userModel.User, error) {
	s, err := Get(statusID)
	if err != nil {
		return nil, err
	}

	u, err := userModel.Get(int(s.UserID))
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
