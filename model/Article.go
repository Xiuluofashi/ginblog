package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	// Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title" validate:"required,min=4,max=20" label:"标题"`
	Cid     int    `gorm:"type:int;not null" json:"cid" validate:"required,gte=1" label:"栏目ID"`
	Desc    int    `gorm:"type:varchar(200)" json:"desc" validate:"required,min=4,max=200" label:"描述"`
	Content string `gorm:"type:longtext" json:"content" validate:"required,min=20" label:"内容"`
	Img     string `gorm:"type:varchar(100)" json:"img" validate:"required,min=4,max=100" label:"图片地址"`
}

// 文章标题不能为空
func CheckArticle(Title string) (code int) {
	if len(Title) == 0 {
		return errmsg.ERROR
	} else {
		return errmsg.SUCCESS
	}
}

// 增加文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 根据文章ID查询文章是否存在
func CheckArticleByID(id int) (code int) {
	var article Article
	
	db.Select("id").Where("id= ?", id).First(&article)
	if article.ID <= 0 {
	   return errmsg.ERROR_ART_NOT_EXIST 
	}
	return errmsg.SUCCESS    
 }

 // 修改文章
func EditArticle(id int, data *Article) int {
    var article Article
    var maps = make(map[string]interface{})
    maps["title"] = data.Title
    maps["cid"] = data.Cid
    maps["desc"] = data.Desc
    maps["content"] = data.Content
    maps["img"] = data.Img

   err = db.Model(&article).Where("id = ? ", id).Updates(maps).Error
   if err != nil {
      return errmsg.ERROR
   }
   return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id = ? ", id).Delete(&article).Error
	if err != nil {
	   return errmsg.ERROR
	}
	return errmsg.SUCCESS
 }

 //  查询分类下的所有文章
func GetCateArtricle(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64
	// 如果栏目id=0，查询所有栏目下的文章
	// 否则查询对应栏目id的文章
	 if id == 0 {
		err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&cateArtList).Count(&total).Error
	 }else {
		err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		   "cid =?", id).Find(&cateArtList).Count(&total).Error
	 }
	if err != nil {
	   return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
 }

//  根据文章id获取某文章
 func GetArticleInfo (id int) (interface{}, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
	   return nil, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
 }