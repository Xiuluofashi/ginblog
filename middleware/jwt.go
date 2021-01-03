package middleware

import (
   "ginblog/utils"
   "ginblog/utils/errmsg"
   "github.com/dgrijalva/jwt-go"
   "github.com/gin-gonic/gin"
   "net/http"
   "strings"
   "time"
)

// 定义一个秘钥参数
var JeyKey = []byte(utils.JwtKey)

// 定义一个请求格式结构体
type MyClaims struct{
    // 因为是使用用户名 所以这里需要定义
    Username string `json:"username"`
    jwt.StandardClaims    // jwt结构体
}

// 生成token
func SetToken(username string)(string, int){
    expireTime := time.Now().Add(10 * time.Hour)    // 设置有效时间
    SetClaims := MyClaims{
        Username:username,
        StandardClaims:jwt.StandardClaims{
            ExpiresAt: expireTime.Unix(),    // 有效时间 时间戳
            Issuer:"ginblog",    // 签发人
        },
    }
    
    // NewWithClaims方法需要传递两个参数
    // 一个是签发的方法
    // 一个是定义的SetClaims 
    // 该方法返回一个 结构体Token的指针
    reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
    
    // 调用reqClaim.SignedString 返回生成的token字符串和err
    token, err := reqClaim.SignedString(JeyKey)
    if err != nil{
        return "", errmsg.ERROR
    }
    return token, errmsg.SUCCESS
}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
   var claims MyClaims

   setToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
      return JeyKey, nil
   })

   if err != nil {
      if ve, ok := err.(*jwt.ValidationError); ok {
         if ve.Errors&jwt.ValidationErrorMalformed != 0 {
            return nil, errmsg.ERROR_TOKEN_WRONG
         } else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
            return nil, errmsg.ERROR_TOKEN_RUNTIME
         } else {
            return nil, errmsg.ERROR_TOKEN_TYPE_WRONG
         }
      }
   }
   if setToken != nil {
      if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
         return key, errmsg.SUCCESS
      } else {
         return nil, errmsg.ERROR_TOKEN_WRONG
      }
   }
   return nil, errmsg.ERROR_TOKEN_WRONG
}

// jwt中间件验证
func JwtToken() gin.HandlerFunc {
   return func(c *gin.Context) {
      var code int
      // 获取请求头的Authorization Token信息
      tokenHeader := c.Request.Header.Get("Authorization")
      // token为空 返回错误信息
      if tokenHeader == "" {
         code = errmsg.ERROR_TOKEN_EXIST
         c.JSON(http.StatusOK, gin.H{
            "code":    code,
            "message": errmsg.GetErrMsg(code),
         })
         c.Abort()
         return
      }
      // token格式错误 返回错误信息
      checkToken := strings.Split(tokenHeader, " ")
      if len(checkToken) == 0 {
         code = errmsg.ERROR_TOKEN_TYPE_WRONG
         c.JSON(http.StatusOK, gin.H{
            "code":    code,
            "message": errmsg.GetErrMsg(code),
         })
         c.Abort()
         return
      }
      // token验证失败 返回错误信息
      if len(checkToken) != 2 && checkToken[0] != "Bearer" {
         code = errmsg.ERROR_TOKEN_TYPE_WRONG
         c.JSON(http.StatusOK, gin.H{
            "code":    code,
            "message": errmsg.GetErrMsg(code),
         })
         c.Abort()
         return
      }
      // token验证成功
      key, tCode := CheckToken(checkToken[1])
      if tCode != errmsg.SUCCESS {
         code = tCode
         c.JSON(http.StatusOK, gin.H{
            "code":    code,
            "message": errmsg.GetErrMsg(code),
         })
         c.Abort()
         return
      }
      c.Set("username", key)
      c.Next()
   }
}