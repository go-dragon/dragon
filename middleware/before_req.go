package middleware

import (
	"net/http"
)

func beforeReq(r *http.Request, w http.ResponseWriter) {
	// 这里可以做限流的处理
}
