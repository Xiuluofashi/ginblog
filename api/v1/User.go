package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"ginblog/utils/validator"
)

var code int

// 添加用户
func AddUser(c *gin.Context) {
	//定义User的结构体
	var data model.User

	var msg string
	_ = c.ShouldBindJSON(&data)
	// 1.验证页面传递过来的参数
	msg, code = validator.Validate(&data)
   if code != errmsg.SUCCESS {
      c.JSON(
         http.StatusOK, gin.H{
            "status":  code,
            "message": msg,
         })
      return
   }

	// 2.验证用户名是否存在
	code = model.CheckUserByName(data.Username)
	if code == errmsg.SUCCESS {
		pwd, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)
		data.Password = string(pwd)
		code=model.CreateUser(&data)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
}

// 修改用户信息
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	// 1.查询用户名是否被使用，ID是否存在
	code = model.CheckBeforeEdit(id, data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})

}

// 查询单个用户
func GetUserInfo(c *gin.Context) {
	// 获取传递的分类id
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetUserInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteUser(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
}
