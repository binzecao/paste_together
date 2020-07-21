package common

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, httpStatusCode int, payload interface{}) {
	bytes, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(httpStatusCode)
	w.Write(bytes)
}

//func ResponseJsonSuccess(w http.ResponseWriter, httpStatusCode int, data interface{}) {
//	payload := make(map[string]interface{})
//	payload["code"] = httpStatusCode
//	payload["msg"] = ""
//	payload["data"] = data
//	ResponseJson(w, http.StatusOK, payload)
//}

func ResponseJsonError(w http.ResponseWriter, err error) {
	ResponseJsonErrorWithData(w, err, struct{}{})
}

func ResponseJsonErrorWithData(w http.ResponseWriter, err error, data interface{}) {
	payload := make(map[string]interface{})
	switch e := err.(type) {
	case Error:
		payload["code"] = e.StatusCode()
		payload["msg"] = e.Error()
		payload["data"] = data
		ResponseJson(w, e.StatusCode(), payload)
	default:
		payload["code"] = http.StatusInternalServerError
		payload["msg"] = e.Error()
		payload["data"] = data
		ResponseJson(w, http.StatusInternalServerError, e.Error())
	}
}
