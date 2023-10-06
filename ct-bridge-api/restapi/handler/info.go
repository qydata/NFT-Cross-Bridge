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
				CtChainID:              e.Config.SwapConfig.CtChainID,
				CtErc20SwapAgent:      e.Config.SwapConfig.CtErc20SwapAgent,
				CtErc721SwapAgent:      e.Config.SwapConfig.CtErc721SwapAgent,
				CtErc1155SwapAgent:     e.Config.SwapConfig.CtErc1155SwapAgent,
				CooChainID:              e.Config.SwapConfig.CooChainID,
				CooErc20SwapAgent:      e.Config.SwapConfig.CooErc20SwapAgent,
				CooErc721SwapAgent:      e.Config.SwapConfig.CooErc721SwapAgent,
				CooErc1155SwapAgent:     e.Config.SwapConfig.CooErc1155SwapAgent,
			}
			return svc_info.NewGetInfoOK().WithPayload(&info)
		},
	}
}
