// Code generated by oa3 (https://github.com/aarondl/oa3). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.
package oa3gen

import (
	"net/http"
	"time"

	"github.com/aarondl/chrono"
	"github.com/aarondl/oa3/support"
	"github.com/aarondl/opt/null"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/go-chi/chi/v5"
)

// Interface is the interface that an application server must implement
// in order to use this package.
//
// A great api
type Interface interface {
	// Authenticate post /auth
	Authenticate(w http.ResponseWriter, r *http.Request) (AuthenticateResponse, error)
	// TestInlinePrimitiveBody get /test/inline
	TestInlinePrimitiveBody(w http.ResponseWriter, r *http.Request, body string) (TestInlinePrimitiveBodyResponse, error)
	// TestInline post /test/inline
	TestInline(w http.ResponseWriter, r *http.Request, body TestInlineInline) (TestInlineResponse, error)
	// GetUser get /users/{id}
	// Retrieves a user with a long description that spans multiple lines so
	// that we can see that both wrapping and long-line support is not
	// bleeding over the sacred 80 char limit.
	GetUser(w http.ResponseWriter, r *http.Request, id string, validStr omitnull.Val[string], reqValidStr null.Val[string], validInt omit.Val[int], reqValidInt int, validNum omit.Val[float64], reqValidNum float64, validBool omit.Val[bool], reqValidBool bool, reqStrFormat string, dateTime chrono.DateTime, date chrono.Date, timeVal chrono.Time, durationVal time.Duration) (GetUserResponse, error)
	// SetUser post /users/{id}
	// Sets a user
	SetUser(w http.ResponseWriter, r *http.Request, body *Primitives) (SetUserResponse, error)
}

type (
	// GoServer implements all the wrapper functionality for the API
	GoServer struct {
		impl      Interface
		converter support.ValidationConverter
		router    *chi.Mux
	}
)

// NewGoServer constructor
func NewGoServer(
	impl Interface,
	cnv support.ValidationConverter,
	eh support.ErrorHandler,
	mw support.MW,
) http.Handler {

	o := GoServer{
		impl:      impl,
		converter: cnv,
		router:    chi.NewRouter(),
	}

	// Untagged operations
	o.router.Group(func(r chi.Router) {
		if m, ok := mw[""]; ok {
			r.Use(m...)
		}
		r.Method(http.MethodPost, `/auth`, eh.Wrap(o.authenticateOp))
		r.Method(http.MethodPost, `/test/inline`, eh.Wrap(o.testinlineOp))
		r.Method(http.MethodGet, `/test/inline`, eh.Wrap(o.testinlineprimitivebodyOp))
	})
	// users tagged operations
	o.router.Group(func(r chi.Router) {
		if m, ok := mw["users"]; ok {
			r.Use(m...)
		}
		r.Method(http.MethodGet, `/users/{id}`, eh.Wrap(o.getuserOp))
		r.Method(http.MethodPost, `/users/{id}`, eh.Wrap(o.setuserOp))
	})

	return o
}

// ServeHTTP implements http.Handler
func (o GoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	o.router.ServeHTTP(w, r)
}

type validatable interface {
	validateSchema() support.Errors
}

// Validate a schema
func Validate[T validatable](toValidate T) support.Errors {
	return toValidate.validateSchema()
}

// AuthenticateResponse one-of enforcer
//
// Implementors:
// - HTTPStatusOk
type AuthenticateResponse interface {
	AuthenticateImpl()
}

// AuthenticateImpl implements AuthenticateResponse(200) for HTTPStatusOk
func (HTTPStatusOk) AuthenticateImpl() {}

// TestInlinePrimitiveBodyResponse one-of enforcer
//
// Implementors:
// - HTTPStatusOk
type TestInlinePrimitiveBodyResponse interface {
	TestInlinePrimitiveBodyImpl()
}

// TestInlinePrimitiveBodyImpl implements TestInlinePrimitiveBodyResponse(200) for HTTPStatusOk
func (HTTPStatusOk) TestInlinePrimitiveBodyImpl() {}

// TestInlineResponse one-of enforcer
//
// Implementors:
// - TestInline200Inline
// - TestInline201Inline
type TestInlineResponse interface {
	TestInlineImpl()
}

// TestInlineImpl implements TestInlineHeadersResponse(200) for
func (TestInline200Inline) TestInlineImpl() {}

// TestInlineImpl implements TestInlineHeadersResponse(201) for
func (TestInline201Inline) TestInlineImpl() {}

// GetUserResponse one-of enforcer
//
// Implementors:
// - HTTPStatusNotModified
type GetUserResponse interface {
	GetUserImpl()
}

// GetUserImpl implements GetUserResponse(304) for HTTPStatusNotModified
func (HTTPStatusNotModified) GetUserImpl() {}

// SetUserResponse one-of enforcer
//
// Implementors:
// - SetUser200HeadersResponse
// - Primitives - #/components/schemas/Primitives
type SetUserResponse interface {
	SetUserImpl()
}

// SetUser200WrappedResponse wraps the normal body response with a
// struct to be able to additionally return headers or differentiate between
// multiple response codes with the same response body.
type SetUser200WrappedResponse struct {
	HeaderXResponseHeader omit.Val[string]
	Body                  Primitives
}

// SetUserImpl implements SetUserResponse(200) for SetUser200WrappedResponse
func (SetUser200WrappedResponse) SetUserImpl() {}

// SetUserdefaultWrappedResponse wraps the normal body response with a
// struct to be able to additionally return headers or differentiate between
// multiple response codes with the same response body.
type SetUserdefaultWrappedResponse struct {
	Body Primitives
}

// SetUserImpl implements SetUserResponse(default) for SetUserdefaultWrappedResponse
func (SetUserdefaultWrappedResponse) SetUserImpl() {}

// HTTPStatusNotModified is an empty response
type HTTPStatusNotModified struct{}

// HTTPStatusOk is an empty response
type HTTPStatusOk struct{}

/*
Here is a copy pastable list of function signatures
for implementing the main interface

// Authenticate post /auth
func (a API) Authenticate(w http.ResponseWriter, r *http.Request) (oa3gen.AuthenticateResponse, error) {
    panic("not implemented")
}
// TestInlinePrimitiveBody get /test/inline
func (a API) TestInlinePrimitiveBody(w http.ResponseWriter, r *http.Request, body string) (oa3gen.TestInlinePrimitiveBodyResponse, error) {
    panic("not implemented")
}
// TestInline post /test/inline
func (a API) TestInline(w http.ResponseWriter, r *http.Request, body oa3gen.TestInlineInline) (oa3gen.TestInlineResponse, error) {
    panic("not implemented")
}
// GetUser get /users/{id}
// Retrieves a user with a long description that spans multiple lines so
// that we can see that both wrapping and long-line support is not
// bleeding over the sacred 80 char limit.
func (a API) GetUser(w http.ResponseWriter, r *http.Request, id string, validStr omitnull.Val[string], reqValidStr null.Val[string], validInt omit.Val[int], reqValidInt int, validNum omit.Val[float64], reqValidNum float64, validBool omit.Val[bool], reqValidBool bool, reqStrFormat string, dateTime chrono.DateTime, date chrono.Date, timeVal chrono.Time, durationVal time.Duration) (oa3gen.GetUserResponse, error) {
    panic("not implemented")
}
// SetUser post /users/{id}
// Sets a user
func (a API) SetUser(w http.ResponseWriter, r *http.Request, body *oa3gen.Primitives) (oa3gen.SetUserResponse, error) {
    panic("not implemented")
}
*/
