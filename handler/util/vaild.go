package util

import "regexp"

//松测是否为数字
func IsNumeric(str string) bool {
	return regexp.MustCompile("^[0-9]+$").MatchString(str)
}
