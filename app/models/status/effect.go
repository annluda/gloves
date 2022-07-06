package status

import (
	"gloves/database"

	"github.com/lexkong/log"
)

// Create -
func (s *Status) Create() (err error) {
	if err = database.DB.Create(&s).Error; err != nil {
		log.Warnf("创建失败: %v", err)
		return err
	}

	return nil
}

// Delete -
func Delete(id int) (err error) {
	status := &Status{}
	status.BaseModel.ID = uint(id)

	if err = database.DB.Delete(&status).Error; err != nil {
		log.Warnf("删除失败: %v", err)
		return err
	}

	return nil
}
