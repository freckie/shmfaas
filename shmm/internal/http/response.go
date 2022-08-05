package http

import (
	"encoding/json"
	"net/http"

	klog "k8s.io/klog/v2"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// ResponseOK makes a 200 response.
func ResponseOK(w http.ResponseWriter, r *http.Request, msg string, data interface{}) {
	klog.InfofDepth(1, "[200] %s %s", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&Response{
		StatusCode: 200,
		Message:    msg,
		Data:       data,
	})
}

// ResponseError makes an error response.
func ResponseError(w http.ResponseWriter, r *http.Request, errorCode int, errorMsg string) {
	klog.ErrorfDepth(1, "[%d] %s %s :: %s", errorCode, r.Method, r.URL.Path, errorMsg)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(&Response{
		StatusCode: errorCode,
		Message:    errorMsg,
		Data:       nil,
	})
}

// Response404 handles 404 errors.
func Response404(w http.ResponseWriter, r *http.Request) {
	klog.ErrorfDepth(1, "[404] %s %s", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(&Response{
		StatusCode: 404,
		Message:    "404 Not Found",
		Data:       nil,
	})
}
