package db

import "wiki-link/model/error"

//获取错误码
func GetErrors() (errs []error.Error) {
	MainDb.Find(&errs)
	return
}
