package config

import (
	"encoding/json"
	"io/ioutil"

	svc "github.com/qydata/ct-evm-compatible-bridge-api/services"
	ulog "github.com/qydata/ct-evm-compatible-bridge-api/utils/log"
)

// Config represents configuration from config.json.
type Config struct {
	Logs       ulog.LogsConfig  `json:"logs"`
	CacheTTLs  map[string]int64 `json:"cache_ttls"`
	DB         DBConfig         `json:"db"`
	SwapConfig SwapConfig       `json:"swap_config"`
	CorsConfig CorsConfig       `json:"cors"`
}

type Consumer struct {
	Username string
	Key      string
}

type SwapConfig struct {
	CtChainID              int64  `json:"ct_chain_id"`
	CtErc20SwapAgent      string `json:"ct_erc_20_swap_agent"`
	CtErc721SwapAgent      string `json:"ct_erc_721_swap_agent"`
	CtErc1155SwapAgent     string `json:"ct_erc_1155_swap_agent"`
	CooChainID              int64  `json:"coo_chain_id"`
	CooErc20SwapAgent      string `json:"coo_erc_20_swap_agent"`
	CooErc721SwapAgent      string `json:"coo_erc_721_swap_agent"`
	CooErc1155SwapAgent     string `json:"coo_erc_1155_swap_agent"`
}

type DBConfig struct {
	DSN      string `json:"dsn"`
	LogLevel string `json:"log_level"`
}

type CorsConfig struct {
	AllowedOrigins []string `json:"allowed_origins"`
}

// InitConfigFromFile initializes a new Env from configuration file.
func InitConfigFromFile(configFileName string) *Config {
	bz, err := ioutil.ReadFile(configFileName)
	if err != nil {
		panic(err)
	}

	var configOpts Config
	if err := json.Unmarshal(bz, &configOpts); err != nil {
		panic(err)
	}

	return &configOpts
}

func InitConfigFromSecret(secretName, secretRegion string) *Config {
	bzStr, err := svc.GetSecret(secretName, secretRegion)
	if err != nil {
		panic(err)
	}

	var configOpts Config
	if err := json.Unmarshal([]byte(bzStr), &configOpts); err != nil {
		panic(err)
	}

	return &configOpts
}
