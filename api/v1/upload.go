package v1

import (
   "ginblog/utils/errmsg"
   "github.com/gin-gonic/gin"
   "net/http"
   "ginblog/servers"
)

func UpLoad(c *gin.Context) {
   file, fileHeader, _ := c.Request.FormFile("file")

   fileSize := fileHeader.Size

   url, code := servers.UpLoadFile(file, fileSize)

   c.JSON(http.StatusOK, gin.H{
      "status":  code,
      "message": errmsg.GetErrMsg(code),
      "url":     url,
   })

}