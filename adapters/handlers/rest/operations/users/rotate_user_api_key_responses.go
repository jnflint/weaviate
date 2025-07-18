//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// RotateUserAPIKeyOKCode is the HTTP code returned for type RotateUserAPIKeyOK
const RotateUserAPIKeyOKCode int = 200

/*
RotateUserAPIKeyOK ApiKey successfully changed

swagger:response rotateUserApiKeyOK
*/
type RotateUserAPIKeyOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserAPIKey `json:"body,omitempty"`
}

// NewRotateUserAPIKeyOK creates RotateUserAPIKeyOK with default headers values
func NewRotateUserAPIKeyOK() *RotateUserAPIKeyOK {

	return &RotateUserAPIKeyOK{}
}

// WithPayload adds the payload to the rotate user Api key o k response
func (o *RotateUserAPIKeyOK) WithPayload(payload *models.UserAPIKey) *RotateUserAPIKeyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rotate user Api key o k response
func (o *RotateUserAPIKeyOK) SetPayload(payload *models.UserAPIKey) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RotateUserAPIKeyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RotateUserAPIKeyBadRequestCode is the HTTP code returned for type RotateUserAPIKeyBadRequest
const RotateUserAPIKeyBadRequestCode int = 400

/*
RotateUserAPIKeyBadRequest Malformed request.

swagger:response rotateUserApiKeyBadRequest
*/
type RotateUserAPIKeyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewRotateUserAPIKeyBadRequest creates RotateUserAPIKeyBadRequest with default headers values
func NewRotateUserAPIKeyBadRequest() *RotateUserAPIKeyBadRequest {

	return &RotateUserAPIKeyBadRequest{}
}

// WithPayload adds the payload to the rotate user Api key bad request response
func (o *RotateUserAPIKeyBadRequest) WithPayload(payload *models.ErrorResponse) *RotateUserAPIKeyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rotate user Api key bad request response
func (o *RotateUserAPIKeyBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RotateUserAPIKeyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RotateUserAPIKeyUnauthorizedCode is the HTTP code returned for type RotateUserAPIKeyUnauthorized
const RotateUserAPIKeyUnauthorizedCode int = 401

/*
RotateUserAPIKeyUnauthorized Unauthorized or invalid credentials.

swagger:response rotateUserApiKeyUnauthorized
*/
type RotateUserAPIKeyUnauthorized struct {
}

// NewRotateUserAPIKeyUnauthorized creates RotateUserAPIKeyUnauthorized with default headers values
func NewRotateUserAPIKeyUnauthorized() *RotateUserAPIKeyUnauthorized {

	return &RotateUserAPIKeyUnauthorized{}
}

// WriteResponse to the client
func (o *RotateUserAPIKeyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// RotateUserAPIKeyForbiddenCode is the HTTP code returned for type RotateUserAPIKeyForbidden
const RotateUserAPIKeyForbiddenCode int = 403

/*
RotateUserAPIKeyForbidden Forbidden

swagger:response rotateUserApiKeyForbidden
*/
type RotateUserAPIKeyForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewRotateUserAPIKeyForbidden creates RotateUserAPIKeyForbidden with default headers values
func NewRotateUserAPIKeyForbidden() *RotateUserAPIKeyForbidden {

	return &RotateUserAPIKeyForbidden{}
}

// WithPayload adds the payload to the rotate user Api key forbidden response
func (o *RotateUserAPIKeyForbidden) WithPayload(payload *models.ErrorResponse) *RotateUserAPIKeyForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rotate user Api key forbidden response
func (o *RotateUserAPIKeyForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RotateUserAPIKeyForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RotateUserAPIKeyNotFoundCode is the HTTP code returned for type RotateUserAPIKeyNotFound
const RotateUserAPIKeyNotFoundCode int = 404

/*
RotateUserAPIKeyNotFound user not found

swagger:response rotateUserApiKeyNotFound
*/
type RotateUserAPIKeyNotFound struct {
}

// NewRotateUserAPIKeyNotFound creates RotateUserAPIKeyNotFound with default headers values
func NewRotateUserAPIKeyNotFound() *RotateUserAPIKeyNotFound {

	return &RotateUserAPIKeyNotFound{}
}

// WriteResponse to the client
func (o *RotateUserAPIKeyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// RotateUserAPIKeyUnprocessableEntityCode is the HTTP code returned for type RotateUserAPIKeyUnprocessableEntity
const RotateUserAPIKeyUnprocessableEntityCode int = 422

/*
RotateUserAPIKeyUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous.

swagger:response rotateUserApiKeyUnprocessableEntity
*/
type RotateUserAPIKeyUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewRotateUserAPIKeyUnprocessableEntity creates RotateUserAPIKeyUnprocessableEntity with default headers values
func NewRotateUserAPIKeyUnprocessableEntity() *RotateUserAPIKeyUnprocessableEntity {

	return &RotateUserAPIKeyUnprocessableEntity{}
}

// WithPayload adds the payload to the rotate user Api key unprocessable entity response
func (o *RotateUserAPIKeyUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *RotateUserAPIKeyUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rotate user Api key unprocessable entity response
func (o *RotateUserAPIKeyUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RotateUserAPIKeyUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RotateUserAPIKeyInternalServerErrorCode is the HTTP code returned for type RotateUserAPIKeyInternalServerError
const RotateUserAPIKeyInternalServerErrorCode int = 500

/*
RotateUserAPIKeyInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response rotateUserApiKeyInternalServerError
*/
type RotateUserAPIKeyInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewRotateUserAPIKeyInternalServerError creates RotateUserAPIKeyInternalServerError with default headers values
func NewRotateUserAPIKeyInternalServerError() *RotateUserAPIKeyInternalServerError {

	return &RotateUserAPIKeyInternalServerError{}
}

// WithPayload adds the payload to the rotate user Api key internal server error response
func (o *RotateUserAPIKeyInternalServerError) WithPayload(payload *models.ErrorResponse) *RotateUserAPIKeyInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rotate user Api key internal server error response
func (o *RotateUserAPIKeyInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RotateUserAPIKeyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
