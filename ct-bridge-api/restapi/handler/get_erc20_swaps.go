package handler

import (
	"net/http"

	"github.com/go-openapi/swag"

	"github.com/go-openapi/runtime/middleware"

	"github.com/qydata/ct-evm-compatible-bridge-api/models"
	"github.com/qydata/ct-evm-compatible-bridge-api/restapi/operations"
	"github.com/qydata/ct-evm-compatible-bridge-api/restapi/operations/erc_20_swaps"
	"github.com/qydata/ct-evm-compatible-bridge-api/utils/env"
)

type GetERC20SwapsHandler struct {
	*env.Env
	H func(e *env.Env, params erc_20_swaps.GetErc20SwapsParams) middleware.Responder
}

func (h GetERC20SwapsHandler) Serve(params erc_20_swaps.GetErc20SwapsParams) middleware.Responder {
	responder := h.H(h.Env, params)
	return responder
}

// NewGetERC20SwapsHandler creates a new GetERC20SwapsHandler instance.
func NewGetERC20SwapsHandler(e *env.Env, api *operations.BscEvmCompatibleBridgeAPIAPI) GetERC20SwapsHandler {
	return GetERC20SwapsHandler{
		Env: e,
		H: func(e *env.Env, params erc_20_swaps.GetErc20SwapsParams) middleware.Responder {

			total, SwapsList, err := e.ERC20SwapDao.GetSwaps(params)
			if err != nil {
				return erc_20_swaps.NewGetErc20SwapsBadRequest().WithPayload(&models.Error{Code: http.StatusBadRequest, Message: swag.String(err.Error())})
			}
			res := models.Erc20Swaps{Total: total, Erc20Swaps: make([]*models.Erc20Swap, 0, len(SwapsList))}

			//log.Printf("SwapsList", res)
			for _, s := range SwapsList {

				res.Erc20Swaps = append(res.Erc20Swaps, &models.Erc20Swap{
					CreatedAt:     s.CreatedAt.String(),
					DstChainID:    s.DstChainID,
					DstTokenAddr:  s.DstTokenAddr,
					FillTxHash:    s.FillTxHash,
					Recipient:     s.Recipient,
					RequestTxHash: s.RequestTxHash,
					Sender:        s.Sender,
					SrcChainID:    s.SrcChainID,
					SrcTokenAddr:  s.SrcTokenAddr,
					State:         string(s.State),
					SwapDirection: string(s.SwapDirection),
					Id:            s.ID,
					Amount:        s.Amount,
					UpdatedAt:     s.UpdatedAt.String(),
				})
			}

			return erc_20_swaps.NewGetErc20SwapsOK().WithPayload(&res)
		},
	}
}
