// Code generated by go-swagger; DO NOT EDIT.

package routemanagement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetGreetingOKCode is the HTTP code returned for type GetGreetingOK
const GetGreetingOKCode int = 200

/*GetGreetingOK returns a greeting

swagger:response getGreetingOK
*/
type GetGreetingOK struct {

	/*contains the actual greeting as plain text
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetGreetingOK creates GetGreetingOK with default headers values
func NewGetGreetingOK() *GetGreetingOK {

	return &GetGreetingOK{}
}

// WithPayload adds the payload to the get greeting o k response
func (o *GetGreetingOK) WithPayload(payload string) *GetGreetingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get greeting o k response
func (o *GetGreetingOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGreetingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
