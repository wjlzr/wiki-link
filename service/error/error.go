package error

import (
	"wiki-link/common/lang"
	"wiki-link/db"
	"wiki-link/model/error"
)

var (
	errsMap = make(map[int64]map[string]string)
)

//初始化错误提示
func InitErrors() {
	errors := getErrors()

	for _, err := range errors {
		langs := make(map[string]string)
		langs[lang.ZH_CN] = err.ZhCN
		langs[lang.ZH_HK] = err.ZhHK
		langs[lang.ZH_TW] = err.ZhTW
		langs[lang.EN] = err.En
		langs[lang.VI] = err.Vi
		langs[lang.TH] = err.Th
		langs[lang.FR] = err.Fr
		langs[lang.ID] = err.Id
		langs[lang.ES] = err.Es
		langs[lang.RU] = err.Ru
		langs[lang.DE] = err.De
		langs[lang.FIL] = err.Fil
		langs[lang.IT] = err.It
		langs[lang.HI] = err.Hi
		langs[lang.JA] = err.Ja
		errsMap[err.Code] = langs
	}
}

//根据编码、语言获取错误信息
func Errors(code int64, lang string) string {
	return errsMap[code][lang]
}

//获取错误信息
func getErrors() []error.Error {
	return db.GetErrors()
}
