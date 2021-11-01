package http

import "errors"

// OKLink {"code":429,"msg":"API_VISIT_TOO_QUICK","detailMsg":"API_VISIT_TOO_QUICK","data":null}
var (
	ErrVisitLimit = errors.New("visit limit")
	ErrMaxRequest = errors.New("max request")
)
