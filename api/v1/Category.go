package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)

	// 验证分类名是否已存在
	code := model.CheckCateByName(data.Name)
	if code == errmsg.SUCCESS {
		code = model.CreateCate(&data)
	}

	// 返回状态码和成功信息
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 获取分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.GetCate(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 根据分类id获取某分类
func GetCateInfo(c *gin.Context) {
	// 获取传递的分类id
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetCateInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
}

// 修改分类信息
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.CheckCateByID(id)
	if code == errmsg.ERROR_CATE_NOT_EXIST {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
		return
	}

	code = model.CheckCateByName(data.Name)
	if code == errmsg.SUCCESS {

		model.EditCate(id, &data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
}

// 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// 1、检查传递的id是否有对应的分类信息
	 code = model.CheckCateByID(id)
	 if code == errmsg.SUCCESS{
		code = model.DeleteCate(id)
	 }
 
	c.JSON(
	   http.StatusOK, gin.H{
		  "status":  code,
		  "message": errmsg.GetErrMsg(code),
	   })
}
