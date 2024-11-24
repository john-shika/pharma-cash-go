package extras

import (
	"errors"
	"github.com/labstack/echo/v4"
	"io/fs"
	"net/http"
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

func GetJwtTokenFromEchoContext(ctx echo.Context) (string, error) {
	var ok bool
	var token string
	nokocore.KeepVoid(ok, token)

	req := ctx.Request()
	authorization := req.Header.Get("Authorization")
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
			httpStatusCode := nokocore.HttpStatusCode(httpError.Code)
			httpStatusCodeValue := httpStatusCode.ToString()
			messageBody := schemas.NewMessageBody(false, httpError.Code, string(httpStatusCodeValue), message, nil)
			err = ctx.JSON(httpError.Code, messageBody)
			return
		}

		var validationError *nokocore.ValidateError
		if errors.As(err, &validationError) {
			message := "Validation failed."
			fields := validationError.Fields()
			err = NewMessageBodyUnprocessableEntity(ctx, message, &nokocore.MapAny{
				"fields": fields,
			})
			return
		}

		err = NewMessageBodyInternalServerError(ctx, err.Error(), nil)
	}
}

func NewMessageBodyContinue(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeContinue.ToInt(), schemas.NewMessageBodyContinue(message, data))
}

func NewMessageBodySwitchingProtocols(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeSwitchingProtocols.ToInt(), schemas.NewMessageBodySwitchingProtocols(message, data))
}

func NewMessageBodyProcessing(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeProcessing.ToInt(), schemas.NewMessageBodyProcessing(message, data))
}

func NewMessageBodyEarlyHints(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeEarlyHints.ToInt(), schemas.NewMessageBodyEarlyHints(message, data))
}

func NewMessageBodyOk(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeOk.ToInt(), schemas.NewMessageBodyOk(message, data))
}

func NewMessageBodyCreated(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeCreated.ToInt(), schemas.NewMessageBodyCreated(message, data))
}

func NewMessageBodyAccepted(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeAccepted.ToInt(), schemas.NewMessageBodyAccepted(message, data))
}

func NewMessageBodyNonAuthoritativeInformation(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeNonAuthoritativeInformation.ToInt(), schemas.NewMessageBodyNonAuthoritativeInformation(message, data))
}

func NewMessageBodyNoContent(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeNoContent.ToInt(), schemas.NewMessageBodyNoContent(message, data))
}

func NewMessageBodyResetContent(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeResetContent.ToInt(), schemas.NewMessageBodyResetContent(message, data))
}

func NewMessageBodyPartialContent(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodePartialContent.ToInt(), schemas.NewMessageBodyPartialContent(message, data))
}

func NewMessageBodyMultiStatus(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeMultiStatus.ToInt(), schemas.NewMessageBodyMultiStatus(message, data))
}

func NewMessageBodyAlreadyReported(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeAlreadyReported.ToInt(), schemas.NewMessageBodyAlreadyReported(message, data))
}

func NewMessageBodyImUsed(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeImUsed.ToInt(), schemas.NewMessageBodyImUsed(message, data))
}

func NewMessageBodyMultipleChoices(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeMultipleChoices.ToInt(), schemas.NewMessageBodyMultipleChoices(message, data))
}

func NewMessageBodyMovedPermanently(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeMovedPermanently.ToInt(), schemas.NewMessageBodyMovedPermanently(message, data))
}

func NewMessageBodyFound(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeFound.ToInt(), schemas.NewMessageBodyFound(message, data))
}

func NewMessageBodySeeOther(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeSeeOther.ToInt(), schemas.NewMessageBodySeeOther(message, data))
}

func NewMessageBodyNotModified(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeNotModified.ToInt(), schemas.NewMessageBodyNotModified(message, data))
}

func NewMessageBodyUseProxy(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeUseProxy.ToInt(), schemas.NewMessageBodyUseProxy(message, data))
}

func NewMessageBodyUnused(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeUnused.ToInt(), schemas.NewMessageBodyUnused(message, data))
}

func NewMessageBodyTemporaryRedirect(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeTemporaryRedirect.ToInt(), schemas.NewMessageBodyTemporaryRedirect(message, data))
}

func NewMessageBodyPermanentRedirect(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodePermanentRedirect.ToInt(), schemas.NewMessageBodyPermanentRedirect(message, data))
}

func NewMessageBodyBadRequest(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeBadRequest.ToInt(), schemas.NewMessageBodyBadRequest(message, data))
}

func NewMessageBodyUnauthorized(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeUnauthorized.ToInt(), schemas.NewMessageBodyUnauthorized(message, data))
}

func NewMessageBodyPaymentRequired(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodePaymentRequired.ToInt(), schemas.NewMessageBodyPaymentRequired(message, data))
}

func NewMessageBodyForbidden(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeForbidden.ToInt(), schemas.NewMessageBodyForbidden(message, data))
}

func NewMessageBodyNotFound(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeNotFound.ToInt(), schemas.NewMessageBodyNotFound(message, data))
}

func NewMessageBodyMethodNotAllowed(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeMethodNotAllowed.ToInt(), schemas.NewMessageBodyMethodNotAllowed(message, data))
}

func NewMessageBodyNotAcceptable(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeNotAcceptable.ToInt(), schemas.NewMessageBodyNotAcceptable(message, data))
}

func NewMessageBodyProxyAuthenticationRequired(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeProxyAuthenticationRequired.ToInt(), schemas.NewMessageBodyProxyAuthenticationRequired(message, data))
}

func NewMessageBodyRequestTimeout(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeRequestTimeout.ToInt(), schemas.NewMessageBodyRequestTimeout(message, data))
}

func NewMessageBodyConflict(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeConflict.ToInt(), schemas.NewMessageBodyConflict(message, data))
}

func NewMessageBodyGone(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeGone.ToInt(), schemas.NewMessageBodyGone(message, data))
}

func NewMessageBodyLengthRequired(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeLengthRequired.ToInt(), schemas.NewMessageBodyLengthRequired(message, data))
}

func NewMessageBodyPreconditionFailed(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodePreconditionFailed.ToInt(), schemas.NewMessageBodyPreconditionFailed(message, data))
}

func NewMessageBodyPayloadTooLarge(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodePayloadTooLarge.ToInt(), schemas.NewMessageBodyPayloadTooLarge(message, data))
}

func NewMessageBodyRequestUriTooLong(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeRequestUriTooLong.ToInt(), schemas.NewMessageBodyRequestUriTooLong(message, data))
}

func NewMessageBodyUnsupportedMediaType(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeUnsupportedMediaType.ToInt(), schemas.NewMessageBodyUnsupportedMediaType(message, data))
}

func NewMessageBodyRequestedRangeNotSatisfiable(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeRequestedRangeNotSatisfiable.ToInt(), schemas.NewMessageBodyRequestedRangeNotSatisfiable(message, data))
}

func NewMessageBodyExpectationFailed(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeExpectationFailed.ToInt(), schemas.NewMessageBodyExpectationFailed(message, data))
}

func NewMessageBodyImATeapot(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeImATeapot.ToInt(), schemas.NewMessageBodyImATeapot(message, data))
}

func NewMessageBodyInsufficientSpaceOnResource(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeInsufficientSpaceOnResource.ToInt(), schemas.NewMessageBodyInsufficientSpaceOnResource(message, data))
}

func NewMessageBodyMethodFailure(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeMethodFailure.ToInt(), schemas.NewMessageBodyMethodFailure(message, data))
}

func NewMessageBodyMisdirectedRequest(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeMisdirectedRequest.ToInt(), schemas.NewMessageBodyMisdirectedRequest(message, data))
}

func NewMessageBodyUnprocessableEntity(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeUnprocessableEntity.ToInt(), schemas.NewMessageBodyUnprocessableEntity(message, data))
}

func NewMessageBodyLocked(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeLocked.ToInt(), schemas.NewMessageBodyLocked(message, data))
}

func NewMessageBodyFailedDependency(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeFailedDependency.ToInt(), schemas.NewMessageBodyFailedDependency(message, data))
}

func NewMessageBodyUpgradeRequired(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeUpgradeRequired.ToInt(), schemas.NewMessageBodyUpgradeRequired(message, data))
}

func NewMessageBodyPreconditionRequired(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodePreconditionRequired.ToInt(), schemas.NewMessageBodyPreconditionRequired(message, data))
}

func NewMessageBodyTooManyRequests(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeTooManyRequests.ToInt(), schemas.NewMessageBodyTooManyRequests(message, data))
}

func NewMessageBodyRequestHeaderFieldsTooLarge(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeRequestHeaderFieldsTooLarge.ToInt(), schemas.NewMessageBodyRequestHeaderFieldsTooLarge(message, data))
}

func NewMessageBodyUnavailableForLegalReasons(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeUnavailableForLegalReasons.ToInt(), schemas.NewMessageBodyUnavailableForLegalReasons(message, data))
}

func NewMessageBodyInternalServerError(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeInternalServerError.ToInt(), schemas.NewMessageBodyInternalServerError(message, data))
}

func NewMessageBodyNotImplemented(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeNotImplemented.ToInt(), schemas.NewMessageBodyNotImplemented(message, data))
}

func NewMessageBodyBadGateway(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeBadGateway.ToInt(), schemas.NewMessageBodyBadGateway(message, data))
}

func NewMessageBodyServiceUnavailable(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeServiceUnavailable.ToInt(), schemas.NewMessageBodyServiceUnavailable(message, data))
}

func NewMessageBodyGatewayTimeout(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeGatewayTimeout.ToInt(), schemas.NewMessageBodyGatewayTimeout(message, data))
}

func NewMessageBodyHttpVersionNotSupported(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeHttpVersionNotSupported.ToInt(), schemas.NewMessageBodyHttpVersionNotSupported(message, data))
}

func NewMessageBodyVariantAlsoNegotiates(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeVariantAlsoNegotiates.ToInt(), schemas.NewMessageBodyVariantAlsoNegotiates(message, data))
}

func NewMessageBodyInsufficientStorage(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeInsufficientStorage.ToInt(), schemas.NewMessageBodyInsufficientStorage(message, data))
}

func NewMessageBodyLoopDetected(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeLoopDetected.ToInt(), schemas.NewMessageBodyLoopDetected(message, data))
}

func NewMessageBodyNotExtended(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeNotExtended.ToInt(), schemas.NewMessageBodyNotExtended(message, data))
}

func NewMessageBodyNetworkAuthenticationRequired(ctx echo.Context, message string, data any) error {
	return ctx.JSON(nokocore.HttpStatusCodeNetworkAuthenticationRequired.ToInt(), schemas.NewMessageBodyNetworkAuthenticationRequired(message, data))
}
