package token

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	contractabi "github.com/qydata/ct-evm-compatible-bridge-core/abi"
)

type IToken interface {
	URI(opts *bind.CallOpts, tokenAddr string) (string, error)
}

type Token struct {
	client *ethclient.Client
}

func NewToken(c *ethclient.Client) *Token {
	return &Token{
		client: c,
	}
}

func (t *Token) bind(addr string) (*contractabi.ERC1155Token, error) {
	tokenAddr := common.HexToAddress(addr)
	token, err := contractabi.NewERC1155Token(tokenAddr, t.client)
	if err != nil {
		return nil, errors.Wrap(err, "[Token.bind]: failed to bind token address")
	}

	return token, nil
}

func (t *Token) URI(opts *bind.CallOpts, tokenAddr string) (string, error) {
	println(tokenAddr)
	token, err := t.bind(tokenAddr)
	if err != nil {
		return "", errors.Wrap(err, "[Token.URI]: failed to bind token")
	}

	uri, err := token.Uri(opts, big.NewInt(1))    // TODO Token Id need  edit
	if err != nil {
		return "", errors.Wrap(err, "[Token.URI]: failed to get token URI")
	}

	return uri, nil
}
