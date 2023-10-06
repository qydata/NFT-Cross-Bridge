package model

import (
	"gorm.io/gorm"

	"github.com/qydata/ct-evm-compatible-bridge-core/model/block"
	"github.com/qydata/ct-evm-compatible-bridge-core/model/erc1155"
	"github.com/qydata/ct-evm-compatible-bridge-core/model/erc20"
	"github.com/qydata/ct-evm-compatible-bridge-core/model/erc721"
)

func InitTables(db *gorm.DB) {
	db.AutoMigrate(&block.Log{})
	db.AutoMigrate(&erc20.SwapPair{})
	db.AutoMigrate(&erc20.Swap{})
	db.AutoMigrate(&erc721.SwapPair{})
	db.AutoMigrate(&erc721.Swap{})
	db.AutoMigrate(&erc1155.SwapPair{})
	db.AutoMigrate(&erc1155.Swap{})
}
