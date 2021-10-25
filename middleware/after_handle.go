package middleware

import (
	"net/http"
)

func afterHandle(r *http.Request, w http.ResponseWriter) error{
	return nil
}
