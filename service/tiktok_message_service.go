package service

import (
	"gorm.io/gorm"
	"log"
	"tiktok/common"
	"tiktok/models"
)

var TiktokMessageService = tiktokMessageService{}

// tiktokMessageService 业务层
type tiktokMessageService struct {
}

func (t tiktokMessageService) db() *gorm.DB {
	return common.Db
}

// List 分页列表
func (t tiktokMessageService) List(page, size int, v *models.TiktokMessage) map[string]interface{} {
	// 结果
	var lists []models.TiktokMessage
	t.db().Model(&v).Where(&v).Order("").Offset((page - 1) * size).Limit(size).Find(&lists)
	// 统计
	var total int64
	t.db().Model(&v).Where(&v).Count(&total)
	data := make(map[string]interface{})
	data["list"] = lists
	data["total"] = total
	return data
}

// One 根据主键Id查询记录
func (t tiktokMessageService) One(id interface{}) models.TiktokMessage {
	var v models.TiktokMessage
	db := t.db().Find(&v, id)
	if db.RowsAffected != 1 {
		log.Println("未找到数据！")
	}
	return v
}

// Update 修改记录 true -> 操作成功
func (t tiktokMessageService) Update(v models.TiktokMessage) bool {
	tx := t.db().Model(&v).Updates(v)
	if tx.Error != nil {
		log.Panicln(tx.Error.Error())
		return false
	}
	return true
}

// Insert 插入记录 true -> 操作成功 注: 主键也传递进来的话，那么就会执行更新或插入操作
func (t tiktokMessageService) Insert(v models.TiktokMessage) bool {
	tx := t.db().Save(&v)
	if tx.Error != nil {
		log.Panicln(tx.Error.Error())
		return false
	}
	return true
}

// Delete 根据主键删除 true -> 操作成功
func (t tiktokMessageService) Delete(ids []int) bool {
	tx := t.db().Delete(models.TiktokMessage{}, ids)
	if tx.Error != nil {
		log.Panicln(tx.Error.Error())
		return false
	}
	return true
}
