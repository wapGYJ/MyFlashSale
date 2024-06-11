package handler

import (
	"myFlashSale/common/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"myFlashSale/goods/api/internal/logic"
	"myFlashSale/goods/api/internal/svc"
)

func GetGoodsFromCacheHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetGoodsFromCacheReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetGoodsFromCacheLogic(r.Context(), svcCtx)
		resp, err := l.GetGoodsFromCache(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
