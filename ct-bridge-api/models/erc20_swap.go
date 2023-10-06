// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Erc20Swap erc20 swap
//
// swagger:model Erc20Swap
type Erc20Swap struct {

	// amounts
	Amount string `json:"amount"`

	// created at
	CreatedAt string `json:"created_at"`

	// dst chain id
	DstChainID string `json:"dst_chain_id"`

	// dst token addr
	DstTokenAddr string `json:"dst_token_addr"`

	// fill tx hash
	FillTxHash string `json:"fill_tx_hash"`

	// ids
	Id string `json:"id"`

	// recipient
	Recipient string `json:"recipient"`

	// request tx hash
	RequestTxHash string `json:"request_tx_hash"`

	// sender
	Sender string `json:"sender"`

	// src chain id
	SrcChainID string `json:"src_chain_id"`

	// src token addr
	SrcTokenAddr string `json:"src_token_addr"`

	// state
	State string `json:"state"`

	// swap direction
	SwapDirection string `json:"swap_direction"`

	// updated at
	UpdatedAt string `json:"updated_at"`
}

// Validate validates this erc20 swap
func (m *Erc20Swap) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this erc20 swap based on context it is used
func (m *Erc20Swap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Erc20Swap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Erc20Swap) UnmarshalBinary(b []byte) error {
	var res Erc20Swap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}