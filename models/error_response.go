package models

type ErrorResponse struct {
	Message string `json:"error"`
	Code    int32  `json:"code"`
}
