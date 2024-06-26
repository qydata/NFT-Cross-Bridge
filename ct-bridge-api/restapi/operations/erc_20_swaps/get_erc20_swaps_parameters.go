// Code generated by go-swagger; DO NOT EDIT.

package erc_20_swaps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetErc20SwapsParams creates a new GetErc20SwapsParams object
// with the default values initialized.
func NewGetErc20SwapsParams() GetErc20SwapsParams {

	var (
		// initialize parameters with default values

		limitDefault  = int32(100)
		offsetDefault = int32(0)
	)

	return GetErc20SwapsParams{
		Limit: &limitDefault,

		Offset: &offsetDefault,
	}
}

// GetErc20SwapsParams contains all the bound params for the get erc20 swaps operation
// typically these are obtained from a http.Request
//
// swagger:parameters getErc20Swaps
type GetErc20SwapsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*limit
	  Maximum: 1000
	  Minimum: 0
	  In: query
	  Default: 100
	*/
	Limit *int32
	/*offset
	  Maximum: 1000
	  Minimum: 0
	  In: query
	  Default: 0
	*/
	Offset *int32
	/*request_tx_hash
	  Pattern: ^(0x)[0-9A-Fa-f]{64}$
	  In: query
	*/
	RequestTxHash *string
	/*address
	  Required: true
	  Pattern: ^(0x)[0-9A-Fa-f]{40}$
	  In: query
	*/
	Sender string
	/*state
	  In: query
	*/
	State *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetErc20SwapsParams() beforehand.
func (o *GetErc20SwapsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	qRequestTxHash, qhkRequestTxHash, _ := qs.GetOK("request_tx_hash")
	if err := o.bindRequestTxHash(qRequestTxHash, qhkRequestTxHash, route.Formats); err != nil {
		res = append(res, err)
	}

	qSender, qhkSender, _ := qs.GetOK("sender")
	if err := o.bindSender(qSender, qhkSender, route.Formats); err != nil {
		res = append(res, err)
	}

	qState, qhkState, _ := qs.GetOK("state")
	if err := o.bindState(qState, qhkState, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetErc20SwapsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetErc20SwapsParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int32", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *GetErc20SwapsParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", int64(*o.Limit), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "query", int64(*o.Limit), 1000, false); err != nil {
		return err
	}

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *GetErc20SwapsParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetErc20SwapsParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int32", raw)
	}
	o.Offset = &value

	if err := o.validateOffset(formats); err != nil {
		return err
	}

	return nil
}

// validateOffset carries on validations for parameter Offset
func (o *GetErc20SwapsParams) validateOffset(formats strfmt.Registry) error {

	if err := validate.MinimumInt("offset", "query", int64(*o.Offset), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("offset", "query", int64(*o.Offset), 1000, false); err != nil {
		return err
	}

	return nil
}

// bindRequestTxHash binds and validates parameter RequestTxHash from query.
func (o *GetErc20SwapsParams) bindRequestTxHash(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.RequestTxHash = &raw

	if err := o.validateRequestTxHash(formats); err != nil {
		return err
	}

	return nil
}

// validateRequestTxHash carries on validations for parameter RequestTxHash
func (o *GetErc20SwapsParams) validateRequestTxHash(formats strfmt.Registry) error {

	if err := validate.Pattern("request_tx_hash", "query", *o.RequestTxHash, `^(0x)[0-9A-Fa-f]{64}$`); err != nil {
		return err
	}

	return nil
}

// bindSender binds and validates parameter Sender from query.
func (o *GetErc20SwapsParams) bindSender(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("sender", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("sender", "query", raw); err != nil {
		return err
	}
	o.Sender = raw

	if err := o.validateSender(formats); err != nil {
		return err
	}

	return nil
}

// validateSender carries on validations for parameter Sender
func (o *GetErc20SwapsParams) validateSender(formats strfmt.Registry) error {

	if err := validate.Pattern("sender", "query", o.Sender, `^(0x)[0-9A-Fa-f]{40}$`); err != nil {
		return err
	}

	return nil
}

// bindState binds and validates parameter State from query.
func (o *GetErc20SwapsParams) bindState(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.State = &raw

	if err := o.validateState(formats); err != nil {
		return err
	}

	return nil
}

// validateState carries on validations for parameter State
func (o *GetErc20SwapsParams) validateState(formats strfmt.Registry) error {

	if err := validate.EnumCase("state", "query", *o.State, []interface{}{"request_ongoing", "request_rejected", "request_confirmed", "fill_tx_dry_run_failed", "fill_tx_created", "fill_tx_sent", "fill_tx_confirmed", "fill_tx_failed", "fill_tx_missing"}, true); err != nil {
		return err
	}

	return nil
}
