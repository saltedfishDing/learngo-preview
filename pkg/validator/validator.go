package validator

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Init() {
	//注册翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)

	// 注册一个获取json tag的方法
	validate.RegisterTagNameFunc(
		func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

			if name == "_" {
				return ""
			}

			return name
		})
	//注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

//Translate 翻译错误信息
func Translate(err error) map[string]string {
	if errors, ok := err.(validator.ValidationErrors); ok {
		return removeTopStruct(errors.Translate(trans))
	}
	return map[string]string{}
}

// 去除结构体名称前缀
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
