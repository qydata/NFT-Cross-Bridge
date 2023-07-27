// Code generated by go-swagger; DO NOT EDIT.

package erc_721_swap_pairs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/qydata/ct-evm-compatible-bridge-api/models"
)

// GetErc721SwapPairsOKCode is the HTTP code returned for type GetErc721SwapPairsOK
const GetErc721SwapPairsOKCode int = 200

/*GetErc721SwapPairsOK Success

swagger:response getErc721SwapPairsOK
*/
type GetErc721SwapPairsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Erc721SwapPairs `json:"body,omitempty"`
}

// NewGetErc721SwapPairsOK creates GetErc721SwapPairsOK with default headers values
func NewGetErc721SwapPairsOK() *GetErc721SwapPairsOK {

	return &GetErc721SwapPairsOK{}
}

// WithPayload adds the payload to the get erc721 swap pairs o k response
func (o *GetErc721SwapPairsOK) WithPayload(payload *models.Erc721SwapPairs) *GetErc721SwapPairsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get erc721 swap pairs o k response
func (o *GetErc721SwapPairsOK) SetPayload(payload *models.Erc721SwapPairs) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetErc721SwapPairsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetErc721SwapPairsBadRequestCode is the HTTP code returned for type GetErc721SwapPairsBadRequest
const GetErc721SwapPairsBadRequestCode int = 400

/*GetErc721SwapPairsBadRequest Bad Request

swagger:response getErc721SwapPairsBadRequest
*/
type GetErc721SwapPairsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetErc721SwapPairsBadRequest creates GetErc721SwapPairsBadRequest with default headers values
func NewGetErc721SwapPairsBadRequest() *GetErc721SwapPairsBadRequest {

	return &GetErc721SwapPairsBadRequest{}
}

// WithPayload adds the payload to the get erc721 swap pairs bad request response
func (o *GetErc721SwapPairsBadRequest) WithPayload(payload *models.Error) *GetErc721SwapPairsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get erc721 swap pairs bad request response
func (o *GetErc721SwapPairsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetErc721SwapPairsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetErc721SwapPairsInternalServerErrorCode is the HTTP code returned for type GetErc721SwapPairsInternalServerError
const GetErc721SwapPairsInternalServerErrorCode int = 500

/*GetErc721SwapPairsInternalServerError internal server error

swagger:response getErc721SwapPairsInternalServerError
*/
type GetErc721SwapPairsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetErc721SwapPairsInternalServerError creates GetErc721SwapPairsInternalServerError with default headers values
func NewGetErc721SwapPairsInternalServerError() *GetErc721SwapPairsInternalServerError {

	return &GetErc721SwapPairsInternalServerError{}
}

// WithPayload adds the payload to the get erc721 swap pairs internal server error response
func (o *GetErc721SwapPairsInternalServerError) WithPayload(payload *models.Error) *GetErc721SwapPairsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get erc721 swap pairs internal server error response
func (o *GetErc721SwapPairsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetErc721SwapPairsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}