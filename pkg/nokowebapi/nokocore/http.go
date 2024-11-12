package nokocore

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
	HttpStatusCodeValuesContinue                      HttpStatusCodeValue = "CONTINUE"
	HttpStatusCodeValuesSwitchingProtocols            HttpStatusCodeValue = "SWITCHING_PROTOCOLS"
	HttpStatusCodeValuesProcessing                    HttpStatusCodeValue = "PROCESSING"
	HttpStatusCodeValuesEarlyHints                    HttpStatusCodeValue = "EARLY_HINTS"
	HttpStatusCodeValuesOk                            HttpStatusCodeValue = "OK"
	HttpStatusCodeValuesCreated                       HttpStatusCodeValue = "CREATED"
	HttpStatusCodeValuesAccepted                      HttpStatusCodeValue = "ACCEPTED"
	HttpStatusCodeValuesNonAuthoritativeInformation   HttpStatusCodeValue = "NON_AUTHORITATIVE_INFORMATION"
	HttpStatusCodeValuesNoContent                     HttpStatusCodeValue = "NO_CONTENT"
	HttpStatusCodeValuesResetContent                  HttpStatusCodeValue = "RESET_CONTENT"
	HttpStatusCodeValuesPartialContent                HttpStatusCodeValue = "PARTIAL_CONTENT"
	HttpStatusCodeValuesMultiStatus                   HttpStatusCodeValue = "MULTI_STATUS"
	HttpStatusCodeValuesAlreadyReported               HttpStatusCodeValue = "ALREADY_REPORTED"
	HttpStatusCodeValuesImUsed                        HttpStatusCodeValue = "IM_USED"
	HttpStatusCodeValuesMultipleChoices               HttpStatusCodeValue = "MULTI_CHOICES"
	HttpStatusCodeValuesMovedPermanently              HttpStatusCodeValue = "MOVED_PERMANENTLY"
	HttpStatusCodeValuesFound                         HttpStatusCodeValue = "FOUND"
	HttpStatusCodeValuesSeeOther                      HttpStatusCodeValue = "SEE_OTHER"
	HttpStatusCodeValuesNotModified                   HttpStatusCodeValue = "NOT_MODIFIED"
	HttpStatusCodeValuesUseProxy                      HttpStatusCodeValue = "USE_PROXY"
	HttpStatusCodeValuesUnused                        HttpStatusCodeValue = "UNUSED"
	HttpStatusCodeValuesTemporaryRedirect             HttpStatusCodeValue = "TEMPORARY_REDIRECT"
	HttpStatusCodeValuesPermanentRedirect             HttpStatusCodeValue = "PERMANENT_REDIRECT"
	HttpStatusCodeValuesBadRequest                    HttpStatusCodeValue = "BAD_REQUEST"
	HttpStatusCodeValuesUnauthorized                  HttpStatusCodeValue = "UNAUTHORIZED"
	HttpStatusCodeValuesPaymentRequired               HttpStatusCodeValue = "PAYMENT_REQUIRED"
	HttpStatusCodeValuesForbidden                     HttpStatusCodeValue = "FORBIDDEN"
	HttpStatusCodeValuesNotFound                      HttpStatusCodeValue = "NOT_FOUND"
	HttpStatusCodeValuesMethodNotAllowed              HttpStatusCodeValue = "METHOD_NOT_ALLOWED"
	HttpStatusCodeValuesNotAcceptable                 HttpStatusCodeValue = "NOT_ACCEPTABLE"
	HttpStatusCodeValuesProxyAuthenticationRequired   HttpStatusCodeValue = "PROXY_AUTHENTICATION_REQUIRED"
	HttpStatusCodeValuesRequestTimeout                HttpStatusCodeValue = "REQUEST_TIMEOUT"
	HttpStatusCodeValuesConflict                      HttpStatusCodeValue = "CONFLICT"
	HttpStatusCodeValuesGone                          HttpStatusCodeValue = "GONE"
	HttpStatusCodeValuesLengthRequired                HttpStatusCodeValue = "LENGTH_REQUIRED"
	HttpStatusCodeValuesPreconditionFailed            HttpStatusCodeValue = "PRECONDITION_FAILED"
	HttpStatusCodeValuesPayloadTooLarge               HttpStatusCodeValue = "PAYLOAD_TOO_LARGE"
	HttpStatusCodeValuesRequestUriTooLong             HttpStatusCodeValue = "REQUEST_URI_TOO_LONG"
	HttpStatusCodeValuesUnsupportedMediaType          HttpStatusCodeValue = "UNSUPPORTED_MEDIA_TYPE"
	HttpStatusCodeValuesRequestedRangeNotSatisfiable  HttpStatusCodeValue = "REQUESTED_RANGE_NOT_SATISFIABLE"
	HttpStatusCodeValuesExpectationFailed             HttpStatusCodeValue = "EXPECTATION_FAILED"
	HttpStatusCodeValuesImATeapot                     HttpStatusCodeValue = "IM_A_TEAPOT"
	HttpStatusCodeValuesInsufficientSpaceOnResource   HttpStatusCodeValue = "INSUFFICIENT_SPACE_ON_RESOURCE"
	HttpStatusCodeValuesMethodFailure                 HttpStatusCodeValue = "METHOD_FAILURE"
	HttpStatusCodeValuesMisdirectedRequest            HttpStatusCodeValue = "MISDIRECTED_REQUEST"
	HttpStatusCodeValuesUnprocessableEntity           HttpStatusCodeValue = "UNPROCESSABLE_ENTITY"
	HttpStatusCodeValuesLocked                        HttpStatusCodeValue = "LOCKED"
	HttpStatusCodeValuesFailedDependency              HttpStatusCodeValue = "FAILED_DEPENDENCY"
	HttpStatusCodeValuesUpgradeRequired               HttpStatusCodeValue = "UPGRADE_REQUIRED"
	HttpStatusCodeValuesPreconditionRequired          HttpStatusCodeValue = "PRECONDITION_REQUIRED"
	HttpStatusCodeValuesTooManyRequests               HttpStatusCodeValue = "TOO_MANY_REQUESTS"
	HttpStatusCodeValuesRequestHeaderFieldsTooLarge   HttpStatusCodeValue = "REQUEST_HEADER_FIELDS_TOO_LARGE"
	HttpStatusCodeValuesUnavailableForLegalReasons    HttpStatusCodeValue = "UNAVAILABLE_FOR_LEGAL_REASONS"
	HttpStatusCodeValuesInternalServerError           HttpStatusCodeValue = "INTERNAL_SERVER_ERROR"
	HttpStatusCodeValuesNotImplemented                HttpStatusCodeValue = "NOT_IMPLEMENTED"
	HttpStatusCodeValuesBadGateway                    HttpStatusCodeValue = "BAD_GATEWAY"
	HttpStatusCodeValuesServiceUnavailable            HttpStatusCodeValue = "SERVICE_UNAVAILABLE"
	HttpStatusCodeValuesGatewayTimeout                HttpStatusCodeValue = "GATEWAY_TIMEOUT"
	HttpStatusCodeValuesHttpVersionNotSupported       HttpStatusCodeValue = "HTTP_VERSION_NOT_SUPPORTED"
	HttpStatusCodeValuesVariantAlsoNegotiates         HttpStatusCodeValue = "VARIANT_ALSO_NEGOTIATES"
	HttpStatusCodeValuesInsufficientStorage           HttpStatusCodeValue = "INSUFFICIENT_STORAGE"
	HttpStatusCodeValuesLoopDetected                  HttpStatusCodeValue = "LOOP_DETECTED"
	HttpStatusCodeValuesNotExtended                   HttpStatusCodeValue = "NOT_EXTENDED"
	HttpStatusCodeValuesNetworkAuthenticationRequired HttpStatusCodeValue = "NETWORK_AUTHENTICATION_REQUIRED"
)

func GetValueFromHttpStatusCode(code HttpStatusCode) HttpStatusCodeValue {
	switch code {
	case HttpStatusCodeContinue:
		return HttpStatusCodeValuesContinue
	case HttpStatusCodeSwitchingProtocols:
		return HttpStatusCodeValuesSwitchingProtocols
	case HttpStatusCodeProcessing:
		return HttpStatusCodeValuesProcessing
	case HttpStatusCodeEarlyHints:
		return HttpStatusCodeValuesEarlyHints
	case HttpStatusCodeOk:
		return HttpStatusCodeValuesOk
	case HttpStatusCodeCreated:
		return HttpStatusCodeValuesCreated
	case HttpStatusCodeAccepted:
		return HttpStatusCodeValuesAccepted
	case HttpStatusCodeNonAuthoritativeInformation:
		return HttpStatusCodeValuesNonAuthoritativeInformation
	case HttpStatusCodeNoContent:
		return HttpStatusCodeValuesNoContent
	case HttpStatusCodeResetContent:
		return HttpStatusCodeValuesResetContent
	case HttpStatusCodePartialContent:
		return HttpStatusCodeValuesPartialContent
	case HttpStatusCodeMultiStatus:
		return HttpStatusCodeValuesMultiStatus
	case HttpStatusCodeAlreadyReported:
		return HttpStatusCodeValuesAlreadyReported
	case HttpStatusCodeImUsed:
		return HttpStatusCodeValuesImUsed
	case HttpStatusCodeMultipleChoices:
		return HttpStatusCodeValuesMultipleChoices
	case HttpStatusCodeMovedPermanently:
		return HttpStatusCodeValuesMovedPermanently
	case HttpStatusCodeFound:
		return HttpStatusCodeValuesFound
	case HttpStatusCodeSeeOther:
		return HttpStatusCodeValuesSeeOther
	case HttpStatusCodeNotModified:
		return HttpStatusCodeValuesNotModified
	case HttpStatusCodeUseProxy:
		return HttpStatusCodeValuesUseProxy
	case HttpStatusCodeUnused:
		return HttpStatusCodeValuesUnused
	case HttpStatusCodeTemporaryRedirect:
		return HttpStatusCodeValuesTemporaryRedirect
	case HttpStatusCodePermanentRedirect:
		return HttpStatusCodeValuesPermanentRedirect
	case HttpStatusCodeBadRequest:
		return HttpStatusCodeValuesBadRequest
	case HttpStatusCodeUnauthorized:
		return HttpStatusCodeValuesUnauthorized
	case HttpStatusCodePaymentRequired:
		return HttpStatusCodeValuesPaymentRequired
	case HttpStatusCodeForbidden:
		return HttpStatusCodeValuesForbidden
	case HttpStatusCodeNotFound:
		return HttpStatusCodeValuesNotFound
	case HttpStatusCodeMethodNotAllowed:
		return HttpStatusCodeValuesMethodNotAllowed
	case HttpStatusCodeNotAcceptable:
		return HttpStatusCodeValuesNotAcceptable
	case HttpStatusCodeProxyAuthenticationRequired:
		return HttpStatusCodeValuesProxyAuthenticationRequired
	case HttpStatusCodeRequestTimeout:
		return HttpStatusCodeValuesRequestTimeout
	case HttpStatusCodeConflict:
		return HttpStatusCodeValuesConflict
	case HttpStatusCodeGone:
		return HttpStatusCodeValuesGone
	case HttpStatusCodeLengthRequired:
		return HttpStatusCodeValuesLengthRequired
	case HttpStatusCodePreconditionFailed:
		return HttpStatusCodeValuesPreconditionFailed
	case HttpStatusCodePayloadTooLarge:
		return HttpStatusCodeValuesPayloadTooLarge
	case HttpStatusCodeRequestUriTooLong:
		return HttpStatusCodeValuesRequestUriTooLong
	case HttpStatusCodeUnsupportedMediaType:
		return HttpStatusCodeValuesUnsupportedMediaType
	case HttpStatusCodeRequestedRangeNotSatisfiable:
		return HttpStatusCodeValuesRequestedRangeNotSatisfiable
	case HttpStatusCodeExpectationFailed:
		return HttpStatusCodeValuesExpectationFailed
	case HttpStatusCodeImATeapot:
		return HttpStatusCodeValuesImATeapot
	case HttpStatusCodeInsufficientSpaceOnResource:
		return HttpStatusCodeValuesInsufficientSpaceOnResource
	case HttpStatusCodeMethodFailure:
		return HttpStatusCodeValuesMethodFailure
	case HttpStatusCodeMisdirectedRequest:
		return HttpStatusCodeValuesMisdirectedRequest
	case HttpStatusCodeUnprocessableEntity:
		return HttpStatusCodeValuesUnprocessableEntity
	case HttpStatusCodeLocked:
		return HttpStatusCodeValuesLocked
	case HttpStatusCodeFailedDependency:
		return HttpStatusCodeValuesFailedDependency
	case HttpStatusCodeUpgradeRequired:
		return HttpStatusCodeValuesUpgradeRequired
	case HttpStatusCodePreconditionRequired:
		return HttpStatusCodeValuesPreconditionRequired
	case HttpStatusCodeTooManyRequests:
		return HttpStatusCodeValuesTooManyRequests
	case HttpStatusCodeRequestHeaderFieldsTooLarge:
		return HttpStatusCodeValuesRequestHeaderFieldsTooLarge
	case HttpStatusCodeUnavailableForLegalReasons:
		return HttpStatusCodeValuesUnavailableForLegalReasons
	case HttpStatusCodeInternalServerError:
		return HttpStatusCodeValuesInternalServerError
	case HttpStatusCodeNotImplemented:
		return HttpStatusCodeValuesNotImplemented
	case HttpStatusCodeBadGateway:
		return HttpStatusCodeValuesBadGateway
	case HttpStatusCodeServiceUnavailable:
		return HttpStatusCodeValuesServiceUnavailable
	case HttpStatusCodeGatewayTimeout:
		return HttpStatusCodeValuesGatewayTimeout
	case HttpStatusCodeHttpVersionNotSupported:
		return HttpStatusCodeValuesHttpVersionNotSupported
	case HttpStatusCodeVariantAlsoNegotiates:
		return HttpStatusCodeValuesVariantAlsoNegotiates
	case HttpStatusCodeInsufficientStorage:
		return HttpStatusCodeValuesInsufficientStorage
	case HttpStatusCodeLoopDetected:
		return HttpStatusCodeValuesLoopDetected
	case HttpStatusCodeNotExtended:
		return HttpStatusCodeValuesNotExtended
	case HttpStatusCodeNetworkAuthenticationRequired:
		return HttpStatusCodeValuesNetworkAuthenticationRequired
	default:
		panic("invalid http status code")
	}
}

func ParseHttpStatusCode(value string) HttpStatusCode {
	switch HttpStatusCodeValue(ToSnakeCaseUpper(value)) {
	case HttpStatusCodeValuesContinue:
		return HttpStatusCodeContinue
	case HttpStatusCodeValuesSwitchingProtocols:
		return HttpStatusCodeSwitchingProtocols
	case HttpStatusCodeValuesProcessing:
		return HttpStatusCodeProcessing
	case HttpStatusCodeValuesEarlyHints:
		return HttpStatusCodeEarlyHints
	case HttpStatusCodeValuesOk:
		return HttpStatusCodeOk
	case HttpStatusCodeValuesCreated:
		return HttpStatusCodeCreated
	case HttpStatusCodeValuesAccepted:
		return HttpStatusCodeAccepted
	case HttpStatusCodeValuesNonAuthoritativeInformation:
		return HttpStatusCodeNonAuthoritativeInformation
	case HttpStatusCodeValuesNoContent:
		return HttpStatusCodeNoContent
	case HttpStatusCodeValuesResetContent:
		return HttpStatusCodeResetContent
	case HttpStatusCodeValuesPartialContent:
		return HttpStatusCodePartialContent
	case HttpStatusCodeValuesMultiStatus:
		return HttpStatusCodeMultiStatus
	case HttpStatusCodeValuesAlreadyReported:
		return HttpStatusCodeAlreadyReported
	case HttpStatusCodeValuesImUsed:
		return HttpStatusCodeImUsed
	case HttpStatusCodeValuesMultipleChoices:
		return HttpStatusCodeMultipleChoices
	case HttpStatusCodeValuesMovedPermanently:
		return HttpStatusCodeMovedPermanently
	case HttpStatusCodeValuesFound:
		return HttpStatusCodeFound
	case HttpStatusCodeValuesSeeOther:
		return HttpStatusCodeSeeOther
	case HttpStatusCodeValuesNotModified:
		return HttpStatusCodeNotModified
	case HttpStatusCodeValuesUseProxy:
		return HttpStatusCodeUseProxy
	case HttpStatusCodeValuesUnused:
		return HttpStatusCodeUnused
	case HttpStatusCodeValuesTemporaryRedirect:
		return HttpStatusCodeTemporaryRedirect
	case HttpStatusCodeValuesPermanentRedirect:
		return HttpStatusCodePermanentRedirect
	case HttpStatusCodeValuesBadRequest:
		return HttpStatusCodeBadRequest
	case HttpStatusCodeValuesUnauthorized:
		return HttpStatusCodeUnauthorized
	case HttpStatusCodeValuesPaymentRequired:
		return HttpStatusCodePaymentRequired
	case HttpStatusCodeValuesForbidden:
		return HttpStatusCodeForbidden
	case HttpStatusCodeValuesNotFound:
		return HttpStatusCodeNotFound
	case HttpStatusCodeValuesMethodNotAllowed:
		return HttpStatusCodeMethodNotAllowed
	case HttpStatusCodeValuesNotAcceptable:
		return HttpStatusCodeNotAcceptable
	case HttpStatusCodeValuesProxyAuthenticationRequired:
		return HttpStatusCodeProxyAuthenticationRequired
	case HttpStatusCodeValuesRequestTimeout:
		return HttpStatusCodeRequestTimeout
	case HttpStatusCodeValuesConflict:
		return HttpStatusCodeConflict
	case HttpStatusCodeValuesGone:
		return HttpStatusCodeGone
	case HttpStatusCodeValuesLengthRequired:
		return HttpStatusCodeLengthRequired
	case HttpStatusCodeValuesPreconditionFailed:
		return HttpStatusCodePreconditionFailed
	case HttpStatusCodeValuesPayloadTooLarge:
		return HttpStatusCodePayloadTooLarge
	case HttpStatusCodeValuesRequestUriTooLong:
		return HttpStatusCodeRequestUriTooLong
	case HttpStatusCodeValuesUnsupportedMediaType:
		return HttpStatusCodeUnsupportedMediaType
	case HttpStatusCodeValuesRequestedRangeNotSatisfiable:
		return HttpStatusCodeRequestedRangeNotSatisfiable
	case HttpStatusCodeValuesExpectationFailed:
		return HttpStatusCodeExpectationFailed
	case HttpStatusCodeValuesImATeapot:
		return HttpStatusCodeImATeapot
	case HttpStatusCodeValuesInsufficientSpaceOnResource:
		return HttpStatusCodeInsufficientSpaceOnResource
	case HttpStatusCodeValuesMethodFailure:
		return HttpStatusCodeMethodFailure
	case HttpStatusCodeValuesMisdirectedRequest:
		return HttpStatusCodeMisdirectedRequest
	case HttpStatusCodeValuesUnprocessableEntity:
		return HttpStatusCodeUnprocessableEntity
	case HttpStatusCodeValuesLocked:
		return HttpStatusCodeLocked
	case HttpStatusCodeValuesFailedDependency:
		return HttpStatusCodeFailedDependency
	case HttpStatusCodeValuesUpgradeRequired:
		return HttpStatusCodeUpgradeRequired
	case HttpStatusCodeValuesPreconditionRequired:
		return HttpStatusCodePreconditionRequired
	case HttpStatusCodeValuesTooManyRequests:
		return HttpStatusCodeTooManyRequests
	case HttpStatusCodeValuesRequestHeaderFieldsTooLarge:
		return HttpStatusCodeRequestHeaderFieldsTooLarge
	case HttpStatusCodeValuesUnavailableForLegalReasons:
		return HttpStatusCodeUnavailableForLegalReasons
	case HttpStatusCodeValuesInternalServerError:
		return HttpStatusCodeInternalServerError
	case HttpStatusCodeValuesNotImplemented:
		return HttpStatusCodeNotImplemented
	case HttpStatusCodeValuesBadGateway:
		return HttpStatusCodeBadGateway
	case HttpStatusCodeValuesServiceUnavailable:
		return HttpStatusCodeServiceUnavailable
	case HttpStatusCodeValuesGatewayTimeout:
		return HttpStatusCodeGatewayTimeout
	case HttpStatusCodeValuesHttpVersionNotSupported:
		return HttpStatusCodeHttpVersionNotSupported
	case HttpStatusCodeValuesVariantAlsoNegotiates:
		return HttpStatusCodeVariantAlsoNegotiates
	case HttpStatusCodeValuesInsufficientStorage:
		return HttpStatusCodeInsufficientStorage
	case HttpStatusCodeValuesLoopDetected:
		return HttpStatusCodeLoopDetected
	case HttpStatusCodeValuesNotExtended:
		return HttpStatusCodeNotExtended
	case HttpStatusCodeValuesNetworkAuthenticationRequired:
		return HttpStatusCodeNetworkAuthenticationRequired
	default:
		panic("invalid http status text")
	}
}
