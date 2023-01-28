package service

import (
	"gorm.io/gorm"
	"log"
	"tiktok/common"
	"tiktok/models"
)

var TiktokUserService = tiktokUserService{}

// tiktokUserService 业务层
type tiktokUserService struct {
}

func (t tiktokUserService) db() *gorm.DB {
	return common.Db
}

// List 分页列表
func (t tiktokUserService) List(page, size int, v *models.TiktokUser) map[string]interface{} {
	// 结果
	var lists []models.TiktokUser
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
func (t tiktokUserService) One(id interface{}) models.TiktokUser {
	var v models.TiktokUser
	db := t.db().Find(&v, id)
	if db.RowsAffected != 1 {
		log.Println("未找到数据！")
	}
	return v
}

// Update 修改记录 true -> 操作成功
func (t tiktokUserService) Update(v models.TiktokUser) bool {
	tx := t.db().Model(&v).Updates(v)
	if tx.Error != nil {
		log.Panicln(tx.Error.Error())
		return false
	}
	return true
}

// Insert 插入记录 true -> 操作成功 注: 主键也传递进来的话，那么就会执行更新或插入操作
func (t tiktokUserService) Insert(v models.TiktokUser) bool {
	tx := t.db().Save(&v)
	if tx.Error != nil {
		log.Panicln(tx.Error.Error())
		return false
	}
	return true
}

// Delete 根据主键删除 true -> 操作成功
func (t tiktokUserService) Delete(ids []int) bool {
	tx := t.db().Delete(models.TiktokUser{}, ids)
	if tx.Error != nil {
		log.Panicln(tx.Error.Error())
		return false
	}
	return true
}
