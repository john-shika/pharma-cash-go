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

func NewMessageBodyContinue(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeContinue
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodySwitchingProtocols(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeSwitchingProtocols
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyProcessing(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeProcessing
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyEarlyHints(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeEarlyHints
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyOk(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeOk
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyCreated(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeCreated
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyAccepted(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeAccepted
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyNonAuthoritativeInformation(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNonAuthoritativeInformation
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyNoContent(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNoContent
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyResetContent(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeResetContent
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyPartialContent(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodePartialContent
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyMultiStatus(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeMultiStatus
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyAlreadyReported(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeAlreadyReported
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyImUsed(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeImUsed
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyMultipleChoices(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeMultipleChoices
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyMovedPermanently(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeMovedPermanently
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyFound(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeFound
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodySeeOther(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeSeeOther
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyNotModified(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNotModified
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyUseProxy(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUseProxy
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyUnused(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUnused
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyTemporaryRedirect(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeTemporaryRedirect
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyPermanentRedirect(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodePermanentRedirect
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyBadRequest(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeBadRequest
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyUnauthorized(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUnauthorized
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyPaymentRequired(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodePaymentRequired
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyForbidden(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeForbidden
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyNotFound(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNotFound
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyMethodNotAllowed(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeMethodNotAllowed
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyNotAcceptable(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNotAcceptable
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyProxyAuthenticationRequired(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeProxyAuthenticationRequired
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyRequestTimeout(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeRequestTimeout
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyConflict(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeConflict
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyGone(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeGone
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyLengthRequired(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeLengthRequired
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyPreconditionFailed(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodePreconditionFailed
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyPayloadTooLarge(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodePayloadTooLarge
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyRequestUriTooLong(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeRequestUriTooLong
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyUnsupportedMediaType(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUnsupportedMediaType
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyRequestedRangeNotSatisfiable(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeRequestedRangeNotSatisfiable
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyExpectationFailed(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeExpectationFailed
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyImATeapot(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeImATeapot
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyInsufficientSpaceOnResource(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeInsufficientSpaceOnResource
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyMethodFailure(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeMethodFailure
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyMisdirectedRequest(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeMisdirectedRequest
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyUnprocessableEntity(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUnprocessableEntity
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyLocked(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeLocked
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyFailedDependency(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeFailedDependency
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyUpgradeRequired(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUpgradeRequired
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyPreconditionRequired(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodePreconditionRequired
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyTooManyRequests(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeTooManyRequests
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyRequestHeaderFieldsTooLarge(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeRequestHeaderFieldsTooLarge
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyUnavailableForLegalReasons(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUnavailableForLegalReasons
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyInternalServerError(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeInternalServerError
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyNotImplemented(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNotImplemented
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyBadGateway(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeBadGateway
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyServiceUnavailable(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeServiceUnavailable
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyGatewayTimeout(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeGatewayTimeout
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyHttpVersionNotSupported(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeHttpVersionNotSupported
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyVariantAlsoNegotiates(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeVariantAlsoNegotiates
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyInsufficientStorage(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeInsufficientStorage
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyLoopDetected(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeLoopDetected
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyNotExtended(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNotExtended
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}

func NewMessageBodyNetworkAuthenticationRequired(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNetworkAuthenticationRequired
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data)
}
