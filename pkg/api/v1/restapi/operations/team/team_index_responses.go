// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/umschlag/umschlag-api/pkg/api/v1/models"
)

// TeamIndexOKCode is the HTTP code returned for type TeamIndexOK
const TeamIndexOKCode int = 200

/*TeamIndexOK A collection of teams

swagger:response teamIndexOK
*/
type TeamIndexOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Team `json:"body,omitempty"`
}

// NewTeamIndexOK creates TeamIndexOK with default headers values
func NewTeamIndexOK() *TeamIndexOK {

	return &TeamIndexOK{}
}

// WithPayload adds the payload to the team index o k response
func (o *TeamIndexOK) WithPayload(payload []*models.Team) *TeamIndexOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team index o k response
func (o *TeamIndexOK) SetPayload(payload []*models.Team) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamIndexOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Team, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// TeamIndexForbiddenCode is the HTTP code returned for type TeamIndexForbidden
const TeamIndexForbiddenCode int = 403

/*TeamIndexForbidden User is not authorized

swagger:response teamIndexForbidden
*/
type TeamIndexForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamIndexForbidden creates TeamIndexForbidden with default headers values
func NewTeamIndexForbidden() *TeamIndexForbidden {

	return &TeamIndexForbidden{}
}

// WithPayload adds the payload to the team index forbidden response
func (o *TeamIndexForbidden) WithPayload(payload *models.GeneralError) *TeamIndexForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team index forbidden response
func (o *TeamIndexForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamIndexForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TeamIndexDefault Some error unrelated to the handler

swagger:response teamIndexDefault
*/
type TeamIndexDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamIndexDefault creates TeamIndexDefault with default headers values
func NewTeamIndexDefault(code int) *TeamIndexDefault {
	if code <= 0 {
		code = 500
	}

	return &TeamIndexDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the team index default response
func (o *TeamIndexDefault) WithStatusCode(code int) *TeamIndexDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the team index default response
func (o *TeamIndexDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the team index default response
func (o *TeamIndexDefault) WithPayload(payload *models.GeneralError) *TeamIndexDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team index default response
func (o *TeamIndexDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamIndexDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
