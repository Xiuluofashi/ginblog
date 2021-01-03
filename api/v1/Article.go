package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	code = model.CheckArticle(data.Title)
	if code == errmsg.SUCCESS {
		model.CreateArticle(&data)
	}

	//操作完成,返回状态码
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 修改文章
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data) // 解析JSON数据方法

	// 1、检查传递的id是否有对应的文章信息，没有则不能修改
	code = model.CheckArticleByID(id)
	if code == errmsg.ERROR_ART_NOT_EXIST {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
	}

	model.EditArticle(id, &data)

	if code == errmsg.ERROR_CATENAME_USED {
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if code == errmsg.SUCCESS {
		code = model.DeleteArticle(id)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
}

// 获取栏目文章列表
func GetArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
    pageNum, _ := strconv.Atoi(c.Query("pagenum"))
    id, _ := strconv.Atoi(c.Param("id"))
    
    switch {
    case pageSize >= 100:
       pageSize = 100
    case pageSize <= 0:
       pageSize = 10
    }
    
    if pageNum == 0 {
       pageNum = 1
    }
    
    data, code, total := model.GetCateArtricle(id, pageSize, pageNum)
    
    c.JSON(http.StatusOK, gin.H{
       "status":  code,
       "data":    data,
       "total":   total,
       "message": errmsg.GetErrMsg(code),
    })
}

// 根据文章id获取某文章
func GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
   data, code := model.GetArticleInfo(id)
   c.JSON(http.StatusOK, gin.H{
      "status":  code,
      "data":    data,
      "message": errmsg.GetErrMsg(code),
   })
}
