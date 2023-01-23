// Package httperror is a convenience package that defines error objects
// for each 4xx and 5xx HTTP status.
package httperror

import (
	"fmt"
	"net/http"
)

var (
	// 4xx
	BadRequest                   = &Error{http.StatusBadRequest}
	Unauthorized                 = &Error{http.StatusUnauthorized}
	PaymentRequired              = &Error{http.StatusPaymentRequired}
	Forbidden                    = &Error{http.StatusForbidden}
	NotFound                     = &Error{http.StatusNotFound}
	MethodNotAllowed             = &Error{http.StatusMethodNotAllowed}
	NotAcceptable                = &Error{http.StatusNotAcceptable}
	ProxyAuthRequired            = &Error{http.StatusProxyAuthRequired}
	RequestTimeout               = &Error{http.StatusRequestTimeout}
	Conflict                     = &Error{http.StatusConflict}
	Gone                         = &Error{http.StatusGone}
	LengthRequired               = &Error{http.StatusLengthRequired}
	PreconditionFailed           = &Error{http.StatusPreconditionFailed}
	RequestEntityTooLarge        = &Error{http.StatusRequestEntityTooLarge}
	RequestURITooLong            = &Error{http.StatusRequestURITooLong}
	UnsupportedMediaType         = &Error{http.StatusUnsupportedMediaType}
	RequestedRangeNotSatisfiable = &Error{http.StatusRequestedRangeNotSatisfiable}
	ExpectationFailed            = &Error{http.StatusExpectationFailed}
	Teapot                       = &Error{http.StatusTeapot}
	MisdirectedRequest           = &Error{http.StatusMisdirectedRequest}
	UnprocessableEntity          = &Error{http.StatusUnprocessableEntity}
	Locked                       = &Error{http.StatusLocked}
	FailedDependency             = &Error{http.StatusFailedDependency}
	TooEarly                     = &Error{http.StatusTooEarly}
	UpgradeRequired              = &Error{http.StatusUpgradeRequired}
	PreconditionRequired         = &Error{http.StatusPreconditionRequired}
	TooManyRequests              = &Error{http.StatusTooManyRequests}
	RequestHeaderFieldsTooLarge  = &Error{http.StatusRequestHeaderFieldsTooLarge}
	UnavailableForLegalReasons   = &Error{http.StatusUnavailableForLegalReasons}
	// 5xx
	InternalServerError           = &Error{http.StatusInternalServerError}
	NotImplemented                = &Error{http.StatusNotImplemented}
	BadGateway                    = &Error{http.StatusBadGateway}
	ServiceUnavailable            = &Error{http.StatusServiceUnavailable}
	GatewayTimeout                = &Error{http.StatusGatewayTimeout}
	HTTPVersionNotSupported       = &Error{http.StatusHTTPVersionNotSupported}
	VariantAlsoNegotiates         = &Error{http.StatusVariantAlsoNegotiates}
	InsufficientStorage           = &Error{http.StatusInsufficientStorage}
	LoopDetected                  = &Error{http.StatusLoopDetected}
	NotExtended                   = &Error{http.StatusNotExtended}
	NetworkAuthenticationRequired = &Error{http.StatusNetworkAuthenticationRequired}
)

type Error struct {
	StatusCode int
}

func (e *Error) Error() string {
	return fmt.Sprintf("http error: %d %s", e.StatusCode, http.StatusText(e.StatusCode))
}
