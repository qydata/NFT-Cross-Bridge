// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceInfo service info
//
// swagger:model ServiceInfo
type ServiceInfo struct {

	// Ct chain id
	CtChainID int64 `json:"ct_chain_id"`

	// Ct erc 1155 swap agent
	CtErc1155SwapAgent string `json:"ct_erc_1155_swap_agent"`

	// bsc erc 721 swap agent
	CtErc721SwapAgent string `json:"ct_erc_721_swap_agent"`

	// bsc erc 20 swap agent
	CtErc20SwapAgent string `json:"ct_erc_20_swap_agent"`

	// coo chain id
	CooChainID int64 `json:"coo_chain_id"`

	// bsc erc 1155 swap agent
	CooErc1155SwapAgent string `json:"coo_erc_1155_swap_agent"`

	// Coo erc 721 swap agent
	CooErc721SwapAgent string `json:"coo_erc_721_swap_agent"`

	// Coo erc 721 swap agent
	CooErc20SwapAgent string `json:"coo_erc_20_swap_agent"`

	// 25
	SparkchainChainID                 int64  `json:"sparkchain_chain_id"`
	SparkchainErc20SwapAgent          string `json:"sparkchain_erc_20_swap_agent"`
	SparkchainErc721SwapAgent         string `json:"sparkchain_erc_721_swap_agent"`
	SparkchainErc1155SwapAgent        string `json:"sparkchain_erc_1155_swap_agent"`
	ChainmakerChainID                 int64  `json:"chainmaker_chain_id"`
	ChainmakerErc20SwapAgent          string `json:"chainmaker_erc_20_swap_agent"`
	ChainmakerErc721SwapAgent         string `json:"chainmaker_erc_721_swap_agent"`
	ChainmakerErc1155SwapAgent        string `json:"chainmaker_erc_1155_swap_agent"`
	BsnChainID                        int64  `json:"bsn_chain_id"`
	BsnErc20SwapAgent                 string `json:"bsn_erc_20_swap_agent"`
	BsnErc721SwapAgent                string `json:"bsn_erc_721_swap_agent"`
	BsnErc1155SwapAgent               string `json:"bsn_erc_1155_swap_agent"`
	HyperchainChainID                 int64  `json:"hyperchain_chain_id"`
	HyperchainErc20SwapAgent          string `json:"hyperchain_erc_20_swap_agent"`
	HyperchainErc721SwapAgent         string `json:"hyperchain_erc_721_swap_agent"`
	HyperchainErc1155SwapAgent        string `json:"hyperchain_erc_1155_swap_agent"`
	BubichainChainID                  int64  `json:"bubichain_chain_id"`
	BubichainErc20SwapAgent           string `json:"bubichain_erc_20_swap_agent"`
	BubichainErc721SwapAgent          string `json:"bubichain_erc_721_swap_agent"`
	BubichainErc1155SwapAgent         string `json:"bubichain_erc_1155_swap_agent"`
	AntblockchainChainID              int64  `json:"antblockchain_chain_id"`
	AntblockchainErc20SwapAgent       string `json:"antblockchain_erc_20_swap_agent"`
	AntblockchainErc721SwapAgent      string `json:"antblockchain_erc_721_swap_agent"`
	AntblockchainErc1155SwapAgent     string `json:"antblockchain_erc_1155_swap_agent"`
	TencentblockchainChainID          int64  `json:"tencentblockchain_chain_id"`
	TencentblockchainErc20SwapAgent   string `json:"tencentblockchain_erc_20_swap_agent"`
	TencentblockchainErc721SwapAgent  string `json:"tencentblockchain_erc_721_swap_agent"`
	TencentblockchainErc1155SwapAgent string `json:"tencentblockchain_erc_1155_swap_agent"`
	JdblockchainChainID               int64  `json:"jdblockchain_chain_id"`
	JdblockchainErc20SwapAgent        string `json:"jdblockchain_erc_20_swap_agent"`
	JdblockchainErc721SwapAgent       string `json:"jdblockchain_erc_721_swap_agent"`
	JdblockchainErc1155SwapAgent      string `json:"jdblockchain_erc_1155_swap_agent"`
	XuperchainChainID                 int64  `json:"xuperchain_chain_id"`
	XuperchainErc20SwapAgent          string `json:"xuperchain_erc_20_swap_agent"`
	XuperchainErc721SwapAgent         string `json:"xuperchain_erc_721_swap_agent"`
	XuperchainErc1155SwapAgent        string `json:"xuperchain_erc_1155_swap_agent"`
	NeoChainID                        int64  `json:"neo_chain_id"`
	NeoErc20SwapAgent                 string `json:"neo_erc_20_swap_agent"`
	NeoErc721SwapAgent                string `json:"neo_erc_721_swap_agent"`
	NeoErc1155SwapAgent               string `json:"neo_erc_1155_swap_agent"`
	QtumChainID                       int64  `json:"qtum_chain_id"`
	QtumErc20SwapAgent                string `json:"qtum_erc_20_swap_agent"`
	QtumErc721SwapAgent               string `json:"qtum_erc_721_swap_agent"`
	QtumErc1155SwapAgent              string `json:"qtum_erc_1155_swap_agent"`
	VechainChainID                    int64  `json:"vechain_chain_id"`
	VechainErc20SwapAgent             string `json:"vechain_erc_20_swap_agent"`
	VechainErc721SwapAgent            string `json:"vechain_erc_721_swap_agent"`
	VechainErc1155SwapAgent           string `json:"vechain_erc_1155_swap_agent"`
	GxchainChainID                    int64  `json:"gxchain_chain_id"`
	GxchainErc20SwapAgent             string `json:"gxchain_erc_20_swap_agent"`
	GxchainErc721SwapAgent            string `json:"gxchain_erc_721_swap_agent"`
	GxchainErc1155SwapAgent           string `json:"gxchain_erc_1155_swap_agent"`
	OraclechainChainID                int64  `json:"oraclechain_chain_id"`
	OraclechainErc20SwapAgent         string `json:"oraclechain_erc_20_swap_agent"`
	OraclechainErc721SwapAgent        string `json:"oraclechain_erc_721_swap_agent"`
	OraclechainErc1155SwapAgent       string `json:"oraclechain_erc_1155_swap_agent"`
	OntologyChainID                   int64  `json:"ontology_chain_id"`
	OntologyErc20SwapAgent            string `json:"ontology_erc_20_swap_agent"`
	OntologyErc721SwapAgent           string `json:"ontology_erc_721_swap_agent"`
	OntologyErc1155SwapAgent          string `json:"ontology_erc_1155_swap_agent"`
	BytomChainID                      int64  `json:"bytom_chain_id"`
	BytomErc20SwapAgent               string `json:"bytom_erc_20_swap_agent"`
	BytomErc721SwapAgent              string `json:"bytom_erc_721_swap_agent"`
	BytomErc1155SwapAgent             string `json:"bytom_erc_1155_swap_agent"`
	DarwiniaChainID                   int64  `json:"darwinia_chain_id"`
	DarwiniaErc20SwapAgent            string `json:"darwinia_erc_20_swap_agent"`
	DarwiniaErc721SwapAgent           string `json:"darwinia_erc_721_swap_agent"`
	DarwiniaErc1155SwapAgent          string `json:"darwinia_erc_1155_swap_agent"`
	HecoChainID                       int64  `json:"heco_chain_id"`
	HecoErc20SwapAgent                string `json:"heco_erc_20_swap_agent"`
	HecoErc721SwapAgent               string `json:"heco_erc_721_swap_agent"`
	HecoErc1155SwapAgent              string `json:"heco_erc_1155_swap_agent"`
	PolkadotChainID                   int64  `json:"polkadot_chain_id"`
	PolkadotErc20SwapAgent            string `json:"polkadot_erc_20_swap_agent"`
	PolkadotErc721SwapAgent           string `json:"polkadot_erc_721_swap_agent"`
	PolkadotErc1155SwapAgent          string `json:"polkadot_erc_1155_swap_agent"`
	ConfluxChainID                    int64  `json:"conflux_chain_id"`
	ConfluxErc20SwapAgent             string `json:"conflux_erc_20_swap_agent"`
	ConfluxErc721SwapAgent            string `json:"conflux_erc_721_swap_agent"`
	ConfluxErc1155SwapAgent           string `json:"conflux_erc_1155_swap_agent"`
	BitsharesChainID                  int64  `json:"bitshares_chain_id"`
	BitsharesErc20SwapAgent           string `json:"bitshares_erc_20_swap_agent"`
	BitsharesErc721SwapAgent          string `json:"bitshares_erc_721_swap_agent"`
	BitsharesErc1155SwapAgent         string `json:"bitshares_erc_1155_swap_agent"`
	YoyowChainID                      int64  `json:"yoyow_chain_id"`
	YoyowErc20SwapAgent               string `json:"yoyow_erc_20_swap_agent"`
	YoyowErc721SwapAgent              string `json:"yoyow_erc_721_swap_agent"`
	YoyowErc1155SwapAgent             string `json:"yoyow_erc_1155_swap_agent"`
	UltrainChainID                    int64  `json:"ultrain_chain_id"`
	UltrainErc20SwapAgent             string `json:"ultrain_erc_20_swap_agent"`
	UltrainErc721SwapAgent            string `json:"ultrain_erc_721_swap_agent"`
	UltrainErc1155SwapAgent           string `json:"ultrain_erc_1155_swap_agent"`
	CocosbcxChainID                   int64  `json:"cocosbcx_chain_id"`
	CocosbcxErc20SwapAgent            string `json:"cocosbcx_erc_20_swap_agent"`
	CocosbcxErc721SwapAgent           string `json:"cocosbcx_erc_721_swap_agent"`
	CocosbcxErc1155SwapAgent          string `json:"cocosbcx_erc_1155_swap_agent"`
	AchainChainID                     int64  `json:"achain_chain_id"`
	AchainErc20SwapAgent              string `json:"achain_erc_20_swap_agent"`
	AchainErc721SwapAgent             string `json:"achain_erc_721_swap_agent"`
	AchainErc1155SwapAgent            string `json:"achain_erc_1155_swap_agent"`
}

// Validate validates this service info
func (m *ServiceInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service info based on context it is used
func (m *ServiceInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceInfo) UnmarshalBinary(b []byte) error {
	var res ServiceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
