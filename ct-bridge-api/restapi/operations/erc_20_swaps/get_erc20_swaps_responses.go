// Code generated by go-swagger; DO NOT EDIT.

package erc_20_swaps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/qydata/ct-evm-compatible-bridge-api/models"
)

// GetErc20SwapsOKCode is the HTTP code returned for type GetErc20SwapsOK
const GetErc20SwapsOKCode int = 200

/*GetErc20SwapsOK Success

swagger:response getErc20SwapsOK
*/
type GetErc20SwapsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Erc20Swaps `json:"body,omitempty"`
}

// NewGetErc20SwapsOK creates GetErc20SwapsOK with default headers values
func NewGetErc20SwapsOK() *GetErc20SwapsOK {

	return &GetErc20SwapsOK{}
}

// WithPayload adds the payload to the get erc20 swaps o k response
func (o *GetErc20SwapsOK) WithPayload(payload *models.Erc20Swaps) *GetErc20SwapsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get erc20 swaps o k response
func (o *GetErc20SwapsOK) SetPayload(payload *models.Erc20Swaps) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetErc20SwapsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetErc20SwapsBadRequestCode is the HTTP code returned for type GetErc20SwapsBadRequest
const GetErc20SwapsBadRequestCode int = 400

/*GetErc20SwapsBadRequest Bad Request

swagger:response getErc20SwapsBadRequest
*/
type GetErc20SwapsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetErc20SwapsBadRequest creates GetErc20SwapsBadRequest with default headers values
func NewGetErc20SwapsBadRequest() *GetErc20SwapsBadRequest {

	return &GetErc20SwapsBadRequest{}
}

// WithPayload adds the payload to the get erc20 swaps bad request response
func (o *GetErc20SwapsBadRequest) WithPayload(payload *models.Error) *GetErc20SwapsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get erc20 swaps bad request response
func (o *GetErc20SwapsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetErc20SwapsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetErc20SwapsInternalServerErrorCode is the HTTP code returned for type GetErc20SwapsInternalServerError
const GetErc20SwapsInternalServerErrorCode int = 500

/*GetErc20SwapsInternalServerError internal server error

swagger:response getErc20SwapsInternalServerError
*/
type GetErc20SwapsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetErc20SwapsInternalServerError creates GetErc20SwapsInternalServerError with default headers values
func NewGetErc20SwapsInternalServerError() *GetErc20SwapsInternalServerError {

	return &GetErc20SwapsInternalServerError{}
}

// WithPayload adds the payload to the get erc20 swaps internal server error response
func (o *GetErc20SwapsInternalServerError) WithPayload(payload *models.Error) *GetErc20SwapsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get erc20 swaps internal server error response
func (o *GetErc20SwapsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetErc20SwapsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
