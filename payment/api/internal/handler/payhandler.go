package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"myFlashSale/common/types"
	"myFlashSale/payment/api/internal/logic"
	"myFlashSale/payment/api/internal/svc"
)

func PayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PayReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPayLogic(r.Context(), svcCtx)
		resp, err := l.Pay(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
