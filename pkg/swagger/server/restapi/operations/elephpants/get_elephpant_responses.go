// Code generated by go-swagger; DO NOT EDIT.

package elephpants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/scraly/elephpants-api/pkg/swagger/server/models"
)

// GetElephpantOKCode is the HTTP code returned for type GetElephpantOK
const GetElephpantOKCode int = 200

/*
GetElephpantOK A elephpant

swagger:response getElephpantOK
*/
type GetElephpantOK struct {
	/*

	 */
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`

	/*
	  In: Body
	*/
	Payload *models.Elephpant `json:"body,omitempty"`
}

// NewGetElephpantOK creates GetElephpantOK with default headers values
func NewGetElephpantOK() *GetElephpantOK {

	return &GetElephpantOK{}
}

// WithAccessControlAllowOrigin adds the accessControlAllowOrigin to the get elephpant o k response
func (o *GetElephpantOK) WithAccessControlAllowOrigin(accessControlAllowOrigin string) *GetElephpantOK {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
	return o
}

// SetAccessControlAllowOrigin sets the accessControlAllowOrigin to the get elephpant o k response
func (o *GetElephpantOK) SetAccessControlAllowOrigin(accessControlAllowOrigin string) {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
}

// WithPayload adds the payload to the get elephpant o k response
func (o *GetElephpantOK) WithPayload(payload *models.Elephpant) *GetElephpantOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get elephpant o k response
func (o *GetElephpantOK) SetPayload(payload *models.Elephpant) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetElephpantOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Access-Control-Allow-Origin

	accessControlAllowOrigin := o.AccessControlAllowOrigin
	if accessControlAllowOrigin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", accessControlAllowOrigin)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetElephpantNotFoundCode is the HTTP code returned for type GetElephpantNotFound
const GetElephpantNotFoundCode int = 404

/*
GetElephpantNotFound A elephpant with the specified Name was not found.

swagger:response getElephpantNotFound
*/
type GetElephpantNotFound struct {
	/*

	 */
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`
}

// NewGetElephpantNotFound creates GetElephpantNotFound with default headers values
func NewGetElephpantNotFound() *GetElephpantNotFound {

	return &GetElephpantNotFound{}
}

// WithAccessControlAllowOrigin adds the accessControlAllowOrigin to the get elephpant not found response
func (o *GetElephpantNotFound) WithAccessControlAllowOrigin(accessControlAllowOrigin string) *GetElephpantNotFound {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
	return o
}

// SetAccessControlAllowOrigin sets the accessControlAllowOrigin to the get elephpant not found response
func (o *GetElephpantNotFound) SetAccessControlAllowOrigin(accessControlAllowOrigin string) {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
}

// WriteResponse to the client
func (o *GetElephpantNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Access-Control-Allow-Origin

	accessControlAllowOrigin := o.AccessControlAllowOrigin
	if accessControlAllowOrigin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", accessControlAllowOrigin)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
