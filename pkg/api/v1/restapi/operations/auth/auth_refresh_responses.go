// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/umschlag/umschlag-api/pkg/api/v1/models"
)

// AuthRefreshOKCode is the HTTP code returned for type AuthRefreshOK
const AuthRefreshOKCode int = 200

/*AuthRefreshOK A refreshed token with expire

swagger:response authRefreshOK
*/
type AuthRefreshOK struct {

	/*
	  In: Body
	*/
	Payload *models.AuthToken `json:"body,omitempty"`
}

// NewAuthRefreshOK creates AuthRefreshOK with default headers values
func NewAuthRefreshOK() *AuthRefreshOK {

	return &AuthRefreshOK{}
}

// WithPayload adds the payload to the auth refresh o k response
func (o *AuthRefreshOK) WithPayload(payload *models.AuthToken) *AuthRefreshOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the auth refresh o k response
func (o *AuthRefreshOK) SetPayload(payload *models.AuthToken) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthRefreshOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AuthRefreshUnauthorizedCode is the HTTP code returned for type AuthRefreshUnauthorized
const AuthRefreshUnauthorizedCode int = 401

/*AuthRefreshUnauthorized Unauthorized if failed to generate

swagger:response authRefreshUnauthorized
*/
type AuthRefreshUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAuthRefreshUnauthorized creates AuthRefreshUnauthorized with default headers values
func NewAuthRefreshUnauthorized() *AuthRefreshUnauthorized {

	return &AuthRefreshUnauthorized{}
}

// WithPayload adds the payload to the auth refresh unauthorized response
func (o *AuthRefreshUnauthorized) WithPayload(payload *models.GeneralError) *AuthRefreshUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the auth refresh unauthorized response
func (o *AuthRefreshUnauthorized) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthRefreshUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AuthRefreshDefault Some error unrelated to the handler

swagger:response authRefreshDefault
*/
type AuthRefreshDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAuthRefreshDefault creates AuthRefreshDefault with default headers values
func NewAuthRefreshDefault(code int) *AuthRefreshDefault {
	if code <= 0 {
		code = 500
	}

	return &AuthRefreshDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the auth refresh default response
func (o *AuthRefreshDefault) WithStatusCode(code int) *AuthRefreshDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the auth refresh default response
func (o *AuthRefreshDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the auth refresh default response
func (o *AuthRefreshDefault) WithPayload(payload *models.GeneralError) *AuthRefreshDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the auth refresh default response
func (o *AuthRefreshDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthRefreshDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
