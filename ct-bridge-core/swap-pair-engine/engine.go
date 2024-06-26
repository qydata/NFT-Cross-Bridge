package engine

import (
	erc20agent "github.com/qydata/ct-evm-compatible-bridge-core/agent/erc20"
	"math/big"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	erc1155agent "github.com/qydata/ct-evm-compatible-bridge-core/agent/erc1155"
	erc721agent "github.com/qydata/ct-evm-compatible-bridge-core/agent/erc721"
	"github.com/qydata/ct-evm-compatible-bridge-core/client"
	"github.com/qydata/ct-evm-compatible-bridge-core/recorder"
)

type Recorder interface {
}

type Config struct {
	ExplorerURL               string
	PrivateKey                string
	ChainID                   *big.Int
	ConfirmNum                int64
	MaxTrackRetry             int64
	ERC20SwapAgentAddresses   map[string]common.Address
	ERC721SwapAgentAddresses  map[string]common.Address
	ERC1155SwapAgentAddresses map[string]common.Address
}

type Dependencies struct {
	Client           map[string]client.ETHClient
	DB               *gorm.DB
	Recorder         map[string]recorder.IRecorder
	ERC20SwapAgent   map[string]erc20agent.SwapAgent
	ERC721SwapAgent  map[string]erc721agent.SwapAgent
	ERC1155SwapAgent map[string]erc1155agent.SwapAgent
}

type Engine struct {
	conf *Config
	deps *Dependencies
}

func NewEngine(c *Config, d *Dependencies) *Engine {
	return &Engine{
		conf: c,
		deps: d,
	}
}
