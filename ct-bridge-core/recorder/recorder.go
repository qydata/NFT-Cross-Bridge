package recorder

import (
	erc20agent "github.com/qydata/ct-evm-compatible-bridge-core/agent/erc20"
	erc20token "github.com/qydata/ct-evm-compatible-bridge-core/token/erc20"
	"math/big"

	"gorm.io/gorm"

	erc1155agent "github.com/qydata/ct-evm-compatible-bridge-core/agent/erc1155"
	erc721agent "github.com/qydata/ct-evm-compatible-bridge-core/agent/erc721"
	"github.com/qydata/ct-evm-compatible-bridge-core/client"
	corecommon "github.com/qydata/ct-evm-compatible-bridge-core/common"
	"github.com/qydata/ct-evm-compatible-bridge-core/model/block"
	erc1155token "github.com/qydata/ct-evm-compatible-bridge-core/token/erc1155"
	erc721token "github.com/qydata/ct-evm-compatible-bridge-core/token/erc721"
)

type IRecorder interface {
	Block(height int64) (*corecommon.Block, error)
	ChainID() string
	Delete(tx *gorm.DB, height int64) error
	LatestBlockCached() *corecommon.Block
	Record(tx *gorm.DB, block *block.Log) error
}

type Config struct {
	ChainID   *big.Int
	ChainName string
	HMACKey   string
}

type Dependencies struct {
	Client           map[string]client.ETHClient
	DB               *gorm.DB
	ERC20SwapAgent   map[string]erc20agent.SwapAgent
	ERC20Token       map[string]erc20token.IToken
	ERC721SwapAgent  map[string]erc721agent.SwapAgent
	ERC721Token      map[string]erc721token.IToken
	ERC1155SwapAgent map[string]erc1155agent.SwapAgent
	ERC1155Token     map[string]erc1155token.IToken
}

type Recorder struct {
	latestBlockCached *corecommon.Block
	conf              *Config
	deps              *Dependencies
}

func NewRecorder(c *Config, d *Dependencies) *Recorder {
	return &Recorder{
		conf: c,
		deps: d,
	}
}
