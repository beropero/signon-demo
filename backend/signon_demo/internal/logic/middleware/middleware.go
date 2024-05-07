package middleware

import "signon_demo/internal/service"


type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(&sMiddleware{})
}