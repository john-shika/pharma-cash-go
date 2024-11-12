package nokocore

type HttpStatusCode int
type HttpStatusCodeValue string

const (
	HttpStatusCodeContinue                      HttpStatusCode = 100
	HttpStatusCodeSwitchingProtocols                           = 101
	HttpStatusCodeProcessing                                   = 102
	HttpStatusCodeEarlyHints                                   = 103
	HttpStatusCodeOk                                           = 200
	HttpStatusCodeCreated                                      = 201
	HttpStatusCodeAccepted                                     = 202
	HttpStatusCodeNonAuthoritativeInformation                  = 203
	HttpStatusCodeNoContent                                    = 204
	HttpStatusCodeResetContent                                 = 205
	HttpStatusCodePartialContent                               = 206
	HttpStatusCodeMultiStatus                                  = 207
	HttpStatusCodeAlreadyReported                              = 208
	HttpStatusCodeImUsed                                       = 226
	HttpStatusCodeMultipleChoices                              = 300
	HttpStatusCodeMovedPermanently                             = 301
	HttpStatusCodeFound                                        = 302
	HttpStatusCodeSeeOther                                     = 303
	HttpStatusCodeNotModified                                  = 304
	HttpStatusCodeUseProxy                                     = 305
	HttpStatusCodeUnused                                       = 306
	HttpStatusCodeTemporaryRedirect                            = 307
	HttpStatusCodePermanentRedirect                            = 308
	HttpStatusCodeBadRequest                                   = 400
	HttpStatusCodeUnauthorized                                 = 401
	HttpStatusCodePaymentRequired                              = 402
	HttpStatusCodeForbidden                                    = 403
	HttpStatusCodeNotFound                                     = 404
	HttpStatusCodeMethodNotAllowed                             = 405
	HttpStatusCodeNotAcceptable                                = 406
	HttpStatusCodeProxyAuthenticationRequired                  = 407
	HttpStatusCodeRequestTimeout                               = 408
	HttpStatusCodeConflict                                     = 409
	HttpStatusCodeGone                                         = 410
	HttpStatusCodeLengthRequired                               = 411
	HttpStatusCodePreconditionFailed                           = 412
	HttpStatusCodePayloadTooLarge                              = 413
	HttpStatusCodeRequestUriTooLong                            = 414
	HttpStatusCodeUnsupportedMediaType                         = 415
	HttpStatusCodeRequestedRangeNotSatisfiable                 = 416
	HttpStatusCodeExpectationFailed                            = 417
	HttpStatusCodeImATeapot                                    = 418
	HttpStatusCodeInsufficientSpaceOnResource                  = 419
	HttpStatusCodeMethodFailure                                = 420
	HttpStatusCodeMisdirectedRequest                           = 421
	HttpStatusCodeUnprocessableEntity                          = 422
	HttpStatusCodeLocked                                       = 423
	HttpStatusCodeFailedDependency                             = 424
	HttpStatusCodeUpgradeRequired                              = 426
	HttpStatusCodePreconditionRequired                         = 428
	HttpStatusCodeTooManyRequests                              = 429
	HttpStatusCodeRequestHeaderFieldsTooLarge                  = 431
	HttpStatusCodeUnavailableForLegalReasons                   = 451
	HttpStatusCodeInternalServerError                          = 500
	HttpStatusCodeNotImplemented                               = 501
	HttpStatusCodeBadGateway                                   = 502
	HttpStatusCodeServiceUnavailable                           = 503
	HttpStatusCodeGatewayTimeout                               = 504
	HttpStatusCodeHttpVersionNotSupported                      = 505
	HttpStatusCodeVariantAlsoNegotiates                        = 506
	HttpStatusCodeInsufficientStorage                          = 507
	HttpStatusCodeLoopDetected                                 = 508
	HttpStatusCodeNotExtended                                  = 510
	HttpStatusCodeNetworkAuthenticationRequired                = 511
)

const (
	HttpStatusCodeValuesContinue                      HttpStatusCodeValue = "CONTINUE"
	HttpStatusCodeValuesSwitchingProtocols                                = "SWITCHING_PROTOCOLS"
	HttpStatusCodeValuesProcessing                                        = "PROCESSING"
	HttpStatusCodeValuesEarlyHints                                        = "EARLY_HINTS"
	HttpStatusCodeValuesOk                                                = "OK"
	HttpStatusCodeValuesCreated                                           = "CREATED"
	HttpStatusCodeValuesAccepted                                          = "ACCEPTED"
	HttpStatusCodeValuesNonAuthoritativeInformation                       = "NON_AUTHORITATIVE_INFORMATION"
	HttpStatusCodeValuesNoContent                                         = "NO_CONTENT"
	HttpStatusCodeValuesResetContent                                      = "RESET_CONTENT"
	HttpStatusCodeValuesPartialContent                                    = "PARTIAL_CONTENT"
	HttpStatusCodeValuesMultiStatus                                       = "MULTI_STATUS"
	HttpStatusCodeValuesAlreadyReported                                   = "ALREADY_REPORTED"
	HttpStatusCodeValuesImUsed                                            = "IM_USED"
	HttpStatusCodeValuesMultipleChoices                                   = "MULTI_CHOICES"
	HttpStatusCodeValuesMovedPermanently                                  = "MOVED_PERMANENTLY"
	HttpStatusCodeValuesFound                                             = "FOUND"
	HttpStatusCodeValuesSeeOther                                          = "SEE_OTHER"
	HttpStatusCodeValuesNotModified                                       = "NOT_MODIFIED"
	HttpStatusCodeValuesUseProxy                                          = "USE_PROXY"
	HttpStatusCodeValuesUnused                                            = "UNUSED"
	HttpStatusCodeValuesTemporaryRedirect                                 = "TEMPORARY_REDIRECT"
	HttpStatusCodeValuesPermanentRedirect                                 = "PERMANENT_REDIRECT"
	HttpStatusCodeValuesBadRequest                                        = "BAD_REQUEST"
	HttpStatusCodeValuesUnauthorized                                      = "UNAUTHORIZED"
	HttpStatusCodeValuesPaymentRequired                                   = "PAYMENT_REQUIRED"
	HttpStatusCodeValuesForbidden                                         = "FORBIDDEN"
	HttpStatusCodeValuesNotFound                                          = "NOT_FOUND"
	HttpStatusCodeValuesMethodNotAllowed                                  = "METHOD_NOT_ALLOWED"
	HttpStatusCodeValuesNotAcceptable                                     = "NOT_ACCEPTABLE"
	HttpStatusCodeValuesProxyAuthenticationRequired                       = "PROXY_AUTHENTICATION_REQUIRED"
	HttpStatusCodeValuesRequestTimeout                                    = "REQUEST_TIMEOUT"
	HttpStatusCodeValuesConflict                                          = "CONFLICT"
	HttpStatusCodeValuesGone                                              = "GONE"
	HttpStatusCodeValuesLengthRequired                                    = "LENGTH_REQUIRED"
	HttpStatusCodeValuesPreconditionFailed                                = "PRECONDITION_FAILED"
	HttpStatusCodeValuesPayloadTooLarge                                   = "PAYLOAD_TOO_LARGE"
	HttpStatusCodeValuesRequestUriTooLong                                 = "REQUEST_URI_TOO_LONG"
	HttpStatusCodeValuesUnsupportedMediaType                              = "UNSUPPORTED_MEDIA_TYPE"
	HttpStatusCodeValuesRequestedRangeNotSatisfiable                      = "REQUESTED_RANGE_NOT_SATISFIABLE"
	HttpStatusCodeValuesExpectationFailed                                 = "EXPECTATION_FAILED"
	HttpStatusCodeValuesImATeapot                                         = "IM_A_TEAPOT"
	HttpStatusCodeValuesInsufficientSpaceOnResource                       = "INSUFFICIENT_SPACE_ON_RESOURCE"
	HttpStatusCodeValuesMethodFailure                                     = "METHOD_FAILURE"
	HttpStatusCodeValuesMisdirectedRequest                                = "MISDIRECTED_REQUEST"
	HttpStatusCodeValuesUnprocessableEntity                               = "UNPROCESSABLE_ENTITY"
	HttpStatusCodeValuesLocked                                            = "LOCKED"
	HttpStatusCodeValuesFailedDependency                                  = "FAILED_DEPENDENCY"
	HttpStatusCodeValuesUpgradeRequired                                   = "UPGRADE_REQUIRED"
	HttpStatusCodeValuesPreconditionRequired                              = "PRECONDITION_REQUIRED"
	HttpStatusCodeValuesTooManyRequests                                   = "TOO_MANY_REQUESTS"
	HttpStatusCodeValuesRequestHeaderFieldsTooLarge                       = "REQUEST_HEADER_FIELDS_TOO_LARGE"
	HttpStatusCodeValuesUnavailableForLegalReasons                        = "UNAVAILABLE_FOR_LEGAL_REASONS"
	HttpStatusCodeValuesInternalServerError                               = "INTERNAL_SERVER_ERROR"
	HttpStatusCodeValuesNotImplemented                                    = "NOT_IMPLEMENTED"
	HttpStatusCodeValuesBadGateway                                        = "BAD_GATEWAY"
	HttpStatusCodeValuesServiceUnavailable                                = "SERVICE_UNAVAILABLE"
	HttpStatusCodeValuesGatewayTimeout                                    = "GATEWAY_TIMEOUT"
	HttpStatusCodeValuesHttpVersionNotSupported                           = "HTTP_VERSION_NOT_SUPPORTED"
	HttpStatusCodeValuesVariantAlsoNegotiates                             = "VARIANT_ALSO_NEGOTIATES"
	HttpStatusCodeValuesInsufficientStorage                               = "INSUFFICIENT_STORAGE"
	HttpStatusCodeValuesLoopDetected                                      = "LOOP_DETECTED"
	HttpStatusCodeValuesNotExtended                                       = "NOT_EXTENDED"
	HttpStatusCodeValuesNetworkAuthenticationRequired                     = "NETWORK_AUTHENTICATION_REQUIRED"
)

func (httpStatusCodeValues HttpStatusCodeValue) FromCode(code HttpStatusCode) HttpStatusCodeValue {
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

func (httpStatusCodeValues HttpStatusCodeValue) ParseCode(value string) HttpStatusCode {
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
