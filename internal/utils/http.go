package utils

import (
	"encoding/json"
	"go-web/internal/models"
	"net/http"
	"strings"
)

func SetHeaderJSON(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
}

func SendRes(res http.ResponseWriter, msg string, status int, data interface{}, err error) {
	SetHeaderJSON(res)

	response := models.Response{
		Message: msg,
		Status: status,
		Data: data,
		Err: err,
	}

	res.WriteHeader(status)
	json.NewEncoder(res).Encode(response)
}

func GetClientIP(req *http.Request) string {
	if xff := req.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}

	if xri := req.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	ip := req.RemoteAddr
	ip = strings.Split(ip, ":")[0]
	return ip
}