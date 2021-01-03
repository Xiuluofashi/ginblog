package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `json:"role" validate:"required,gte=2" label:"角色"`
}

// 查询用户是否存在
func CheckUserByName(name string) (code int) {
	var user User
	db.Raw("SELECT Username FROM Users WHERE Username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 创建新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	} else {
		return errmsg.SUCCESS
	}
}

// 更新前检查用户数据
func CheckBeforeEdit(id int, name string) (code int) {
	var user User

	db.Raw("SElECT ID,Username FROM Users WHERE Username = ? ", name).First(&user)
	// 用户名已存在
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}

	db.Raw("SElECT ID,Username FROM Users WHERE ID= ? ", id).First(&user)
	// 用户ID不存在
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}

	return errmsg.SUCCESS
}

// 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	} else {
		db.Model(&user).Where("id=?", id).Updates(maps)
		return errmsg.SUCCESS
	}
}

// 查询用户讯息
func GetUserInfo(id int) (interface{}, int) {
	var user User
	err := db.Select("id, username, role, created_at").Where("ID = ?", id).First(&user).Error
	if user.ID == 0 {
		return nil, errmsg.ERROR_USER_NOT_EXIST
	}

	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User

	// 查询用户ID是否存在
	err := db.Select("id").Where("id = ? ", id).First(&user).Error
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if err != nil {
		return errmsg.ERROR
	}

	// 删除用户信息
	err = db.Where("id = ? ", id).Delete(&user).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 登陆验证
func CheckLogin(username string, password string) int {
    var user User
    
    db.Where("username = ?", username).First(&user)
    
    // 用户不存在
    if user.ID == 0 {
       return errmsg.ERROR_USER_NOT_EXIST
    }
    // 验证用户密码不正确
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
       return errmsg.ERROR_PASSWORD_WRONG
    }
    // 用户没有权限
    if user.Role != 1 {
       return errmsg.ERROR_USER_NO_RIGHT
    }
    return errmsg.SUCCESS
}
