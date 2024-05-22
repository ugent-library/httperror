// Package httperror is a convenience package that defines error objects
// for each 4xx and 5xx HTTP status.
package httperror

import (
	"fmt"
	"net/http"
)

var (
	// 4xx
	BadRequest                   = &Error{StatusCode: http.StatusBadRequest}
	Unauthorized                 = &Error{StatusCode: http.StatusUnauthorized}
	PaymentRequired              = &Error{StatusCode: http.StatusPaymentRequired}
	Forbidden                    = &Error{StatusCode: http.StatusForbidden}
	NotFound                     = &Error{StatusCode: http.StatusNotFound}
	MethodNotAllowed             = &Error{StatusCode: http.StatusMethodNotAllowed}
	NotAcceptable                = &Error{StatusCode: http.StatusNotAcceptable}
	ProxyAuthRequired            = &Error{StatusCode: http.StatusProxyAuthRequired}
	RequestTimeout               = &Error{StatusCode: http.StatusRequestTimeout}
	Conflict                     = &Error{StatusCode: http.StatusConflict}
	Gone                         = &Error{StatusCode: http.StatusGone}
	LengthRequired               = &Error{StatusCode: http.StatusLengthRequired}
	PreconditionFailed           = &Error{StatusCode: http.StatusPreconditionFailed}
	RequestEntityTooLarge        = &Error{StatusCode: http.StatusRequestEntityTooLarge}
	RequestURITooLong            = &Error{StatusCode: http.StatusRequestURITooLong}
	UnsupportedMediaType         = &Error{StatusCode: http.StatusUnsupportedMediaType}
	RequestedRangeNotSatisfiable = &Error{StatusCode: http.StatusRequestedRangeNotSatisfiable}
	ExpectationFailed            = &Error{StatusCode: http.StatusExpectationFailed}
	Teapot                       = &Error{StatusCode: http.StatusTeapot}
	MisdirectedRequest           = &Error{StatusCode: http.StatusMisdirectedRequest}
	UnprocessableEntity          = &Error{StatusCode: http.StatusUnprocessableEntity}
	Locked                       = &Error{StatusCode: http.StatusLocked}
	FailedDependency             = &Error{StatusCode: http.StatusFailedDependency}
	TooEarly                     = &Error{StatusCode: http.StatusTooEarly}
	UpgradeRequired              = &Error{StatusCode: http.StatusUpgradeRequired}
	PreconditionRequired         = &Error{StatusCode: http.StatusPreconditionRequired}
	TooManyRequests              = &Error{StatusCode: http.StatusTooManyRequests}
	RequestHeaderFieldsTooLarge  = &Error{StatusCode: http.StatusRequestHeaderFieldsTooLarge}
	UnavailableForLegalReasons   = &Error{StatusCode: http.StatusUnavailableForLegalReasons}
	// 5xx
	InternalServerError           = &Error{StatusCode: http.StatusInternalServerError}
	NotImplemented                = &Error{StatusCode: http.StatusNotImplemented}
	BadGateway                    = &Error{StatusCode: http.StatusBadGateway}
	ServiceUnavailable            = &Error{StatusCode: http.StatusServiceUnavailable}
	GatewayTimeout                = &Error{StatusCode: http.StatusGatewayTimeout}
	HTTPVersionNotSupported       = &Error{StatusCode: http.StatusHTTPVersionNotSupported}
	VariantAlsoNegotiates         = &Error{StatusCode: http.StatusVariantAlsoNegotiates}
	InsufficientStorage           = &Error{StatusCode: http.StatusInsufficientStorage}
	LoopDetected                  = &Error{StatusCode: http.StatusLoopDetected}
	NotExtended                   = &Error{StatusCode: http.StatusNotExtended}
	NetworkAuthenticationRequired = &Error{StatusCode: http.StatusNetworkAuthenticationRequired}
)

type Error struct {
	StatusCode int
}

func (e *Error) Error() string {
	return fmt.Sprintf("http error %d (%s)", e.StatusCode, http.StatusText(e.StatusCode))
}

func (e *Error) Wrap(err error) error {
	return fmt.Errorf("%w: %w", e, err)
}
