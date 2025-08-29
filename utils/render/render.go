package render

import (
	"encoding/json"
	"net/http"

	"github.com/umarkotak/ytkidd-api/model"
)

type ResponseBody struct {
	Data    any       `json:"data"`
	Success bool      `json:"success"`
	Error   ErrorData `json:"error"`
}

type ErrorData struct {
	Code          string `json:"code,omitempty"`
	Message       string `json:"message,omitempty"`
	InternalError string `json:"internal_error,omitempty"`
}

func Response(w http.ResponseWriter, r *http.Request, statusCode int, data any) {
	if data == nil {
		data = map[string]any{}
	}

	res := ResponseBody{
		Data:    data,
		Success: true,
		Error:   ErrorData{},
	}
	b, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(b)
}

func Error(w http.ResponseWriter, r *http.Request, err error, customMessage string) {
	jxErr, ok := model.ERR_MAP[err]
	if !ok {
		jxErr = model.ERR_MAP[model.ErrUnprocessableEntity]
	}

	message := jxErr.EN
	if customMessage != "" {
		message = customMessage
	}

	res := ResponseBody{
		Data:    map[string]string{},
		Success: false,
		Error: ErrorData{
			Code:          jxErr.ErrorCode,
			Message:       message,
			InternalError: err.Error(),
		},
	}
	b, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(jxErr.HttpCode)
	w.Write(b)
}

func RawError(w http.ResponseWriter, r *http.Request, statusCode int, err error, errCode, errMessage string) {
	message := err.Error()
	if errMessage != "" {
		message = errMessage
	}

	res := ResponseBody{
		Data:    map[string]any{},
		Success: true,
		Error: ErrorData{
			Code:    errCode,
			Message: message,
		},
	}
	b, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(b)
}

func SetCorsHeaders(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
	w.Header().Add(
		"Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-App-Session",
	)
}
