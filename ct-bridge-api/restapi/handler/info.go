package handler

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/qydata/ct-evm-compatible-bridge-api/models"
	"github.com/qydata/ct-evm-compatible-bridge-api/restapi/operations"
	"github.com/qydata/ct-evm-compatible-bridge-api/restapi/operations/svc_info"
	"github.com/qydata/ct-evm-compatible-bridge-api/utils/env"
)

type GetInfosHandler struct {
	*env.Env
	H func(e *env.Env, params svc_info.GetInfoParams) middleware.Responder
}

func (h GetInfosHandler) Serve(params svc_info.GetInfoParams) middleware.Responder {
	responder := h.H(h.Env, params)
	return responder
}

func NewGetInfoHandler(e *env.Env, _ *operations.BscEvmCompatibleBridgeAPIAPI) GetInfosHandler {
	return GetInfosHandler{
		Env: e,
		H: func(e *env.Env, _ svc_info.GetInfoParams) middleware.Responder {
			info := models.ServiceInfo{
				CtChainID:           e.Config.SwapConfig.CtChainID,
				CtErc20SwapAgent:    e.Config.SwapConfig.CtErc20SwapAgent,
				CtErc721SwapAgent:   e.Config.SwapConfig.CtErc721SwapAgent,
				CtErc1155SwapAgent:  e.Config.SwapConfig.CtErc1155SwapAgent,
				CooChainID:          e.Config.SwapConfig.CooChainID,
				CooErc20SwapAgent:   e.Config.SwapConfig.CooErc20SwapAgent,
				CooErc721SwapAgent:  e.Config.SwapConfig.CooErc721SwapAgent,
				CooErc1155SwapAgent: e.Config.SwapConfig.CooErc1155SwapAgent,

				SparkchainChainID:                 e.Config.SwapConfig.SparkchainChainID,
				SparkchainErc20SwapAgent:          e.Config.SwapConfig.SparkchainErc20SwapAgent,
				SparkchainErc721SwapAgent:         e.Config.SwapConfig.SparkchainErc721SwapAgent,
				SparkchainErc1155SwapAgent:        e.Config.SwapConfig.SparkchainErc1155SwapAgent,
				ChainmakerChainID:                 e.Config.SwapConfig.ChainmakerChainID,
				ChainmakerErc20SwapAgent:          e.Config.SwapConfig.ChainmakerErc20SwapAgent,
				ChainmakerErc721SwapAgent:         e.Config.SwapConfig.ChainmakerErc721SwapAgent,
				ChainmakerErc1155SwapAgent:        e.Config.SwapConfig.ChainmakerErc1155SwapAgent,
				BsnChainID:                        e.Config.SwapConfig.BsnChainID,
				BsnErc20SwapAgent:                 e.Config.SwapConfig.BsnErc20SwapAgent,
				BsnErc721SwapAgent:                e.Config.SwapConfig.BsnErc721SwapAgent,
				BsnErc1155SwapAgent:               e.Config.SwapConfig.BsnErc1155SwapAgent,
				HyperchainChainID:                 e.Config.SwapConfig.HyperchainChainID,
				HyperchainErc20SwapAgent:          e.Config.SwapConfig.HyperchainErc20SwapAgent,
				HyperchainErc721SwapAgent:         e.Config.SwapConfig.HyperchainErc721SwapAgent,
				HyperchainErc1155SwapAgent:        e.Config.SwapConfig.HyperchainErc1155SwapAgent,
				BubichainChainID:                  e.Config.SwapConfig.BubichainChainID,
				BubichainErc20SwapAgent:           e.Config.SwapConfig.BubichainErc20SwapAgent,
				BubichainErc721SwapAgent:          e.Config.SwapConfig.BubichainErc721SwapAgent,
				BubichainErc1155SwapAgent:         e.Config.SwapConfig.BubichainErc1155SwapAgent,
				AntblockchainChainID:              e.Config.SwapConfig.AntblockchainChainID,
				AntblockchainErc20SwapAgent:       e.Config.SwapConfig.AntblockchainErc20SwapAgent,
				AntblockchainErc721SwapAgent:      e.Config.SwapConfig.AntblockchainErc721SwapAgent,
				AntblockchainErc1155SwapAgent:     e.Config.SwapConfig.AntblockchainErc1155SwapAgent,
				TencentblockchainChainID:          e.Config.SwapConfig.TencentblockchainChainID,
				TencentblockchainErc20SwapAgent:   e.Config.SwapConfig.TencentblockchainErc20SwapAgent,
				TencentblockchainErc721SwapAgent:  e.Config.SwapConfig.TencentblockchainErc721SwapAgent,
				TencentblockchainErc1155SwapAgent: e.Config.SwapConfig.TencentblockchainErc1155SwapAgent,
				JdblockchainChainID:               e.Config.SwapConfig.JdblockchainChainID,
				JdblockchainErc20SwapAgent:        e.Config.SwapConfig.JdblockchainErc20SwapAgent,
				JdblockchainErc721SwapAgent:       e.Config.SwapConfig.JdblockchainErc721SwapAgent,
				JdblockchainErc1155SwapAgent:      e.Config.SwapConfig.JdblockchainErc1155SwapAgent,
				XuperchainChainID:                 e.Config.SwapConfig.XuperchainChainID,
				XuperchainErc20SwapAgent:          e.Config.SwapConfig.XuperchainErc20SwapAgent,
				XuperchainErc721SwapAgent:         e.Config.SwapConfig.XuperchainErc721SwapAgent,
				XuperchainErc1155SwapAgent:        e.Config.SwapConfig.XuperchainErc1155SwapAgent,
				NeoChainID:                        e.Config.SwapConfig.NeoChainID,
				NeoErc20SwapAgent:                 e.Config.SwapConfig.NeoErc20SwapAgent,
				NeoErc721SwapAgent:                e.Config.SwapConfig.NeoErc721SwapAgent,
				NeoErc1155SwapAgent:               e.Config.SwapConfig.NeoErc1155SwapAgent,
				QtumChainID:                       e.Config.SwapConfig.QtumChainID,
				QtumErc20SwapAgent:                e.Config.SwapConfig.QtumErc20SwapAgent,
				QtumErc721SwapAgent:               e.Config.SwapConfig.QtumErc721SwapAgent,
				QtumErc1155SwapAgent:              e.Config.SwapConfig.QtumErc1155SwapAgent,
				VechainChainID:                    e.Config.SwapConfig.VechainChainID,
				VechainErc20SwapAgent:             e.Config.SwapConfig.VechainErc20SwapAgent,
				VechainErc721SwapAgent:            e.Config.SwapConfig.VechainErc721SwapAgent,
				VechainErc1155SwapAgent:           e.Config.SwapConfig.VechainErc1155SwapAgent,
				GxchainChainID:                    e.Config.SwapConfig.GxchainChainID,
				GxchainErc20SwapAgent:             e.Config.SwapConfig.GxchainErc20SwapAgent,
				GxchainErc721SwapAgent:            e.Config.SwapConfig.GxchainErc721SwapAgent,
				GxchainErc1155SwapAgent:           e.Config.SwapConfig.GxchainErc1155SwapAgent,
				OraclechainChainID:                e.Config.SwapConfig.OraclechainChainID,
				OraclechainErc20SwapAgent:         e.Config.SwapConfig.OraclechainErc20SwapAgent,
				OraclechainErc721SwapAgent:        e.Config.SwapConfig.OraclechainErc721SwapAgent,
				OraclechainErc1155SwapAgent:       e.Config.SwapConfig.OraclechainErc1155SwapAgent,
				OntologyChainID:                   e.Config.SwapConfig.OntologyChainID,
				OntologyErc20SwapAgent:            e.Config.SwapConfig.OntologyErc20SwapAgent,
				OntologyErc721SwapAgent:           e.Config.SwapConfig.OntologyErc721SwapAgent,
				OntologyErc1155SwapAgent:          e.Config.SwapConfig.OntologyErc1155SwapAgent,
				BytomChainID:                      e.Config.SwapConfig.BytomChainID,
				BytomErc20SwapAgent:               e.Config.SwapConfig.BytomErc20SwapAgent,
				BytomErc721SwapAgent:              e.Config.SwapConfig.BytomErc721SwapAgent,
				BytomErc1155SwapAgent:             e.Config.SwapConfig.BytomErc1155SwapAgent,
				DarwiniaChainID:                   e.Config.SwapConfig.DarwiniaChainID,
				DarwiniaErc20SwapAgent:            e.Config.SwapConfig.DarwiniaErc20SwapAgent,
				DarwiniaErc721SwapAgent:           e.Config.SwapConfig.DarwiniaErc721SwapAgent,
				DarwiniaErc1155SwapAgent:          e.Config.SwapConfig.DarwiniaErc1155SwapAgent,
				HecoChainID:                       e.Config.SwapConfig.HecoChainID,
				HecoErc20SwapAgent:                e.Config.SwapConfig.HecoErc20SwapAgent,
				HecoErc721SwapAgent:               e.Config.SwapConfig.HecoErc721SwapAgent,
				HecoErc1155SwapAgent:              e.Config.SwapConfig.HecoErc1155SwapAgent,
				PolkadotChainID:                   e.Config.SwapConfig.PolkadotChainID,
				PolkadotErc20SwapAgent:            e.Config.SwapConfig.PolkadotErc20SwapAgent,
				PolkadotErc721SwapAgent:           e.Config.SwapConfig.PolkadotErc721SwapAgent,
				PolkadotErc1155SwapAgent:          e.Config.SwapConfig.PolkadotErc1155SwapAgent,
				ConfluxChainID:                    e.Config.SwapConfig.ConfluxChainID,
				ConfluxErc20SwapAgent:             e.Config.SwapConfig.ConfluxErc20SwapAgent,
				ConfluxErc721SwapAgent:            e.Config.SwapConfig.ConfluxErc721SwapAgent,
				ConfluxErc1155SwapAgent:           e.Config.SwapConfig.ConfluxErc1155SwapAgent,
				BitsharesChainID:                  e.Config.SwapConfig.BitsharesChainID,
				BitsharesErc20SwapAgent:           e.Config.SwapConfig.BitsharesErc20SwapAgent,
				BitsharesErc721SwapAgent:          e.Config.SwapConfig.BitsharesErc721SwapAgent,
				BitsharesErc1155SwapAgent:         e.Config.SwapConfig.BitsharesErc1155SwapAgent,
				YoyowChainID:                      e.Config.SwapConfig.YoyowChainID,
				YoyowErc20SwapAgent:               e.Config.SwapConfig.YoyowErc20SwapAgent,
				YoyowErc721SwapAgent:              e.Config.SwapConfig.YoyowErc721SwapAgent,
				YoyowErc1155SwapAgent:             e.Config.SwapConfig.YoyowErc1155SwapAgent,
				UltrainChainID:                    e.Config.SwapConfig.UltrainChainID,
				UltrainErc20SwapAgent:             e.Config.SwapConfig.UltrainErc20SwapAgent,
				UltrainErc721SwapAgent:            e.Config.SwapConfig.UltrainErc721SwapAgent,
				UltrainErc1155SwapAgent:           e.Config.SwapConfig.UltrainErc1155SwapAgent,
				CocosbcxChainID:                   e.Config.SwapConfig.CocosbcxChainID,
				CocosbcxErc20SwapAgent:            e.Config.SwapConfig.CocosbcxErc20SwapAgent,
				CocosbcxErc721SwapAgent:           e.Config.SwapConfig.CocosbcxErc721SwapAgent,
				CocosbcxErc1155SwapAgent:          e.Config.SwapConfig.CocosbcxErc1155SwapAgent,
				AchainChainID:                     e.Config.SwapConfig.AchainChainID,
				AchainErc20SwapAgent:              e.Config.SwapConfig.AchainErc20SwapAgent,
				AchainErc721SwapAgent:             e.Config.SwapConfig.AchainErc721SwapAgent,
				AchainErc1155SwapAgent:            e.Config.SwapConfig.AchainErc1155SwapAgent,
			}
			return svc_info.NewGetInfoOK().WithPayload(&info)
		},
	}
}
