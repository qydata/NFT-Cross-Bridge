package handler

import (
	"net/http"

	"github.com/go-openapi/swag"

	"github.com/go-openapi/runtime/middleware"

	"github.com/qydata/ct-evm-compatible-bridge-api/models"
	"github.com/qydata/ct-evm-compatible-bridge-api/restapi/operations"
	"github.com/qydata/ct-evm-compatible-bridge-api/restapi/operations/erc_20_swap_pairs"
	"github.com/qydata/ct-evm-compatible-bridge-api/utils/env"
)

type GetERC20SwapPairsHandler struct {
	*env.Env
	H func(e *env.Env, params erc_20_swap_pairs.GetErc20SwapPairsParams) middleware.Responder
}

func (h GetERC20SwapPairsHandler) Serve(params erc_20_swap_pairs.GetErc20SwapPairsParams) middleware.Responder {
	responder := h.H(h.Env, params)
	return responder
}

// NewGetERC20SwapPairsHandler creates a new GetERC20SwapPairsHandler instance.
func NewGetERC20SwapPairsHandler(e *env.Env, api *operations.BscEvmCompatibleBridgeAPIAPI) GetERC20SwapPairsHandler {
	return GetERC20SwapPairsHandler{
		Env: e,
		H: func(e *env.Env, params erc_20_swap_pairs.GetErc20SwapPairsParams) middleware.Responder {
			total, pairList, err := e.ERC20SwapPairDao.GetSwapPairs(params)
			if err != nil {
				return erc_20_swap_pairs.NewGetErc20SwapPairsBadRequest().WithPayload(&models.Error{Code: http.StatusBadRequest, Message: swag.String(err.Error())})
			}
			res := models.Erc20SwapPairs{Total: total, Pairs: make([]*models.Erc20SwapPair, 0, len(pairList))}
			for _, t := range pairList {
				res.Pairs = append(res.Pairs, &models.Erc20SwapPair{
					Available:      t.Available,
					NAME:           t.NAME,
					SYMBOL:         t.SYMBOL,
					CreateTxHash:   t.CreateTxHash,
					CreatedAt:      t.CreatedAt.String(),
					DstChainID:     t.DstChainID,
					DstTokenAddr:   t.DstTokenAddr,
					RegisterTxHash: t.RegisterTxHash,
					Sponsor:        t.Sponsor,
					SrcChainID:     t.SrcChainID,
					SrcTokenAddr:   t.SrcTokenAddr,
					State:          string(t.State),
					UpdatedAt:      t.UpdatedAt.String(),
				})
			}

			return erc_20_swap_pairs.NewGetErc20SwapPairsOK().WithPayload(&res)
		},
	}
}
