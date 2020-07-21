package common

import "net/http"

// 自定义 ResponseWriter
type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

// 重写 WriteHeader 方法
func (w *CustomResponseWriter) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// 工厂方法
func NewCustomResponseWriter(w *http.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{
		ResponseWriter: *w,
		StatusCode:     http.StatusOK,
	}
}
