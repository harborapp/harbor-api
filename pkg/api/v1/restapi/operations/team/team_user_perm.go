// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// TeamUserPermHandlerFunc turns a function with the right signature into a team user perm handler
type TeamUserPermHandlerFunc func(TeamUserPermParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TeamUserPermHandlerFunc) Handle(params TeamUserPermParams) middleware.Responder {
	return fn(params)
}

// TeamUserPermHandler interface for that can handle valid team user perm params
type TeamUserPermHandler interface {
	Handle(TeamUserPermParams) middleware.Responder
}

// NewTeamUserPerm creates a new http.Handler for the team user perm operation
func NewTeamUserPerm(ctx *middleware.Context, handler TeamUserPermHandler) *TeamUserPerm {
	return &TeamUserPerm{Context: ctx, Handler: handler}
}

/*TeamUserPerm swagger:route PUT /teams/{teamID}/users team teamUserPerm

Update user perms for team

*/
type TeamUserPerm struct {
	Context *middleware.Context
	Handler TeamUserPermHandler
}

func (o *TeamUserPerm) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTeamUserPermParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
