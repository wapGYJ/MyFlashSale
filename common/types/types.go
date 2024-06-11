package types

// goods
type GetGoodsFromCacheReq struct {
}

type GetGoodsFromCacheResp struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type LockStockRequest struct {
	GoodsId int64 `json:"goodsId"`
}

type LockStockResponse struct {
	LockId   int64 `json:"lockId"`
	Quantity int64 `json:"quantity"`
}

type StockReq struct {
}

type StockResp struct {
	ExistStock int64 `json:"existStock"`
}

// user

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserId int64  `json:"id"`
	Token  string `json:"token"`
}

// order
type CreateOrderReq struct {
	GoodsId int64 `json:"goodsId"`
	UserId  int64 `json:"userId"`
}

type CreateOrderResp struct {
	OrderId int64  `json:"orderId"`
	Message string `json:"message"`
}

// payment
type CreatePaymentReq struct {
	OrderId int64 `json:"orderId"`
	UserId  int64 `json:"userId"`
	GoodsId int64 `json:"goodsId"`
}

type CreatePaymentResp struct {
	PaymentId int64 `json:"paymentId"`
	Status    int64 `json:"status"`
}

type PayReq struct {
	UserId    int64 `json:"userId"`
	PaymentId int64 `json:"paymentId"`
	GoodsId   int64 `json:"goodsId"`
}

type PayResp struct {
	Message string `json:"message"`
}
