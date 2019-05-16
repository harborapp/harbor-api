// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/umschlag/umschlag-api/pkg/api/v1/models"
)

// UserTeamDeleteOKCode is the HTTP code returned for type UserTeamDeleteOK
const UserTeamDeleteOKCode int = 200

/*UserTeamDeleteOK Plain success message

swagger:response userTeamDeleteOK
*/
type UserTeamDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserTeamDeleteOK creates UserTeamDeleteOK with default headers values
func NewUserTeamDeleteOK() *UserTeamDeleteOK {

	return &UserTeamDeleteOK{}
}

// WithPayload adds the payload to the user team delete o k response
func (o *UserTeamDeleteOK) WithPayload(payload *models.GeneralError) *UserTeamDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user team delete o k response
func (o *UserTeamDeleteOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserTeamDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserTeamDeleteForbiddenCode is the HTTP code returned for type UserTeamDeleteForbidden
const UserTeamDeleteForbiddenCode int = 403

/*UserTeamDeleteForbidden User is not authorized

swagger:response userTeamDeleteForbidden
*/
type UserTeamDeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserTeamDeleteForbidden creates UserTeamDeleteForbidden with default headers values
func NewUserTeamDeleteForbidden() *UserTeamDeleteForbidden {

	return &UserTeamDeleteForbidden{}
}

// WithPayload adds the payload to the user team delete forbidden response
func (o *UserTeamDeleteForbidden) WithPayload(payload *models.GeneralError) *UserTeamDeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user team delete forbidden response
func (o *UserTeamDeleteForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserTeamDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserTeamDeletePreconditionFailedCode is the HTTP code returned for type UserTeamDeletePreconditionFailed
const UserTeamDeletePreconditionFailedCode int = 412

/*UserTeamDeletePreconditionFailed Failed to parse request body

swagger:response userTeamDeletePreconditionFailed
*/
type UserTeamDeletePreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserTeamDeletePreconditionFailed creates UserTeamDeletePreconditionFailed with default headers values
func NewUserTeamDeletePreconditionFailed() *UserTeamDeletePreconditionFailed {

	return &UserTeamDeletePreconditionFailed{}
}

// WithPayload adds the payload to the user team delete precondition failed response
func (o *UserTeamDeletePreconditionFailed) WithPayload(payload *models.GeneralError) *UserTeamDeletePreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user team delete precondition failed response
func (o *UserTeamDeletePreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserTeamDeletePreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserTeamDeleteUnprocessableEntityCode is the HTTP code returned for type UserTeamDeleteUnprocessableEntity
const UserTeamDeleteUnprocessableEntityCode int = 422

/*UserTeamDeleteUnprocessableEntity Team is not assigned

swagger:response userTeamDeleteUnprocessableEntity
*/
type UserTeamDeleteUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserTeamDeleteUnprocessableEntity creates UserTeamDeleteUnprocessableEntity with default headers values
func NewUserTeamDeleteUnprocessableEntity() *UserTeamDeleteUnprocessableEntity {

	return &UserTeamDeleteUnprocessableEntity{}
}

// WithPayload adds the payload to the user team delete unprocessable entity response
func (o *UserTeamDeleteUnprocessableEntity) WithPayload(payload *models.GeneralError) *UserTeamDeleteUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user team delete unprocessable entity response
func (o *UserTeamDeleteUnprocessableEntity) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserTeamDeleteUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UserTeamDeleteDefault Some error unrelated to the handler

swagger:response userTeamDeleteDefault
*/
type UserTeamDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserTeamDeleteDefault creates UserTeamDeleteDefault with default headers values
func NewUserTeamDeleteDefault(code int) *UserTeamDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &UserTeamDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user team delete default response
func (o *UserTeamDeleteDefault) WithStatusCode(code int) *UserTeamDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user team delete default response
func (o *UserTeamDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user team delete default response
func (o *UserTeamDeleteDefault) WithPayload(payload *models.GeneralError) *UserTeamDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user team delete default response
func (o *UserTeamDeleteDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserTeamDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
