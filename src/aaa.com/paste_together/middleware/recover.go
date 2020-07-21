package middleware

import (
	"aaa.com/paste_together/common"
	"errors"
	"log"
	"net/http"
	"time"
)

func RecoverHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func(start *time.Time, r *http.Request) {
			if r2 := recover(); r2 != nil {
				end := time.Now()
				result := end.Sub(*start)
				_ = result
				log.Printf("Recover from panic: %v", r2)
				//log.Printf("[%s] %q [%v]", r.Method, r.URL.String(), result)
				//log.Printf("[%s] %q [%v] Recover from panic: %v", r.Method, r.URL.String(), result, r)
				common.ResponseJsonError(w, common.AppError{Code: http.StatusInternalServerError, Err: errors.New(http.StatusText(http.StatusInternalServerError))})
			}
		}(&start, r)

		h.ServeHTTP(w, r)

	}
}
