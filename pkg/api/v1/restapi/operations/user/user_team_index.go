// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UserTeamIndexHandlerFunc turns a function with the right signature into a user team index handler
type UserTeamIndexHandlerFunc func(UserTeamIndexParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserTeamIndexHandlerFunc) Handle(params UserTeamIndexParams) middleware.Responder {
	return fn(params)
}

// UserTeamIndexHandler interface for that can handle valid user team index params
type UserTeamIndexHandler interface {
	Handle(UserTeamIndexParams) middleware.Responder
}

// NewUserTeamIndex creates a new http.Handler for the user team index operation
func NewUserTeamIndex(ctx *middleware.Context, handler UserTeamIndexHandler) *UserTeamIndex {
	return &UserTeamIndex{Context: ctx, Handler: handler}
}

/*UserTeamIndex swagger:route GET /users/{userID}/teams user userTeamIndex

Fetch all teams assigned to user

*/
type UserTeamIndex struct {
	Context *middleware.Context
	Handler UserTeamIndexHandler
}

func (o *UserTeamIndex) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserTeamIndexParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
