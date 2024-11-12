package extras

import (
	"github.com/labstack/echo/v4"
	"io/fs"
	"nokowebapi/apis/schemas"
	"nokowebapi/nokocore"
	"strings"
)

type EchoGroupImpl interface {
	Use(middleware ...echo.MiddlewareFunc)
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Any(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route
	Match(methods []string, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route
	Group(prefix string, middleware ...echo.MiddlewareFunc) (sg *echo.Group)
	File(path string, file string)
	RouteNotFound(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Add(method string, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route
	Static(pathPrefix string, fsRoot string)
	StaticFS(pathPrefix string, filesystem fs.FS)
	FileFS(path string, file string, filesystem fs.FS, m ...echo.MiddlewareFunc) *echo.Route
}

func GetJwtTokenFromEchoContext(c echo.Context) (string, error) {
	var ok bool
	var token string
	nokocore.KeepVoid(ok, token)

	authorization := c.Request().Header.Get("Authorization")
	if token = strings.Trim(authorization, " "); token == "" {
		return "", nokocore.ErrJwtTokenInvalid
	}

	if token, ok = strings.CutPrefix(token, "Bearer "); !ok {
		return "", nokocore.ErrJwtTokenInvalid
	}
	if token = strings.Trim(token, " "); token == "" {
		return "", nokocore.ErrJwtTokenInvalid
	}
	return token, nil
}

func NewMessageBodyContinue(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeContinue
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodySwitchingProtocols(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeSwitchingProtocols
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyProcessing(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeProcessing
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyEarlyHints(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeEarlyHints
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyOk(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeOk
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyCreated(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeCreated
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyAccepted(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeAccepted
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyNonAuthoritativeInformation(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeNonAuthoritativeInformation
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyNoContent(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeNoContent
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyResetContent(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeResetContent
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyPartialContent(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodePartialContent
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyMultiStatus(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeMultiStatus
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyAlreadyReported(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeAlreadyReported
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyImUsed(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeImUsed
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyMultipleChoices(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeMultipleChoices
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyMovedPermanently(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeMovedPermanently
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyFound(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeFound
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodySeeOther(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeSeeOther
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyNotModified(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeNotModified
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyUseProxy(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeUseProxy
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyUnused(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeUnused
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyTemporaryRedirect(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeTemporaryRedirect
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyPermanentRedirect(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodePermanentRedirect
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyBadRequest(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeBadRequest
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyUnauthorized(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeUnauthorized
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyPaymentRequired(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodePaymentRequired
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyForbidden(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeForbidden
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyNotFound(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeNotFound
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyMethodNotAllowed(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeMethodNotAllowed
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyNotAcceptable(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeNotAcceptable
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyProxyAuthenticationRequired(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeProxyAuthenticationRequired
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyRequestTimeout(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeRequestTimeout
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyConflict(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeConflict
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyGone(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeGone
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyLengthRequired(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeLengthRequired
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyPreconditionFailed(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodePreconditionFailed
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyPayloadTooLarge(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodePayloadTooLarge
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyRequestUriTooLong(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeRequestUriTooLong
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyUnsupportedMediaType(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeUnsupportedMediaType
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyRequestedRangeNotSatisfiable(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeRequestedRangeNotSatisfiable
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyExpectationFailed(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeExpectationFailed
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyImATeapot(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeImATeapot
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyInsufficientSpaceOnResource(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeInsufficientSpaceOnResource
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyMethodFailure(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeMethodFailure
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyMisdirectedRequest(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeMisdirectedRequest
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyUnprocessableEntity(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeUnprocessableEntity
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyLocked(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeLocked
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyFailedDependency(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeFailedDependency
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyUpgradeRequired(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeUpgradeRequired
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyPreconditionRequired(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodePreconditionRequired
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyTooManyRequests(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeTooManyRequests
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyRequestHeaderFieldsTooLarge(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeRequestHeaderFieldsTooLarge
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyUnavailableForLegalReasons(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeUnavailableForLegalReasons
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyInternalServerError(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeInternalServerError
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyNotImplemented(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeNotImplemented
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyBadGateway(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeBadGateway
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyServiceUnavailable(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeServiceUnavailable
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyGatewayTimeout(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeGatewayTimeout
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyHttpVersionNotSupported(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeHttpVersionNotSupported
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyVariantAlsoNegotiates(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeVariantAlsoNegotiates
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyInsufficientStorage(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeInsufficientStorage
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyLoopDetected(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeLoopDetected
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyNotExtended(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeNotExtended
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}

func NewMessageBodyNetworkAuthenticationRequired(ctx echo.Context, message string, data any) error {
	httpStatusCode := nokocore.HttpStatusCodeNetworkAuthenticationRequired
	httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(httpStatusCode)
	return ctx.JSON(int(httpStatusCode), schemas.NewMessageBody(false, int(httpStatusCode), string(httpStatusCodeValue), message, data))
}
