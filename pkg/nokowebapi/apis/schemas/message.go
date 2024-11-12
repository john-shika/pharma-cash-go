package schemas

import (
	"nokowebapi/nokocore"
)

type MessageBody struct {
	StatusOk   bool   `mapstructure:"status_ok" json:"statusOk,required" yaml:"status_ok"`
	StatusCode int    `mapstructure:"status_code" json:"statusCode,required" yaml:"status_code"`
	Status     string `mapstructure:"status" json:"status,required" yaml:"status"`
	Message    string `mapstructure:"message" json:"message,required" yaml:"message"`
	Data       any    `mapstructure:"data" json:"data,required" yaml:"data"`
}

func NewMessageBody(statusOk bool, statusCode int, status string, message string, data any) *MessageBody {
	return &MessageBody{
		StatusOk:   statusOk,
		StatusCode: statusCode,
		Status:     status,
		Message:    message,
		Data:       data,
	}
}

func NewMessageBodyOk(message string, data any) *MessageBody {
	status := nokocore.Default[nokocore.HttpStatusCodeValue]().FromCode(nokocore.HttpStatusCodeOk)
	return NewMessageBody(false, nokocore.HttpStatusCodeOk, string(status), message, data)
}

func NewMessageBodyCreated(message string, data any) *MessageBody {
	status := nokocore.Default[nokocore.HttpStatusCodeValue]().FromCode(nokocore.HttpStatusCodeCreated)
	return NewMessageBody(false, nokocore.HttpStatusCodeCreated, string(status), message, data)
}

func NewMessageBodyUnauthorized(message string, data any) *MessageBody {
	status := nokocore.Default[nokocore.HttpStatusCodeValue]().FromCode(nokocore.HttpStatusCodeUnauthorized)
	return NewMessageBody(false, nokocore.HttpStatusCodeUnauthorized, string(status), message, data)
}

func NewMessageBodyBadRequest(message string, data any) *MessageBody {
	status := nokocore.Default[nokocore.HttpStatusCodeValue]().FromCode(nokocore.HttpStatusCodeBadRequest)
	return NewMessageBody(false, nokocore.HttpStatusCodeBadRequest, string(status), message, data)
}

func NewMessageBodyNotFound(message string, data any) *MessageBody {
	status := nokocore.Default[nokocore.HttpStatusCodeValue]().FromCode(nokocore.HttpStatusCodeNotFound)
	return NewMessageBody(false, nokocore.HttpStatusCodeNotFound, string(status), message, data)
}

func NewMessageBodyInternalServerError(message string, data any) *MessageBody {
	status := nokocore.Default[nokocore.HttpStatusCodeValue]().FromCode(nokocore.HttpStatusCodeInternalServerError)
	return NewMessageBody(false, nokocore.HttpStatusCodeInternalServerError, string(status), message, data)
}
