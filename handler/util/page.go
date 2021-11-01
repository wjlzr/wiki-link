package util

//分页
//row:  行数
//page: 页数
func Pagination(size, index *int) {
	//条案为空或者行数大于10 或乾小于０
	if size == nil || *size <= 0 {
		*size = 20
	}

	//小于0或者没有页码
	if *index <= 0 || index == nil {
		*index = 0
	} else {
		*index = (*index - 1) * *size
	}
}
