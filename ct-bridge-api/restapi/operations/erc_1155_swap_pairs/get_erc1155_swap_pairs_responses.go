// Code generated by go-swagger; DO NOT EDIT.

package erc_1155_swap_pairs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/qydata/ct-evm-compatible-bridge-api/models"
)

// GetErc1155SwapPairsOKCode is the HTTP code returned for type GetErc1155SwapPairsOK
const GetErc1155SwapPairsOKCode int = 200

/*GetErc1155SwapPairsOK Success

swagger:response getErc1155SwapPairsOK
*/
type GetErc1155SwapPairsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Erc1155SwapPairs `json:"body,omitempty"`
}

// NewGetErc1155SwapPairsOK creates GetErc1155SwapPairsOK with default headers values
func NewGetErc1155SwapPairsOK() *GetErc1155SwapPairsOK {

	return &GetErc1155SwapPairsOK{}
}

// WithPayload adds the payload to the get erc1155 swap pairs o k response
func (o *GetErc1155SwapPairsOK) WithPayload(payload *models.Erc1155SwapPairs) *GetErc1155SwapPairsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get erc1155 swap pairs o k response
func (o *GetErc1155SwapPairsOK) SetPayload(payload *models.Erc1155SwapPairs) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetErc1155SwapPairsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetErc1155SwapPairsBadRequestCode is the HTTP code returned for type GetErc1155SwapPairsBadRequest
const GetErc1155SwapPairsBadRequestCode int = 400

/*GetErc1155SwapPairsBadRequest Bad Request

swagger:response getErc1155SwapPairsBadRequest
*/
type GetErc1155SwapPairsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetErc1155SwapPairsBadRequest creates GetErc1155SwapPairsBadRequest with default headers values
func NewGetErc1155SwapPairsBadRequest() *GetErc1155SwapPairsBadRequest {

	return &GetErc1155SwapPairsBadRequest{}
}

// WithPayload adds the payload to the get erc1155 swap pairs bad request response
func (o *GetErc1155SwapPairsBadRequest) WithPayload(payload *models.Error) *GetErc1155SwapPairsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get erc1155 swap pairs bad request response
func (o *GetErc1155SwapPairsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetErc1155SwapPairsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetErc1155SwapPairsInternalServerErrorCode is the HTTP code returned for type GetErc1155SwapPairsInternalServerError
const GetErc1155SwapPairsInternalServerErrorCode int = 500

/*GetErc1155SwapPairsInternalServerError internal server error

swagger:response getErc1155SwapPairsInternalServerError
*/
type GetErc1155SwapPairsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetErc1155SwapPairsInternalServerError creates GetErc1155SwapPairsInternalServerError with default headers values
func NewGetErc1155SwapPairsInternalServerError() *GetErc1155SwapPairsInternalServerError {

	return &GetErc1155SwapPairsInternalServerError{}
}

// WithPayload adds the payload to the get erc1155 swap pairs internal server error response
func (o *GetErc1155SwapPairsInternalServerError) WithPayload(payload *models.Error) *GetErc1155SwapPairsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get erc1155 swap pairs internal server error response
func (o *GetErc1155SwapPairsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetErc1155SwapPairsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
