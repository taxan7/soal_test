package model

type View struct {
	Status       bool        `json:"status"`
	Code         string      `json:"code"`
	Message      string      `json:"message"`
	ErrorMessage interface{} `json:"error_message"`
	Data         interface{} `json:"data"`
}
