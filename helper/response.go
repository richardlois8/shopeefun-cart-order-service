package helper

import (
	"encoding/json"
	"net/http"
)

const (
	SUCCESS_MESSSAGE string = "Success"
)

type SuccessResponse struct {
	Success  bool   `json:"success"`
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success  bool   `json:"success"`
	Message string `json:"message"`
}

func HandleResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode >= 400 {
		json.NewEncoder(w).Encode(ErrorResponse{
			Success:  false,
			Message: data.(string),
		})
		return
	}else {
		json.NewEncoder(w).Encode(SuccessResponse{
			Success:  true,
			Message: SUCCESS_MESSSAGE,
			Data: data,
		})
	}
}