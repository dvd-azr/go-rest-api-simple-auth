package helpers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   bool
	Code    int
	Message string
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func SuccessJsonResponse(w http.ResponseWriter, data interface{}) {
	var succRes SuccessResponse

	succRes.Success = true
	succRes.Data = data
	outputJSON(w, &succRes, http.StatusOK)
}

func ErrorJsonResponse(w http.ResponseWriter, statusCode int, message string) {
	var errorRes ErrorResponse

	errorRes.Error = true
	errorRes.Code = statusCode
	errorRes.Message = message
	outputJSON(w, &errorRes, statusCode)
}

func outputJSON(w http.ResponseWriter, data interface{}, statusCode int) {

	res, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Menghidari superfluous response
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
