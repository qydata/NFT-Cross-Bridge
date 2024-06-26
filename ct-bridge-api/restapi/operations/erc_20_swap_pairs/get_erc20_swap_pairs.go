// Code generated by go-swagger; DO NOT EDIT.

package erc_20_swap_pairs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetErc20SwapPairsHandlerFunc turns a function with the right signature into a get erc20 swap pairs handler
type GetErc20SwapPairsHandlerFunc func(GetErc20SwapPairsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetErc20SwapPairsHandlerFunc) Handle(params GetErc20SwapPairsParams) middleware.Responder {
	return fn(params)
}

// GetErc20SwapPairsHandler interface for that can handle valid get erc20 swap pairs params
type GetErc20SwapPairsHandler interface {
	Handle(GetErc20SwapPairsParams) middleware.Responder
}

// NewGetErc20SwapPairs creates a new http.Handler for the get erc20 swap pairs operation
func NewGetErc20SwapPairs(ctx *middleware.Context, handler GetErc20SwapPairsHandler) *GetErc20SwapPairs {
	return &GetErc20SwapPairs{Context: ctx, Handler: handler}
}

/* GetErc20SwapPairs swagger:route GET /v1/erc-20-swap-pairs erc_20_swap_pairs getErc20SwapPairs

Gets a list of available ERC20 swap pairs.

*/
type GetErc20SwapPairs struct {
	Context *middleware.Context
	Handler GetErc20SwapPairsHandler
}

func (o *GetErc20SwapPairs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetErc20SwapPairsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
