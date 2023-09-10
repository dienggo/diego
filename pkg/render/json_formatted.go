package render

import (
	"net/http"
)

type Data struct {
	Code    int    `json:"code,omitempty"`
	Name    string `json:"name"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

// JsonFormatted writes the response headers and calls JSON to render data.
// with template json format response
func JsonFormatted(w http.ResponseWriter, data Data) {
	code := data.Code
	data.Code = 0
	Json(w, code, data)
}

// mapJsonFormatted is method to build template json response
func mapJsonFormatted(w http.ResponseWriter, code int, message string, data any) {
	JsonFormatted(w, Data{
		Code:    code,
		Name:    http.StatusText(code),
		Message: message,
		Data:    data,
	})
}

// JsonFormattedPaginate is method to build template json response with pagination format
func JsonFormattedPaginate(w http.ResponseWriter, code int, message string, data any, count int, limit int, page int) {
	prev := page - 1
	next := page + 1

	meta := map[string]any{
		"prev":  prev,
		"next":  next,
		"count": count,
	}
	if prev < 0 {
		meta["prev"] = nil
	}
	if count < limit {
		meta["next"] = nil
	}
	JsonFormatted(w, Data{
		Code:    code,
		Name:    http.StatusText(code),
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// FormattedJson is HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
type FormattedJson struct {
	Writer http.ResponseWriter
}

func (fj FormattedJson) StatusContinue(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusContinue, message, data)
}
func (fj FormattedJson) StatusSwitchingProtocols(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusSwitchingProtocols, message, data)
}
func (fj FormattedJson) StatusProcessing(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusProcessing, message, data)
}
func (fj FormattedJson) StatusEarlyHints(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusEarlyHints, message, data)
}
func (fj FormattedJson) StatusOK(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusOK, message, data)
}
func (fj FormattedJson) StatusCreated(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusCreated, message, data)
}
func (fj FormattedJson) StatusAccepted(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusAccepted, message, data)
}
func (fj FormattedJson) StatusNonAuthoritativeInfo(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusNonAuthoritativeInfo, message, data)
}
func (fj FormattedJson) StatusNoContent(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusNoContent, message, data)
}
func (fj FormattedJson) StatusResetContent(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusResetContent, message, data)
}
func (fj FormattedJson) StatusPartialContent(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusPartialContent, message, data)
}
func (fj FormattedJson) StatusMultiStatus(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusMultiStatus, message, data)
}
func (fj FormattedJson) StatusAlreadyReported(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusAlreadyReported, message, data)
}
func (fj FormattedJson) StatusIMUsed(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusIMUsed, message, data)
}
func (fj FormattedJson) StatusMultipleChoices(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusMultipleChoices, message, data)
}
func (fj FormattedJson) StatusMovedPermanently(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusMovedPermanently, message, data)
}
func (fj FormattedJson) StatusFound(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusFound, message, data)
}
func (fj FormattedJson) StatusSeeOther(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusSeeOther, message, data)
}
func (fj FormattedJson) StatusNotModified(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusNotModified, message, data)
}
func (fj FormattedJson) StatusUseProxy(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusUseProxy, message, data)
}
func (fj FormattedJson) StatusTemporaryRedirect(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusTemporaryRedirect, message, data)
}
func (fj FormattedJson) StatusPermanentRedirect(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusPermanentRedirect, message, data)
}
func (fj FormattedJson) StatusBadRequest(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusBadRequest, message, data)
}
func (fj FormattedJson) StatusUnauthorized(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusUnauthorized, message, data)
}
func (fj FormattedJson) StatusPaymentRequired(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusPaymentRequired, message, data)
}
func (fj FormattedJson) StatusForbidden(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusForbidden, message, data)
}
func (fj FormattedJson) StatusNotFound(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusNotFound, message, data)
}
func (fj FormattedJson) StatusMethodNotAllowed(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusMethodNotAllowed, message, data)
}
func (fj FormattedJson) StatusNotAcceptable(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusNotAcceptable, message, data)
}
func (fj FormattedJson) StatusProxyAuthRequired(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusProxyAuthRequired, message, data)
}
func (fj FormattedJson) StatusRequestTimeout(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusRequestTimeout, message, data)
}
func (fj FormattedJson) StatusConflict(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusConflict, message, data)
}
func (fj FormattedJson) StatusGone(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusGone, message, data)
}
func (fj FormattedJson) StatusLengthRequired(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusLengthRequired, message, data)
}
func (fj FormattedJson) StatusPreconditionFailed(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusPreconditionFailed, message, data)
}
func (fj FormattedJson) StatusRequestEntityTooLarge(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusRequestEntityTooLarge, message, data)
}
func (fj FormattedJson) StatusRequestURITooLong(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusRequestURITooLong, message, data)
}
func (fj FormattedJson) StatusUnsupportedMediaType(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusUnsupportedMediaType, message, data)
}
func (fj FormattedJson) StatusRequestedRangeNotSatisfiable(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusRequestedRangeNotSatisfiable, message, data)
}
func (fj FormattedJson) StatusExpectationFailed(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusExpectationFailed, message, data)
}
func (fj FormattedJson) StatusTeapot(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusTeapot, message, data)
}
func (fj FormattedJson) StatusMisdirectedRequest(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusMisdirectedRequest, message, data)
}
func (fj FormattedJson) StatusUnprocessableEntity(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusUnprocessableEntity, message, data)
}
func (fj FormattedJson) StatusLocked(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusLocked, message, data)
}
func (fj FormattedJson) StatusFailedDependency(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusFailedDependency, message, data)
}
func (fj FormattedJson) StatusTooEarly(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusTooEarly, message, data)
}
func (fj FormattedJson) StatusUpgradeRequired(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusUpgradeRequired, message, data)
}
func (fj FormattedJson) StatusPreconditionRequired(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusPreconditionRequired, message, data)
}
func (fj FormattedJson) StatusTooManyRequests(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusTooManyRequests, message, data)
}
func (fj FormattedJson) StatusRequestHeaderFieldsTooLarge(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusRequestHeaderFieldsTooLarge, message, data)
}
func (fj FormattedJson) StatusUnavailableForLegalReasons(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusUnavailableForLegalReasons, message, data)
}
func (fj FormattedJson) StatusInternalServerError(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusInternalServerError, message, data)
}
func (fj FormattedJson) StatusNotImplemented(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusNotImplemented, message, data)
}
func (fj FormattedJson) StatusBadGateway(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusBadGateway, message, data)
}
func (fj FormattedJson) StatusServiceUnavailable(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusServiceUnavailable, message, data)
}
func (fj FormattedJson) StatusGatewayTimeout(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusGatewayTimeout, message, data)
}
func (fj FormattedJson) StatusHTTPVersionNotSupported(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusHTTPVersionNotSupported, message, data)
}
func (fj FormattedJson) StatusVariantAlsoNegotiates(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusVariantAlsoNegotiates, message, data)
}
func (fj FormattedJson) StatusInsufficientStorage(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusInsufficientStorage, message, data)
}
func (fj FormattedJson) StatusLoopDetected(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusLoopDetected, message, data)
}
func (fj FormattedJson) StatusNotExtended(message string, data any) {
	mapJsonFormatted(fj.Writer, http.StatusNotExtended, message, data)
}
