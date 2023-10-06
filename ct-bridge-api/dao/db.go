package dao

import (
	"github.com/pkg/errors"
	erc20dao "github.com/qydata/ct-evm-compatible-bridge-api/dao/erc20"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	erc1155dao "github.com/qydata/ct-evm-compatible-bridge-api/dao/erc1155"
	erc721dao "github.com/qydata/ct-evm-compatible-bridge-api/dao/erc721"
)

func NewDaoServices(dsn, logLevel string) (
	erc20dao.SwapPairDaoInterface,
	erc20dao.SwapDaoInterface,
	erc721dao.SwapPairDaoInterface,
	erc721dao.SwapDaoInterface,
	erc1155dao.SwapPairDaoInterface,
	erc1155dao.SwapDaoInterface,
	error,
) {
	mysqlConn := mysql.New(mysql.Config{
		DSN: dsn,
	})
	db, err := gorm.Open(mysqlConn, &gorm.Config{
		Logger: logger.Default.LogMode(dbLogLevel(logLevel)),
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, errors.Wrap(err, "[NewDaoServices]: failed to open db")
	}

	return erc20dao.NewSwapPairDao(db), erc20dao.NewSwapDao(db), erc721dao.NewSwapPairDao(db), erc721dao.NewSwapDao(db), erc1155dao.NewSwapPairDao(db), erc1155dao.NewSwapDao(db), nil
}

func dbLogLevel(level string) logger.LogLevel {
	switch level {
	case "SILENT":
		return logger.Silent
	case "ERROR":
		return logger.Error
	case "WARN":
		return logger.Warn
	case "INFO":
		return logger.Info
	}

	return logger.Warn
}
