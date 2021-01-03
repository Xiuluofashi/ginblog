package middleware

import (
   "fmt"
   "github.com/gin-gonic/gin"
   retalog "github.com/lestrrat-go/file-rotatelogs"
   "github.com/rifflock/lfshook"
   "github.com/sirupsen/logrus"
   "math"
   "os"
   "time"
)

func Log() gin.HandlerFunc {
   // 日志存放位置
   filePath := "log/log"
   //linkName := "latest_log.log"
   
   scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
   if err != nil {
      fmt.Println("err:", err)
   }
   logger := logrus.New()

   // 日志输出格式 输出到文件中
   logger.Out = scr

   // 设置日志级别
   logger.SetLevel(logrus.DebugLevel)

   // 设置日志分割 
   logWriter, _ := retalog.New(
      filePath+"%Y%m%d.log",    // 按照年月日的格式生成文件名
      retalog.WithMaxAge(7*24*time.Hour),    // 日志保存时间
      retalog.WithRotationTime(24*time.Hour),    // 分割时间 24h
      //retalog.WithLinkName(linkName),    // 设置最新文件软链接 方便查看最新日志
   )

   // 设置日志分段格式
   writeMap := lfshook.WriterMap{
      logrus.InfoLevel:  logWriter,
      logrus.FatalLevel: logWriter,
      logrus.DebugLevel: logWriter,
      logrus.WarnLevel:  logWriter,
      logrus.ErrorLevel: logWriter,
      logrus.PanicLevel: logWriter,
   }
   Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
      TimestampFormat: "2006-01-02 15:04:05",
   })
   logger.AddHook(Hook)

   return func(c *gin.Context) {
      startTime := time.Now()
      c.Next()
      stopTime := time.Since(startTime)
      spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
      hostName, err := os.Hostname()
      if err != nil {
         hostName = "unknown"
      }
      statusCode := c.Writer.Status()
      clientIp := c.ClientIP()
      userAgent := c.Request.UserAgent()
      dataSize := c.Writer.Size()
      if dataSize < 0 {
         dataSize = 0
      }
      method := c.Request.Method
      path := c.Request.RequestURI

      entry := logger.WithFields(logrus.Fields{
         "HostName":  hostName,
         "status":    statusCode,
         "SpendTime": spendTime,
         "Ip":        clientIp,
         "Method":    method,
         "Path":      path,
         "DataSize":  dataSize,
         "Agent":     userAgent,
      })
      if len(c.Errors) > 0 {
         entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
      }
      if statusCode >= 500 {
         entry.Error()
      } else if statusCode >= 400 {
         entry.Warn()
      } else {
         entry.Info()
      }
   }
}