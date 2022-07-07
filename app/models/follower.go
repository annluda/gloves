package models

import (
	"fmt"
	"gloves/database"
	"strconv"
)

const (
	followersTableName = "followers"
)

// Follower 粉丝
type Follower struct {
	ID         uint `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	UserID     uint `gorm:"column:user_id;not null" sql:"index"`     // 多对多，关联 User Model (关注者)
	FollowerID uint `gorm:"column:follower_id;not null" sql:"index"` // 多对多，关联 User Model (粉丝)
}

// TableName 表名
func (Follower) TableName() string {
	return "followers"
}

// DoFollow 关注
func DoFollow(userID uint, followIDs ...uint) error {
	l := len(followIDs) - 1
	sqlStr := fmt.Sprintf("insert into %s (follower_id, user_id) values ", followersTableName)
	for i, v := range followIDs {
		sqlStr = sqlStr + fmt.Sprintf("(%d, %d)", userID, v)
		if i < l {
			sqlStr = sqlStr + ","
		}
	}
	d := database.DB.Exec(sqlStr)
	return d.Error
}

// DoUnFollow 取消关注
func DoUnFollow(userID uint, followIDs ...uint) error {
	sqlStr := fmt.Sprintf("delete from %s where follower_id = %d and user_id in (", followersTableName, userID)
	l := len(followIDs) - 1
	for i, v := range followIDs {
		sqlStr = sqlStr + strconv.Itoa(int(v))
		if i < l {
			sqlStr = sqlStr + ","
		}
	}
	sqlStr = sqlStr + ")"
	d := database.DB.Exec(sqlStr)
	return d.Error
}

// Followers 获取粉丝列表
func Followers(userID, offset, limit int) (followers []*User, err error) {
	followers = make([]*User, 0)
	joinSQL := fmt.Sprintf("inner join %s on users.id = followers.follower_id", followersTableName)
	if limit == 0 {
		d := database.DB.Model(&User{}).Joins(joinSQL).Where("followers.user_id = ?", userID).Order("id").Find(&followers)
		return followers, d.Error
	} else {
		d := database.DB.Model(&User{}).Joins(joinSQL).Where("followers.user_id = ?", userID).Offset(offset).Limit(limit).Order("id").Find(&followers)
		return followers, d.Error
	}
}

// Followings 获取用户关注人列表
func Followings(userID, offset, limit int) (followers []*User, err error) {
	followers = make([]*User, 0)
	joinSQL := fmt.Sprintf("inner join %s on users.id = followers.user_id", followersTableName)
	if limit == 0 {
		d := database.DB.Model(&User{}).Joins(joinSQL).Where("followers.follower_id = ?", userID).Order("id").Find(&followers)
		return followers, d.Error
	} else {
		d := database.DB.Model(&User{}).Joins(joinSQL).Where("followers.follower_id = ?", userID).Offset(offset).Limit(limit).Order("id").Find(&followers)
		return followers, d.Error
	}
}

// FollowingsIDList 获取用户关注人 ID 列表
func FollowingsIDList(userID int) (followerIDList []uint) {
	followers, _ := Followings(userID, 0, 0)
	followerIDList = make([]uint, 0)
	for _, v := range followers {
		followerIDList = append(followerIDList, v.ID)
	}
	return
}

// FollowingsCount 关注数
func FollowingsCount(userID int) (count int, err error) {
	joinSQL := fmt.Sprintf("inner join %s on users.id = followers.user_id", followersTableName)
	err = database.DB.Model(&User{}).Joins(joinSQL).Where("followers.follower_id = ?", userID).Count(&count).Error
	return
}

// FollowersCount 粉丝数
func FollowersCount(userID int) (count int, err error) {
	joinSQL := fmt.Sprintf("inner join %s on users.id = followers.follower_id", followersTableName)
	err = database.DB.Model(&User{}).Joins(joinSQL).Where("followers.user_id = ?", userID).Count(&count).Error
	return
}

// IsFollowing 已经关注了
func IsFollowing(currentUserID, userID int) bool {
	followerIDList := FollowingsIDList(currentUserID)
	id := uint(userID)
	for _, v := range followerIDList {
		if id == v {
			return true
		}
	}
	return false
}
