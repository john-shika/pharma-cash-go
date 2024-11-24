package schemas

import (
	"nokowebapi/nokocore"
)

type MessageBody struct {
	StatusOk   bool   `mapstructure:"status_ok" json:"statusOk,required"`
	StatusCode int    `mapstructure:"status_code" json:"statusCode,required"`
	Status     string `mapstructure:"status" json:"status,required"`
	Message    string `mapstructure:"message" json:"message,required"`
	Data       any    `mapstructure:"data" json:"data,required"`
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
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodySwitchingProtocols(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeSwitchingProtocols
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyProcessing(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeProcessing
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyEarlyHints(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeEarlyHints
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyOk(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeOk
	return NewMessageBody(true, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyCreated(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeCreated
	return NewMessageBody(true, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyAccepted(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeAccepted
	return NewMessageBody(true, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyNonAuthoritativeInformation(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNonAuthoritativeInformation
	return NewMessageBody(true, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyNoContent(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNoContent
	return NewMessageBody(true, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyResetContent(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeResetContent
	return NewMessageBody(true, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyPartialContent(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodePartialContent
	return NewMessageBody(true, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyMultiStatus(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeMultiStatus
	return NewMessageBody(true, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyAlreadyReported(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeAlreadyReported
	return NewMessageBody(true, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyImUsed(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeImUsed
	return NewMessageBody(true, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyMultipleChoices(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeMultipleChoices
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyMovedPermanently(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeMovedPermanently
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyFound(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeFound
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodySeeOther(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeSeeOther
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyNotModified(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNotModified
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyUseProxy(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUseProxy
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyUnused(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUnused
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyTemporaryRedirect(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeTemporaryRedirect
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyPermanentRedirect(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodePermanentRedirect
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyBadRequest(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeBadRequest
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyUnauthorized(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUnauthorized
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyPaymentRequired(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodePaymentRequired
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyForbidden(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeForbidden
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyNotFound(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNotFound
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyMethodNotAllowed(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeMethodNotAllowed
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyNotAcceptable(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNotAcceptable
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyProxyAuthenticationRequired(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeProxyAuthenticationRequired
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyRequestTimeout(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeRequestTimeout
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyConflict(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeConflict
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyGone(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeGone
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyLengthRequired(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeLengthRequired
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyPreconditionFailed(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodePreconditionFailed
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyPayloadTooLarge(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodePayloadTooLarge
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyRequestUriTooLong(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeRequestUriTooLong
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyUnsupportedMediaType(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUnsupportedMediaType
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyRequestedRangeNotSatisfiable(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeRequestedRangeNotSatisfiable
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyExpectationFailed(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeExpectationFailed
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyImATeapot(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeImATeapot
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyInsufficientSpaceOnResource(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeInsufficientSpaceOnResource
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyMethodFailure(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeMethodFailure
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyMisdirectedRequest(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeMisdirectedRequest
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyUnprocessableEntity(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUnprocessableEntity
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyLocked(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeLocked
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyFailedDependency(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeFailedDependency
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyUpgradeRequired(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUpgradeRequired
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyPreconditionRequired(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodePreconditionRequired
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyTooManyRequests(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeTooManyRequests
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyRequestHeaderFieldsTooLarge(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeRequestHeaderFieldsTooLarge
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyUnavailableForLegalReasons(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeUnavailableForLegalReasons
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyInternalServerError(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeInternalServerError
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyNotImplemented(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNotImplemented
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyBadGateway(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeBadGateway
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyServiceUnavailable(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeServiceUnavailable
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyGatewayTimeout(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeGatewayTimeout
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyHttpVersionNotSupported(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeHttpVersionNotSupported
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyVariantAlsoNegotiates(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeVariantAlsoNegotiates
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyInsufficientStorage(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeInsufficientStorage
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyLoopDetected(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeLoopDetected
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyNotExtended(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNotExtended
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}

func NewMessageBodyNetworkAuthenticationRequired(message string, data any) *MessageBody {
	httpStatusCode := nokocore.HttpStatusCodeNetworkAuthenticationRequired
	return NewMessageBody(false, httpStatusCode.ToInt(), httpStatusCode.ToString(), message, data)
}
