package token

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	contractabi "github.com/qydata/ct-evm-compatible-bridge-core/abi"
)

type IToken interface {
	NAME(opts *bind.CallOpts, tokenAddr string) (string, error)
	SYMBOL(opts *bind.CallOpts, tokenAddr string) (string, error)
	DECIMALS(opts *bind.CallOpts, tokenAddr string) (uint8, error)
}

type Token struct {
	client *ethclient.Client
}

func (t *Token) NAME(opts *bind.CallOpts, tokenAddr string) (string, error) {
	println(tokenAddr)
	token, err := t.bind(tokenAddr)
	if err != nil {
		return "", errors.Wrap(err, "[Token.name]: failed to bind token")
	}
	name, err := token.Name(opts)
	if err != nil {
		return "", errors.Wrap(err, "[Token.name]: failed to get Name")
	}

	return name, nil
}

func (t *Token) SYMBOL(opts *bind.CallOpts, tokenAddr string) (string, error) {
	println(tokenAddr)
	token, err := t.bind(tokenAddr)
	if err != nil {
		return "", errors.Wrap(err, "[Token.symbol]: failed to bind token")
	}

	symbol, err := token.Symbol(opts)
	if err != nil {
		return "", errors.Wrap(err, "[Token.symbol]: failed to get token symbol")
	}

	return symbol, nil
}

func (t *Token) DECIMALS(opts *bind.CallOpts, tokenAddr string) (uint8, error) {
	println(tokenAddr)
	token, err := t.bind(tokenAddr)
	if err != nil {
		return 0, errors.Wrap(err, "[Token.decimals]: failed to bind token")
	}

	decimals, err := token.Decimals(opts)
	if err != nil {
		return 0, errors.Wrap(err, "[Token.decimals]: failed to get token decimals")
	}

	return decimals, nil
}

func NewToken(c *ethclient.Client) *Token {
	return &Token{
		client: c,
	}
}

func (t *Token) bind(addr string) (*contractabi.ERC20Token, error) {
	tokenAddr := common.HexToAddress(addr)
	token, err := contractabi.NewERC20Token(tokenAddr, t.client)
	if err != nil {
		return nil, errors.Wrap(err, "[Token.bind]: failed to bind token address")
	}

	return token, nil
}
