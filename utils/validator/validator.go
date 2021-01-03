package validator

import (
   "fmt"
   "ginblog/utils/errmsg"
   "github.com/go-playground/locales/zh_Hans_CN"
   unTrans "github.com/go-playground/universal-translator"
   "github.com/go-playground/validator/v10"
   zhTrans "github.com/go-playground/validator/v10/translations/zh"
   "reflect"
)

func Validate(data interface{}) (string, int) {
    // 实例化validate
    validate := validator.New()
    // 实例化返回信息语言 转化为中文
    uni := unTrans.New(zh_Hans_CN.New())
    trans, _ := uni.GetTranslator("zh_Hans_CN")
    
    // 注册翻译方法
    err := zhTrans.RegisterDefaultTranslations(validate, trans)
    if err != nil {
      fmt.Println("err:", err)
    } 
    
    // 使用label标签的内容代替数据库字段名
    validate.RegisterTagNameFunc(func(field reflect.StructField) string {
       label := field.Tag.Get("label")
       return label
    })
    
    // 使用validate验证data
    err = validate.Struct(data)
    if err != nil {
        // 如果有错误即验证不通过 循环返回验证错误信息
        for _, v := range err.(validator.ValidationErrors) {
         return v.Translate(trans), errmsg.ERROR
        }
    }
    return "", errmsg.SUCCESS
}