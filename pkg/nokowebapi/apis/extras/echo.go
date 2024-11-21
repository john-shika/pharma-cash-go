package extras

import (
	"errors"
	"github.com/labstack/echo/v4"
	"io/fs"
	"net/http"
	"nokowebapi/apis/schemas"
	"nokowebapi/apis/validators"
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

func EchoHTTPErrorHandler() echo.HTTPErrorHandler {
	return func(err error, ctx echo.Context) {
		req := ctx.Request()

		if req.Method == echo.HEAD {
			err = ctx.NoContent(http.StatusInternalServerError)
			return
		}

		var httpError *echo.HTTPError
		if errors.As(err, &httpError) {
			message := nokocore.ToStringReflect(httpError.Message)
			httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(nokocore.HttpStatusCode(httpError.Code))
			messageBody := schemas.NewMessageBody(false, httpError.Code, string(httpStatusCodeValue), message, nil)
			err = ctx.JSON(httpError.Code, messageBody)
			return
		}

		var validationError *validators.ValidateError
		if errors.As(err, &validationError) {
			message := "Validation failed."
			fields := validationError.Fields()
			err = NewMessageBodyUnprocessableEntity(ctx, message, nokocore.MapAny{
				"fields": fields,
			})
			return
		}

		err = NewMessageBodyInternalServerError(ctx, err.Error(), nil)
	}
}

func NewMessageBodyContinue(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeContinue), schemas.NewMessageBodyContinue(message, data))
}

func NewMessageBodySwitchingProtocols(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeSwitchingProtocols), schemas.NewMessageBodySwitchingProtocols(message, data))
}

func NewMessageBodyProcessing(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeProcessing), schemas.NewMessageBodyProcessing(message, data))
}

func NewMessageBodyEarlyHints(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeEarlyHints), schemas.NewMessageBodyEarlyHints(message, data))
}

func NewMessageBodyOk(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeOk), schemas.NewMessageBodyOk(message, data))
}

func NewMessageBodyCreated(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeCreated), schemas.NewMessageBodyCreated(message, data))
}

func NewMessageBodyAccepted(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeAccepted), schemas.NewMessageBodyAccepted(message, data))
}

func NewMessageBodyNonAuthoritativeInformation(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeNonAuthoritativeInformation), schemas.NewMessageBodyNonAuthoritativeInformation(message, data))
}

func NewMessageBodyNoContent(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeNoContent), schemas.NewMessageBodyNoContent(message, data))
}

func NewMessageBodyResetContent(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeResetContent), schemas.NewMessageBodyResetContent(message, data))
}

func NewMessageBodyPartialContent(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodePartialContent), schemas.NewMessageBodyPartialContent(message, data))
}

func NewMessageBodyMultiStatus(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeMultiStatus), schemas.NewMessageBodyMultiStatus(message, data))
}

func NewMessageBodyAlreadyReported(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeAlreadyReported), schemas.NewMessageBodyAlreadyReported(message, data))
}

func NewMessageBodyImUsed(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeImUsed), schemas.NewMessageBodyImUsed(message, data))
}

func NewMessageBodyMultipleChoices(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeMultipleChoices), schemas.NewMessageBodyMultipleChoices(message, data))
}

func NewMessageBodyMovedPermanently(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeMovedPermanently), schemas.NewMessageBodyMovedPermanently(message, data))
}

func NewMessageBodyFound(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeFound), schemas.NewMessageBodyFound(message, data))
}

func NewMessageBodySeeOther(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeSeeOther), schemas.NewMessageBodySeeOther(message, data))
}

func NewMessageBodyNotModified(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeNotModified), schemas.NewMessageBodyNotModified(message, data))
}

func NewMessageBodyUseProxy(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeUseProxy), schemas.NewMessageBodyUseProxy(message, data))
}

func NewMessageBodyUnused(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeUnused), schemas.NewMessageBodyUnused(message, data))
}

func NewMessageBodyTemporaryRedirect(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeTemporaryRedirect), schemas.NewMessageBodyTemporaryRedirect(message, data))
}

func NewMessageBodyPermanentRedirect(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodePermanentRedirect), schemas.NewMessageBodyPermanentRedirect(message, data))
}

func NewMessageBodyBadRequest(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeBadRequest), schemas.NewMessageBodyBadRequest(message, data))
}

func NewMessageBodyUnauthorized(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeUnauthorized), schemas.NewMessageBodyUnauthorized(message, data))
}

func NewMessageBodyPaymentRequired(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodePaymentRequired), schemas.NewMessageBodyPaymentRequired(message, data))
}

func NewMessageBodyForbidden(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeForbidden), schemas.NewMessageBodyForbidden(message, data))
}

func NewMessageBodyNotFound(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeNotFound), schemas.NewMessageBodyNotFound(message, data))
}

func NewMessageBodyMethodNotAllowed(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeMethodNotAllowed), schemas.NewMessageBodyMethodNotAllowed(message, data))
}

func NewMessageBodyNotAcceptable(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeNotAcceptable), schemas.NewMessageBodyNotAcceptable(message, data))
}

func NewMessageBodyProxyAuthenticationRequired(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeProxyAuthenticationRequired), schemas.NewMessageBodyProxyAuthenticationRequired(message, data))
}

func NewMessageBodyRequestTimeout(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeRequestTimeout), schemas.NewMessageBodyRequestTimeout(message, data))
}

func NewMessageBodyConflict(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeConflict), schemas.NewMessageBodyConflict(message, data))
}

func NewMessageBodyGone(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeGone), schemas.NewMessageBodyGone(message, data))
}

func NewMessageBodyLengthRequired(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeLengthRequired), schemas.NewMessageBodyLengthRequired(message, data))
}

func NewMessageBodyPreconditionFailed(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodePreconditionFailed), schemas.NewMessageBodyPreconditionFailed(message, data))
}

func NewMessageBodyPayloadTooLarge(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodePayloadTooLarge), schemas.NewMessageBodyPayloadTooLarge(message, data))
}

func NewMessageBodyRequestUriTooLong(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeRequestUriTooLong), schemas.NewMessageBodyRequestUriTooLong(message, data))
}

func NewMessageBodyUnsupportedMediaType(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeUnsupportedMediaType), schemas.NewMessageBodyUnsupportedMediaType(message, data))
}

func NewMessageBodyRequestedRangeNotSatisfiable(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeRequestedRangeNotSatisfiable), schemas.NewMessageBodyRequestedRangeNotSatisfiable(message, data))
}

func NewMessageBodyExpectationFailed(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeExpectationFailed), schemas.NewMessageBodyExpectationFailed(message, data))
}

func NewMessageBodyImATeapot(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeImATeapot), schemas.NewMessageBodyImATeapot(message, data))
}

func NewMessageBodyInsufficientSpaceOnResource(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeInsufficientSpaceOnResource), schemas.NewMessageBodyInsufficientSpaceOnResource(message, data))
}

func NewMessageBodyMethodFailure(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeMethodFailure), schemas.NewMessageBodyMethodFailure(message, data))
}

func NewMessageBodyMisdirectedRequest(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeMisdirectedRequest), schemas.NewMessageBodyMisdirectedRequest(message, data))
}

func NewMessageBodyUnprocessableEntity(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeUnprocessableEntity), schemas.NewMessageBodyUnprocessableEntity(message, data))
}

func NewMessageBodyLocked(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeLocked), schemas.NewMessageBodyLocked(message, data))
}

func NewMessageBodyFailedDependency(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeFailedDependency), schemas.NewMessageBodyFailedDependency(message, data))
}

func NewMessageBodyUpgradeRequired(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeUpgradeRequired), schemas.NewMessageBodyUpgradeRequired(message, data))
}

func NewMessageBodyPreconditionRequired(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodePreconditionRequired), schemas.NewMessageBodyPreconditionRequired(message, data))
}

func NewMessageBodyTooManyRequests(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeTooManyRequests), schemas.NewMessageBodyTooManyRequests(message, data))
}

func NewMessageBodyRequestHeaderFieldsTooLarge(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeRequestHeaderFieldsTooLarge), schemas.NewMessageBodyRequestHeaderFieldsTooLarge(message, data))
}

func NewMessageBodyUnavailableForLegalReasons(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeUnavailableForLegalReasons), schemas.NewMessageBodyUnavailableForLegalReasons(message, data))
}

func NewMessageBodyInternalServerError(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeInternalServerError), schemas.NewMessageBodyInternalServerError(message, data))
}

func NewMessageBodyNotImplemented(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeNotImplemented), schemas.NewMessageBodyNotImplemented(message, data))
}

func NewMessageBodyBadGateway(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeBadGateway), schemas.NewMessageBodyBadGateway(message, data))
}

func NewMessageBodyServiceUnavailable(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeServiceUnavailable), schemas.NewMessageBodyServiceUnavailable(message, data))
}

func NewMessageBodyGatewayTimeout(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeGatewayTimeout), schemas.NewMessageBodyGatewayTimeout(message, data))
}

func NewMessageBodyHttpVersionNotSupported(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeHttpVersionNotSupported), schemas.NewMessageBodyHttpVersionNotSupported(message, data))
}

func NewMessageBodyVariantAlsoNegotiates(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeVariantAlsoNegotiates), schemas.NewMessageBodyVariantAlsoNegotiates(message, data))
}

func NewMessageBodyInsufficientStorage(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeInsufficientStorage), schemas.NewMessageBodyInsufficientStorage(message, data))
}

func NewMessageBodyLoopDetected(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeLoopDetected), schemas.NewMessageBodyLoopDetected(message, data))
}

func NewMessageBodyNotExtended(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeNotExtended), schemas.NewMessageBodyNotExtended(message, data))
}

func NewMessageBodyNetworkAuthenticationRequired(ctx echo.Context, message string, data any) error {
	return ctx.JSON(int(nokocore.HttpStatusCodeNetworkAuthenticationRequired), schemas.NewMessageBodyNetworkAuthenticationRequired(message, data))
}
