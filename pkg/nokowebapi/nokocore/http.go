package nokocore

type HttpMethod string

const (
	HttpMethodGet     HttpMethod = "GET"
	HttpMethodPost    HttpMethod = "POST"
	HttpMethodPut     HttpMethod = "PUT"
	HttpMethodPatch   HttpMethod = "PATCH"
	HttpMethodDelete  HttpMethod = "DELETE"
	HttpMethodOptions HttpMethod = "OPTIONS"
	HttpMethodHead    HttpMethod = "HEAD"
	HttpMethodTrace   HttpMethod = "TRACE"
)

type HttpStatusCode int
type HttpStatusCodeValue string

const (
	HttpStatusCodeContinue                      HttpStatusCode = 100
	HttpStatusCodeSwitchingProtocols            HttpStatusCode = 101
	HttpStatusCodeProcessing                    HttpStatusCode = 102
	HttpStatusCodeEarlyHints                    HttpStatusCode = 103
	HttpStatusCodeOk                            HttpStatusCode = 200
	HttpStatusCodeCreated                       HttpStatusCode = 201
	HttpStatusCodeAccepted                      HttpStatusCode = 202
	HttpStatusCodeNonAuthoritativeInformation   HttpStatusCode = 203
	HttpStatusCodeNoContent                     HttpStatusCode = 204
	HttpStatusCodeResetContent                  HttpStatusCode = 205
	HttpStatusCodePartialContent                HttpStatusCode = 206
	HttpStatusCodeMultiStatus                   HttpStatusCode = 207
	HttpStatusCodeAlreadyReported               HttpStatusCode = 208
	HttpStatusCodeImUsed                        HttpStatusCode = 226
	HttpStatusCodeMultipleChoices               HttpStatusCode = 300
	HttpStatusCodeMovedPermanently              HttpStatusCode = 301
	HttpStatusCodeFound                         HttpStatusCode = 302
	HttpStatusCodeSeeOther                      HttpStatusCode = 303
	HttpStatusCodeNotModified                   HttpStatusCode = 304
	HttpStatusCodeUseProxy                      HttpStatusCode = 305
	HttpStatusCodeUnused                        HttpStatusCode = 306
	HttpStatusCodeTemporaryRedirect             HttpStatusCode = 307
	HttpStatusCodePermanentRedirect             HttpStatusCode = 308
	HttpStatusCodeBadRequest                    HttpStatusCode = 400
	HttpStatusCodeUnauthorized                  HttpStatusCode = 401
	HttpStatusCodePaymentRequired               HttpStatusCode = 402
	HttpStatusCodeForbidden                     HttpStatusCode = 403
	HttpStatusCodeNotFound                      HttpStatusCode = 404
	HttpStatusCodeMethodNotAllowed              HttpStatusCode = 405
	HttpStatusCodeNotAcceptable                 HttpStatusCode = 406
	HttpStatusCodeProxyAuthenticationRequired   HttpStatusCode = 407
	HttpStatusCodeRequestTimeout                HttpStatusCode = 408
	HttpStatusCodeConflict                      HttpStatusCode = 409
	HttpStatusCodeGone                          HttpStatusCode = 410
	HttpStatusCodeLengthRequired                HttpStatusCode = 411
	HttpStatusCodePreconditionFailed            HttpStatusCode = 412
	HttpStatusCodePayloadTooLarge               HttpStatusCode = 413
	HttpStatusCodeRequestUriTooLong             HttpStatusCode = 414
	HttpStatusCodeUnsupportedMediaType          HttpStatusCode = 415
	HttpStatusCodeRequestedRangeNotSatisfiable  HttpStatusCode = 416
	HttpStatusCodeExpectationFailed             HttpStatusCode = 417
	HttpStatusCodeImATeapot                     HttpStatusCode = 418
	HttpStatusCodeInsufficientSpaceOnResource   HttpStatusCode = 419
	HttpStatusCodeMethodFailure                 HttpStatusCode = 420
	HttpStatusCodeMisdirectedRequest            HttpStatusCode = 421
	HttpStatusCodeUnprocessableEntity           HttpStatusCode = 422
	HttpStatusCodeLocked                        HttpStatusCode = 423
	HttpStatusCodeFailedDependency              HttpStatusCode = 424
	HttpStatusCodeUpgradeRequired               HttpStatusCode = 426
	HttpStatusCodePreconditionRequired          HttpStatusCode = 428
	HttpStatusCodeTooManyRequests               HttpStatusCode = 429
	HttpStatusCodeRequestHeaderFieldsTooLarge   HttpStatusCode = 431
	HttpStatusCodeUnavailableForLegalReasons    HttpStatusCode = 451
	HttpStatusCodeInternalServerError           HttpStatusCode = 500
	HttpStatusCodeNotImplemented                HttpStatusCode = 501
	HttpStatusCodeBadGateway                    HttpStatusCode = 502
	HttpStatusCodeServiceUnavailable            HttpStatusCode = 503
	HttpStatusCodeGatewayTimeout                HttpStatusCode = 504
	HttpStatusCodeHttpVersionNotSupported       HttpStatusCode = 505
	HttpStatusCodeVariantAlsoNegotiates         HttpStatusCode = 506
	HttpStatusCodeInsufficientStorage           HttpStatusCode = 507
	HttpStatusCodeLoopDetected                  HttpStatusCode = 508
	HttpStatusCodeNotExtended                   HttpStatusCode = 510
	HttpStatusCodeNetworkAuthenticationRequired HttpStatusCode = 511
)

const (
	HttpStatusCodeValueContinue                      HttpStatusCodeValue = "CONTINUE"
	HttpStatusCodeValueSwitchingProtocols            HttpStatusCodeValue = "SWITCHING_PROTOCOLS"
	HttpStatusCodeValueProcessing                    HttpStatusCodeValue = "PROCESSING"
	HttpStatusCodeValueEarlyHints                    HttpStatusCodeValue = "EARLY_HINTS"
	HttpStatusCodeValueOk                            HttpStatusCodeValue = "OK"
	HttpStatusCodeValueCreated                       HttpStatusCodeValue = "CREATED"
	HttpStatusCodeValueAccepted                      HttpStatusCodeValue = "ACCEPTED"
	HttpStatusCodeValueNonAuthoritativeInformation   HttpStatusCodeValue = "NON_AUTHORITATIVE_INFORMATION"
	HttpStatusCodeValueNoContent                     HttpStatusCodeValue = "NO_CONTENT"
	HttpStatusCodeValueResetContent                  HttpStatusCodeValue = "RESET_CONTENT"
	HttpStatusCodeValuePartialContent                HttpStatusCodeValue = "PARTIAL_CONTENT"
	HttpStatusCodeValueMultiStatus                   HttpStatusCodeValue = "MULTI_STATUS"
	HttpStatusCodeValueAlreadyReported               HttpStatusCodeValue = "ALREADY_REPORTED"
	HttpStatusCodeValueImUsed                        HttpStatusCodeValue = "IM_USED"
	HttpStatusCodeValueMultipleChoices               HttpStatusCodeValue = "MULTI_CHOICES"
	HttpStatusCodeValueMovedPermanently              HttpStatusCodeValue = "MOVED_PERMANENTLY"
	HttpStatusCodeValueFound                         HttpStatusCodeValue = "FOUND"
	HttpStatusCodeValueSeeOther                      HttpStatusCodeValue = "SEE_OTHER"
	HttpStatusCodeValueNotModified                   HttpStatusCodeValue = "NOT_MODIFIED"
	HttpStatusCodeValueUseProxy                      HttpStatusCodeValue = "USE_PROXY"
	HttpStatusCodeValueUnused                        HttpStatusCodeValue = "UNUSED"
	HttpStatusCodeValueTemporaryRedirect             HttpStatusCodeValue = "TEMPORARY_REDIRECT"
	HttpStatusCodeValuePermanentRedirect             HttpStatusCodeValue = "PERMANENT_REDIRECT"
	HttpStatusCodeValueBadRequest                    HttpStatusCodeValue = "BAD_REQUEST"
	HttpStatusCodeValueUnauthorized                  HttpStatusCodeValue = "UNAUTHORIZED"
	HttpStatusCodeValuePaymentRequired               HttpStatusCodeValue = "PAYMENT_REQUIRED"
	HttpStatusCodeValueForbidden                     HttpStatusCodeValue = "FORBIDDEN"
	HttpStatusCodeValueNotFound                      HttpStatusCodeValue = "NOT_FOUND"
	HttpStatusCodeValueMethodNotAllowed              HttpStatusCodeValue = "METHOD_NOT_ALLOWED"
	HttpStatusCodeValueNotAcceptable                 HttpStatusCodeValue = "NOT_ACCEPTABLE"
	HttpStatusCodeValueProxyAuthenticationRequired   HttpStatusCodeValue = "PROXY_AUTHENTICATION_REQUIRED"
	HttpStatusCodeValueRequestTimeout                HttpStatusCodeValue = "REQUEST_TIMEOUT"
	HttpStatusCodeValueConflict                      HttpStatusCodeValue = "CONFLICT"
	HttpStatusCodeValueGone                          HttpStatusCodeValue = "GONE"
	HttpStatusCodeValueLengthRequired                HttpStatusCodeValue = "LENGTH_REQUIRED"
	HttpStatusCodeValuePreconditionFailed            HttpStatusCodeValue = "PRECONDITION_FAILED"
	HttpStatusCodeValuePayloadTooLarge               HttpStatusCodeValue = "PAYLOAD_TOO_LARGE"
	HttpStatusCodeValueRequestUriTooLong             HttpStatusCodeValue = "REQUEST_URI_TOO_LONG"
	HttpStatusCodeValueUnsupportedMediaType          HttpStatusCodeValue = "UNSUPPORTED_MEDIA_TYPE"
	HttpStatusCodeValueRequestedRangeNotSatisfiable  HttpStatusCodeValue = "REQUESTED_RANGE_NOT_SATISFIABLE"
	HttpStatusCodeValueExpectationFailed             HttpStatusCodeValue = "EXPECTATION_FAILED"
	HttpStatusCodeValueImATeapot                     HttpStatusCodeValue = "IM_A_TEAPOT"
	HttpStatusCodeValueInsufficientSpaceOnResource   HttpStatusCodeValue = "INSUFFICIENT_SPACE_ON_RESOURCE"
	HttpStatusCodeValueMethodFailure                 HttpStatusCodeValue = "METHOD_FAILURE"
	HttpStatusCodeValueMisdirectedRequest            HttpStatusCodeValue = "MISDIRECTED_REQUEST"
	HttpStatusCodeValueUnprocessableEntity           HttpStatusCodeValue = "UNPROCESSABLE_ENTITY"
	HttpStatusCodeValueLocked                        HttpStatusCodeValue = "LOCKED"
	HttpStatusCodeValueFailedDependency              HttpStatusCodeValue = "FAILED_DEPENDENCY"
	HttpStatusCodeValueUpgradeRequired               HttpStatusCodeValue = "UPGRADE_REQUIRED"
	HttpStatusCodeValuePreconditionRequired          HttpStatusCodeValue = "PRECONDITION_REQUIRED"
	HttpStatusCodeValueTooManyRequests               HttpStatusCodeValue = "TOO_MANY_REQUESTS"
	HttpStatusCodeValueRequestHeaderFieldsTooLarge   HttpStatusCodeValue = "REQUEST_HEADER_FIELDS_TOO_LARGE"
	HttpStatusCodeValueUnavailableForLegalReasons    HttpStatusCodeValue = "UNAVAILABLE_FOR_LEGAL_REASONS"
	HttpStatusCodeValueInternalServerError           HttpStatusCodeValue = "INTERNAL_SERVER_ERROR"
	HttpStatusCodeValueNotImplemented                HttpStatusCodeValue = "NOT_IMPLEMENTED"
	HttpStatusCodeValueBadGateway                    HttpStatusCodeValue = "BAD_GATEWAY"
	HttpStatusCodeValueServiceUnavailable            HttpStatusCodeValue = "SERVICE_UNAVAILABLE"
	HttpStatusCodeValueGatewayTimeout                HttpStatusCodeValue = "GATEWAY_TIMEOUT"
	HttpStatusCodeValueHttpVersionNotSupported       HttpStatusCodeValue = "HTTP_VERSION_NOT_SUPPORTED"
	HttpStatusCodeValueVariantAlsoNegotiates         HttpStatusCodeValue = "VARIANT_ALSO_NEGOTIATES"
	HttpStatusCodeValueInsufficientStorage           HttpStatusCodeValue = "INSUFFICIENT_STORAGE"
	HttpStatusCodeValueLoopDetected                  HttpStatusCodeValue = "LOOP_DETECTED"
	HttpStatusCodeValueNotExtended                   HttpStatusCodeValue = "NOT_EXTENDED"
	HttpStatusCodeValueNetworkAuthenticationRequired HttpStatusCodeValue = "NETWORK_AUTHENTICATION_REQUIRED"
)

func GetValueFromHttpStatusCode(code HttpStatusCode) HttpStatusCodeValue {
	switch code {
	case HttpStatusCodeContinue:
		return HttpStatusCodeValueContinue
	case HttpStatusCodeSwitchingProtocols:
		return HttpStatusCodeValueSwitchingProtocols
	case HttpStatusCodeProcessing:
		return HttpStatusCodeValueProcessing
	case HttpStatusCodeEarlyHints:
		return HttpStatusCodeValueEarlyHints
	case HttpStatusCodeOk:
		return HttpStatusCodeValueOk
	case HttpStatusCodeCreated:
		return HttpStatusCodeValueCreated
	case HttpStatusCodeAccepted:
		return HttpStatusCodeValueAccepted
	case HttpStatusCodeNonAuthoritativeInformation:
		return HttpStatusCodeValueNonAuthoritativeInformation
	case HttpStatusCodeNoContent:
		return HttpStatusCodeValueNoContent
	case HttpStatusCodeResetContent:
		return HttpStatusCodeValueResetContent
	case HttpStatusCodePartialContent:
		return HttpStatusCodeValuePartialContent
	case HttpStatusCodeMultiStatus:
		return HttpStatusCodeValueMultiStatus
	case HttpStatusCodeAlreadyReported:
		return HttpStatusCodeValueAlreadyReported
	case HttpStatusCodeImUsed:
		return HttpStatusCodeValueImUsed
	case HttpStatusCodeMultipleChoices:
		return HttpStatusCodeValueMultipleChoices
	case HttpStatusCodeMovedPermanently:
		return HttpStatusCodeValueMovedPermanently
	case HttpStatusCodeFound:
		return HttpStatusCodeValueFound
	case HttpStatusCodeSeeOther:
		return HttpStatusCodeValueSeeOther
	case HttpStatusCodeNotModified:
		return HttpStatusCodeValueNotModified
	case HttpStatusCodeUseProxy:
		return HttpStatusCodeValueUseProxy
	case HttpStatusCodeUnused:
		return HttpStatusCodeValueUnused
	case HttpStatusCodeTemporaryRedirect:
		return HttpStatusCodeValueTemporaryRedirect
	case HttpStatusCodePermanentRedirect:
		return HttpStatusCodeValuePermanentRedirect
	case HttpStatusCodeBadRequest:
		return HttpStatusCodeValueBadRequest
	case HttpStatusCodeUnauthorized:
		return HttpStatusCodeValueUnauthorized
	case HttpStatusCodePaymentRequired:
		return HttpStatusCodeValuePaymentRequired
	case HttpStatusCodeForbidden:
		return HttpStatusCodeValueForbidden
	case HttpStatusCodeNotFound:
		return HttpStatusCodeValueNotFound
	case HttpStatusCodeMethodNotAllowed:
		return HttpStatusCodeValueMethodNotAllowed
	case HttpStatusCodeNotAcceptable:
		return HttpStatusCodeValueNotAcceptable
	case HttpStatusCodeProxyAuthenticationRequired:
		return HttpStatusCodeValueProxyAuthenticationRequired
	case HttpStatusCodeRequestTimeout:
		return HttpStatusCodeValueRequestTimeout
	case HttpStatusCodeConflict:
		return HttpStatusCodeValueConflict
	case HttpStatusCodeGone:
		return HttpStatusCodeValueGone
	case HttpStatusCodeLengthRequired:
		return HttpStatusCodeValueLengthRequired
	case HttpStatusCodePreconditionFailed:
		return HttpStatusCodeValuePreconditionFailed
	case HttpStatusCodePayloadTooLarge:
		return HttpStatusCodeValuePayloadTooLarge
	case HttpStatusCodeRequestUriTooLong:
		return HttpStatusCodeValueRequestUriTooLong
	case HttpStatusCodeUnsupportedMediaType:
		return HttpStatusCodeValueUnsupportedMediaType
	case HttpStatusCodeRequestedRangeNotSatisfiable:
		return HttpStatusCodeValueRequestedRangeNotSatisfiable
	case HttpStatusCodeExpectationFailed:
		return HttpStatusCodeValueExpectationFailed
	case HttpStatusCodeImATeapot:
		return HttpStatusCodeValueImATeapot
	case HttpStatusCodeInsufficientSpaceOnResource:
		return HttpStatusCodeValueInsufficientSpaceOnResource
	case HttpStatusCodeMethodFailure:
		return HttpStatusCodeValueMethodFailure
	case HttpStatusCodeMisdirectedRequest:
		return HttpStatusCodeValueMisdirectedRequest
	case HttpStatusCodeUnprocessableEntity:
		return HttpStatusCodeValueUnprocessableEntity
	case HttpStatusCodeLocked:
		return HttpStatusCodeValueLocked
	case HttpStatusCodeFailedDependency:
		return HttpStatusCodeValueFailedDependency
	case HttpStatusCodeUpgradeRequired:
		return HttpStatusCodeValueUpgradeRequired
	case HttpStatusCodePreconditionRequired:
		return HttpStatusCodeValuePreconditionRequired
	case HttpStatusCodeTooManyRequests:
		return HttpStatusCodeValueTooManyRequests
	case HttpStatusCodeRequestHeaderFieldsTooLarge:
		return HttpStatusCodeValueRequestHeaderFieldsTooLarge
	case HttpStatusCodeUnavailableForLegalReasons:
		return HttpStatusCodeValueUnavailableForLegalReasons
	case HttpStatusCodeInternalServerError:
		return HttpStatusCodeValueInternalServerError
	case HttpStatusCodeNotImplemented:
		return HttpStatusCodeValueNotImplemented
	case HttpStatusCodeBadGateway:
		return HttpStatusCodeValueBadGateway
	case HttpStatusCodeServiceUnavailable:
		return HttpStatusCodeValueServiceUnavailable
	case HttpStatusCodeGatewayTimeout:
		return HttpStatusCodeValueGatewayTimeout
	case HttpStatusCodeHttpVersionNotSupported:
		return HttpStatusCodeValueHttpVersionNotSupported
	case HttpStatusCodeVariantAlsoNegotiates:
		return HttpStatusCodeValueVariantAlsoNegotiates
	case HttpStatusCodeInsufficientStorage:
		return HttpStatusCodeValueInsufficientStorage
	case HttpStatusCodeLoopDetected:
		return HttpStatusCodeValueLoopDetected
	case HttpStatusCodeNotExtended:
		return HttpStatusCodeValueNotExtended
	case HttpStatusCodeNetworkAuthenticationRequired:
		return HttpStatusCodeValueNetworkAuthenticationRequired
	default:
		panic("invalid http status code")
	}
}

func ParseHttpStatusCode(value string) HttpStatusCode {
	switch HttpStatusCodeValue(ToSnakeCaseUpper(value)) {
	case HttpStatusCodeValueContinue:
		return HttpStatusCodeContinue
	case HttpStatusCodeValueSwitchingProtocols:
		return HttpStatusCodeSwitchingProtocols
	case HttpStatusCodeValueProcessing:
		return HttpStatusCodeProcessing
	case HttpStatusCodeValueEarlyHints:
		return HttpStatusCodeEarlyHints
	case HttpStatusCodeValueOk:
		return HttpStatusCodeOk
	case HttpStatusCodeValueCreated:
		return HttpStatusCodeCreated
	case HttpStatusCodeValueAccepted:
		return HttpStatusCodeAccepted
	case HttpStatusCodeValueNonAuthoritativeInformation:
		return HttpStatusCodeNonAuthoritativeInformation
	case HttpStatusCodeValueNoContent:
		return HttpStatusCodeNoContent
	case HttpStatusCodeValueResetContent:
		return HttpStatusCodeResetContent
	case HttpStatusCodeValuePartialContent:
		return HttpStatusCodePartialContent
	case HttpStatusCodeValueMultiStatus:
		return HttpStatusCodeMultiStatus
	case HttpStatusCodeValueAlreadyReported:
		return HttpStatusCodeAlreadyReported
	case HttpStatusCodeValueImUsed:
		return HttpStatusCodeImUsed
	case HttpStatusCodeValueMultipleChoices:
		return HttpStatusCodeMultipleChoices
	case HttpStatusCodeValueMovedPermanently:
		return HttpStatusCodeMovedPermanently
	case HttpStatusCodeValueFound:
		return HttpStatusCodeFound
	case HttpStatusCodeValueSeeOther:
		return HttpStatusCodeSeeOther
	case HttpStatusCodeValueNotModified:
		return HttpStatusCodeNotModified
	case HttpStatusCodeValueUseProxy:
		return HttpStatusCodeUseProxy
	case HttpStatusCodeValueUnused:
		return HttpStatusCodeUnused
	case HttpStatusCodeValueTemporaryRedirect:
		return HttpStatusCodeTemporaryRedirect
	case HttpStatusCodeValuePermanentRedirect:
		return HttpStatusCodePermanentRedirect
	case HttpStatusCodeValueBadRequest:
		return HttpStatusCodeBadRequest
	case HttpStatusCodeValueUnauthorized:
		return HttpStatusCodeUnauthorized
	case HttpStatusCodeValuePaymentRequired:
		return HttpStatusCodePaymentRequired
	case HttpStatusCodeValueForbidden:
		return HttpStatusCodeForbidden
	case HttpStatusCodeValueNotFound:
		return HttpStatusCodeNotFound
	case HttpStatusCodeValueMethodNotAllowed:
		return HttpStatusCodeMethodNotAllowed
	case HttpStatusCodeValueNotAcceptable:
		return HttpStatusCodeNotAcceptable
	case HttpStatusCodeValueProxyAuthenticationRequired:
		return HttpStatusCodeProxyAuthenticationRequired
	case HttpStatusCodeValueRequestTimeout:
		return HttpStatusCodeRequestTimeout
	case HttpStatusCodeValueConflict:
		return HttpStatusCodeConflict
	case HttpStatusCodeValueGone:
		return HttpStatusCodeGone
	case HttpStatusCodeValueLengthRequired:
		return HttpStatusCodeLengthRequired
	case HttpStatusCodeValuePreconditionFailed:
		return HttpStatusCodePreconditionFailed
	case HttpStatusCodeValuePayloadTooLarge:
		return HttpStatusCodePayloadTooLarge
	case HttpStatusCodeValueRequestUriTooLong:
		return HttpStatusCodeRequestUriTooLong
	case HttpStatusCodeValueUnsupportedMediaType:
		return HttpStatusCodeUnsupportedMediaType
	case HttpStatusCodeValueRequestedRangeNotSatisfiable:
		return HttpStatusCodeRequestedRangeNotSatisfiable
	case HttpStatusCodeValueExpectationFailed:
		return HttpStatusCodeExpectationFailed
	case HttpStatusCodeValueImATeapot:
		return HttpStatusCodeImATeapot
	case HttpStatusCodeValueInsufficientSpaceOnResource:
		return HttpStatusCodeInsufficientSpaceOnResource
	case HttpStatusCodeValueMethodFailure:
		return HttpStatusCodeMethodFailure
	case HttpStatusCodeValueMisdirectedRequest:
		return HttpStatusCodeMisdirectedRequest
	case HttpStatusCodeValueUnprocessableEntity:
		return HttpStatusCodeUnprocessableEntity
	case HttpStatusCodeValueLocked:
		return HttpStatusCodeLocked
	case HttpStatusCodeValueFailedDependency:
		return HttpStatusCodeFailedDependency
	case HttpStatusCodeValueUpgradeRequired:
		return HttpStatusCodeUpgradeRequired
	case HttpStatusCodeValuePreconditionRequired:
		return HttpStatusCodePreconditionRequired
	case HttpStatusCodeValueTooManyRequests:
		return HttpStatusCodeTooManyRequests
	case HttpStatusCodeValueRequestHeaderFieldsTooLarge:
		return HttpStatusCodeRequestHeaderFieldsTooLarge
	case HttpStatusCodeValueUnavailableForLegalReasons:
		return HttpStatusCodeUnavailableForLegalReasons
	case HttpStatusCodeValueInternalServerError:
		return HttpStatusCodeInternalServerError
	case HttpStatusCodeValueNotImplemented:
		return HttpStatusCodeNotImplemented
	case HttpStatusCodeValueBadGateway:
		return HttpStatusCodeBadGateway
	case HttpStatusCodeValueServiceUnavailable:
		return HttpStatusCodeServiceUnavailable
	case HttpStatusCodeValueGatewayTimeout:
		return HttpStatusCodeGatewayTimeout
	case HttpStatusCodeValueHttpVersionNotSupported:
		return HttpStatusCodeHttpVersionNotSupported
	case HttpStatusCodeValueVariantAlsoNegotiates:
		return HttpStatusCodeVariantAlsoNegotiates
	case HttpStatusCodeValueInsufficientStorage:
		return HttpStatusCodeInsufficientStorage
	case HttpStatusCodeValueLoopDetected:
		return HttpStatusCodeLoopDetected
	case HttpStatusCodeValueNotExtended:
		return HttpStatusCodeNotExtended
	case HttpStatusCodeValueNetworkAuthenticationRequired:
		return HttpStatusCodeNetworkAuthenticationRequired
	default:
		panic("invalid http status text")
	}
}
