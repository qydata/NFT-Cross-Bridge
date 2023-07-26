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
				EthChainID:              e.Config.SwapConfig.ETHChainID,
				PolygonChainID:          e.Config.SwapConfig.PolygonChainID,
				FantomChainID:           e.Config.SwapConfig.FantomChainID,
				BscChainID:              e.Config.SwapConfig.BSCChainID,
				BscErc721SwapAgent:      e.Config.SwapConfig.BSCErc721SwapAgent,
				BscErc1155SwapAgent:     e.Config.SwapConfig.BSCErc1155SwapAgent,
				CtChainID:              e.Config.SwapConfig.CtChainID,
				CtErc721SwapAgent:      e.Config.SwapConfig.CtErc721SwapAgent,
				CtErc1155SwapAgent:     e.Config.SwapConfig.CtErc1155SwapAgent,
				CooChainID:              e.Config.SwapConfig.CooChainID,
				CooErc721SwapAgent:      e.Config.SwapConfig.CooErc721SwapAgent,
				CooErc1155SwapAgent:     e.Config.SwapConfig.CooErc1155SwapAgent,
				EthErc721SwapAgent:      e.Config.SwapConfig.EthErc721SwapAgent,
				PolygonErc721SwapAgent:  e.Config.SwapConfig.PolygonErc721SwapAgent,
				FantomErc721SwapAgent:   e.Config.SwapConfig.FantomErc721SwapAgent,
				EthErc1155SwapAgent:     e.Config.SwapConfig.EthErc1155SwapAgent,
				PolygonErc1155SwapAgent: e.Config.SwapConfig.PolygonErc1155SwapAgent,
				FantomErc1155SwapAgent:  e.Config.SwapConfig.FantomErc1155SwapAgent,
			}
			return svc_info.NewGetInfoOK().WithPayload(&info)
		},
	}
}
