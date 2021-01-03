package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查询分类名是否存在
func CheckCateByName(name string) (code int) {
	var cate Category
	db.Select("id").Where("name=?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// 新增分类
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	} else {
		return errmsg.SUCCESS
	}
}

// 获取分类列表
func GetCate(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	err = db.Find(&cate).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR_CATE_NOT_EXIST
	}
	return cate, total
}

// 根据分类id获取某分类
func GetCateInfo(id int) (interface{}, int) {
	var cate Category
	db.Where("id = ?", id).First(&cate)
	if cate.ID == 0 {
		return nil, errmsg.ERROR_CATE_NOT_EXIST
	}

	return cate, errmsg.SUCCESS
}

// 更新前检查
func CheckCateByID(id int) (code int) {
	var cate Category
	db.Select("id").Where("id= ?", id).First(&cate)
	if cate.ID <= 0 {
		return errmsg.ERROR_CATE_NOT_EXIST
	}
	return errmsg.SUCCESS
}

//  更新分类
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ? ", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
