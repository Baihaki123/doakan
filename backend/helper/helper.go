package helper

import "github.com/go-playground/validator/v10"

// Response => public bisa di akses diluar package
// response => private hanya bisa di akses di dalam package tsb
type Response struct {
	Meta Meta        `json:"meta"` // => untuk merubah huruf besar dari Meta jadi huruf kecil meta
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
