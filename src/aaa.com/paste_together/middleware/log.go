package middleware

import (
	"aaa.com/paste_together/common"
	"log"
	"net/http"
	"time"
)

func LogHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 扩展 ResponseWrite
		customResponseWriter := common.NewCustomResponseWriter(&w)

		// 记录执行时间
		start := time.Now()
		h.ServeHTTP(customResponseWriter, r)
		end := time.Now()
		result := end.Sub(start)
		log.Printf("[%s] %q [%v] [%v]", r.Method, r.URL.String(), customResponseWriter.StatusCode, result)
	}
}
