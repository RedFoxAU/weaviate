//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package contextionary_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// C11yExtensionsOKCode is the HTTP code returned for type C11yExtensionsOK
const C11yExtensionsOKCode int = 200

/*C11yExtensionsOK Successfully extended the contextionary with the custom cocnept

swagger:response c11yExtensionsOK
*/
type C11yExtensionsOK struct {

	/*
	  In: Body
	*/
	Payload *models.C11yExtension `json:"body,omitempty"`
}

// NewC11yExtensionsOK creates C11yExtensionsOK with default headers values
func NewC11yExtensionsOK() *C11yExtensionsOK {

	return &C11yExtensionsOK{}
}

// WithPayload adds the payload to the c11y extensions o k response
func (o *C11yExtensionsOK) WithPayload(payload *models.C11yExtension) *C11yExtensionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the c11y extensions o k response
func (o *C11yExtensionsOK) SetPayload(payload *models.C11yExtension) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *C11yExtensionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// C11yExtensionsBadRequestCode is the HTTP code returned for type C11yExtensionsBadRequest
const C11yExtensionsBadRequestCode int = 400

/*C11yExtensionsBadRequest Incorrect request

swagger:response c11yExtensionsBadRequest
*/
type C11yExtensionsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewC11yExtensionsBadRequest creates C11yExtensionsBadRequest with default headers values
func NewC11yExtensionsBadRequest() *C11yExtensionsBadRequest {

	return &C11yExtensionsBadRequest{}
}

// WithPayload adds the payload to the c11y extensions bad request response
func (o *C11yExtensionsBadRequest) WithPayload(payload *models.ErrorResponse) *C11yExtensionsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the c11y extensions bad request response
func (o *C11yExtensionsBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *C11yExtensionsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// C11yExtensionsUnauthorizedCode is the HTTP code returned for type C11yExtensionsUnauthorized
const C11yExtensionsUnauthorizedCode int = 401

/*C11yExtensionsUnauthorized Unauthorized or invalid credentials.

swagger:response c11yExtensionsUnauthorized
*/
type C11yExtensionsUnauthorized struct {
}

// NewC11yExtensionsUnauthorized creates C11yExtensionsUnauthorized with default headers values
func NewC11yExtensionsUnauthorized() *C11yExtensionsUnauthorized {

	return &C11yExtensionsUnauthorized{}
}

// WriteResponse to the client
func (o *C11yExtensionsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// C11yExtensionsForbiddenCode is the HTTP code returned for type C11yExtensionsForbidden
const C11yExtensionsForbiddenCode int = 403

/*C11yExtensionsForbidden Forbidden

swagger:response c11yExtensionsForbidden
*/
type C11yExtensionsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewC11yExtensionsForbidden creates C11yExtensionsForbidden with default headers values
func NewC11yExtensionsForbidden() *C11yExtensionsForbidden {

	return &C11yExtensionsForbidden{}
}

// WithPayload adds the payload to the c11y extensions forbidden response
func (o *C11yExtensionsForbidden) WithPayload(payload *models.ErrorResponse) *C11yExtensionsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the c11y extensions forbidden response
func (o *C11yExtensionsForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *C11yExtensionsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// C11yExtensionsInternalServerErrorCode is the HTTP code returned for type C11yExtensionsInternalServerError
const C11yExtensionsInternalServerErrorCode int = 500

/*C11yExtensionsInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response c11yExtensionsInternalServerError
*/
type C11yExtensionsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewC11yExtensionsInternalServerError creates C11yExtensionsInternalServerError with default headers values
func NewC11yExtensionsInternalServerError() *C11yExtensionsInternalServerError {

	return &C11yExtensionsInternalServerError{}
}

// WithPayload adds the payload to the c11y extensions internal server error response
func (o *C11yExtensionsInternalServerError) WithPayload(payload *models.ErrorResponse) *C11yExtensionsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the c11y extensions internal server error response
func (o *C11yExtensionsInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *C11yExtensionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// C11yExtensionsNotImplementedCode is the HTTP code returned for type C11yExtensionsNotImplemented
const C11yExtensionsNotImplementedCode int = 501

/*C11yExtensionsNotImplemented Not (yet) implemented.

swagger:response c11yExtensionsNotImplemented
*/
type C11yExtensionsNotImplemented struct {
}

// NewC11yExtensionsNotImplemented creates C11yExtensionsNotImplemented with default headers values
func NewC11yExtensionsNotImplemented() *C11yExtensionsNotImplemented {

	return &C11yExtensionsNotImplemented{}
}

// WriteResponse to the client
func (o *C11yExtensionsNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(501)
}
