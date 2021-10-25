package middleware

import (
	"net/http"
)

func Entry(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// 处理前
		beforeHandle(request, response)
		handler.ServeHTTP(response, request)
		// 处理后
		afterHandle(request, response)
	})
}
