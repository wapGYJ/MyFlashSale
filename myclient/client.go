package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL              = "http://localhost:80" // 基础URL，根据实际情况修改
	registerURL          = baseURL + "/api/users/register"
	loginURL             = baseURL + "/api/users/login"
	checkStockURL        = baseURL + "/goods/checkstock"
	lockStockURL         = baseURL + "/goods/lockstock"
	getGoodsFromCacheURL = baseURL + "/goods/getgoodsfromcache"
	createOrderURL       = baseURL + "/order/createorder"
	createPaymentURL     = baseURL + "/payment/createpayment"
	payURL               = baseURL + "/payment/pay"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserId int64  `json:"id"`
	Token  string `json:"token"`
}

type StockResp struct {
	ExistStock int64 `json:"existStock"`
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

type CreateOrderReq struct {
	GoodsId int64 `json:"goodsId"`
	UserId  int64 `json:"userId"`
}

type CreateOrderResp struct {
	OrderId int64  `json:"orderId"`
	Message string `json:"message"`
}

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

func main() {
	//注册
	user := UserRequest{
		Username: "testuser42",
		Password: "testpassword",
	}
	userResponse, err := registerUser(user)
	if err != nil {
		fmt.Println("用户注册失败:", err)
		return
	}
	fmt.Printf("用户成功注册User ID: %d\n", userResponse.UserId)

	//  登录
	loginResponse, err := loginUser(user)
	if err != nil {
		fmt.Println("用户登陆失败:", err)
		return
	}
	token := loginResponse.Token
	fmt.Printf("用户登陆成功Token: %s\n", token)

	//  检查库存
	stockResp, err := checkStock()
	if err != nil {
		fmt.Println("无法查询库存:", err)
		return
	}
	if stockResp.ExistStock <= 0 {
		fmt.Println("手慢了，没有库存了")
		return
	}
	fmt.Printf("库存余量: %d\n", stockResp.ExistStock)

	//从缓存中获取商品
	GetGoodsFromCache, err := getGoodsFromCache()
	if err != nil {
		fmt.Println("获取缓存错误:", err)
		return
	}
	fmt.Printf("成功获取商品缓存: ID: %d, Name: %s, Price: %d\n", GetGoodsFromCache.Id, GetGoodsFromCache.Name, GetGoodsFromCache.Price)

	// 锁定商品，删除商品
	GoodsId := GetGoodsFromCache.Id
	lockStockResp, err := lockStock(GoodsId)
	if err != nil {
		fmt.Println("锁定商品错误:", err)
		return
	}
	fmt.Printf("成功锁定商品, Goods ID: %d, Quantity: %d\n", lockStockResp.LockId, lockStockResp.Quantity)

	// 创建订单
	orderResp, err := createOrder(GoodsId, userResponse.UserId)
	if err != nil {
		fmt.Println("创建订单出错:", err)
		return
	}
	fmt.Printf("订单生成成功, Order ID: %d, Message: %s\n", orderResp.OrderId, orderResp.Message)

	// 支付
	paymentResp, err := createPayment(orderResp.OrderId, userResponse.UserId, GoodsId)
	if err != nil {
		fmt.Println("创建支付失败:", err)
		return
	}
	fmt.Printf("成功创建支付, Payment ID: %d, Status: %d\n", paymentResp.PaymentId, paymentResp.Status)

	// 实际支付
	payResp, err := pay(userResponse.UserId, paymentResp.PaymentId, GoodsId)
	if err != nil {
		fmt.Println("支付出错:", err)
		return
	}
	fmt.Printf("成功支付, Message: %s\n", payResp.Message)
}

func registerUser(user UserRequest) (UserResponse, error) {
	body, err := json.Marshal(user)
	if err != nil {
		return UserResponse{}, err
	}
	resp, err := http.Post(registerURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return UserResponse{}, err
	}
	defer resp.Body.Close()

	var userResponse UserResponse
	err = json.NewDecoder(resp.Body).Decode(&userResponse)
	if err != nil {
		return UserResponse{}, err
	}

	return userResponse, nil
}

func loginUser(user UserRequest) (UserResponse, error) {
	body, err := json.Marshal(user)
	if err != nil {
		return UserResponse{}, err
	}

	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return UserResponse{}, err
	}
	defer resp.Body.Close()

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		return UserResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read and log the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserResponse{}, err
	}
	fmt.Println("Response Body:", string(respBody))

	var userResponse UserResponse
	err = json.Unmarshal(respBody, &userResponse)
	if err != nil {
		return UserResponse{}, err
	}

	return userResponse, nil

}

func checkStock() (StockResp, error) {
	resp, err := http.Get(checkStockURL)
	if err != nil {
		return StockResp{}, err
	}
	defer resp.Body.Close()

	var stockResp StockResp
	err = json.NewDecoder(resp.Body).Decode(&stockResp)
	if err != nil {
		return StockResp{}, err
	}

	return stockResp, nil
}

func getGoodsFromCache() (GetGoodsFromCacheResp, error) {
	resp, err := http.Get(getGoodsFromCacheURL)
	if err != nil {
		return GetGoodsFromCacheResp{}, err
	}
	defer resp.Body.Close()

	var goodsResp GetGoodsFromCacheResp
	err = json.NewDecoder(resp.Body).Decode(&goodsResp)
	if err != nil {
		return GetGoodsFromCacheResp{}, err
	}

	return goodsResp, nil
}

func lockStock(goodsId int64) (LockStockResponse, error) {

	lockStockReq := LockStockRequest{GoodsId: goodsId}
	body, err := json.Marshal(lockStockReq)
	if err != nil {
		return LockStockResponse{}, err
	}
	resp, err := http.Post(lockStockURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return LockStockResponse{}, err
	}
	defer resp.Body.Close()
	var lockStockResp LockStockResponse
	err = json.NewDecoder(resp.Body).Decode(&lockStockResp)
	if err != nil {
		return LockStockResponse{}, err
	}
	return lockStockResp, nil
}

func createOrder(goodsId int64, userId int64) (CreateOrderResp, error) {
	orderReq := CreateOrderReq{
		GoodsId: goodsId,
		UserId:  userId,
	}
	body, err := json.Marshal(orderReq)
	if err != nil {
		return CreateOrderResp{}, err
	}
	resp, err := http.Post(createOrderURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return CreateOrderResp{}, err
	}
	defer resp.Body.Close()

	var orderResp CreateOrderResp
	err = json.NewDecoder(resp.Body).Decode(&orderResp)
	if err != nil {
		return CreateOrderResp{}, err
	}

	return orderResp, nil
}

func createPayment(orderId int64, userId int64, goodsId int64) (CreatePaymentResp, error) {
	paymentReq := CreatePaymentReq{
		OrderId: orderId,
		UserId:  userId,
		GoodsId: goodsId,
	}
	body, err := json.Marshal(paymentReq)
	if err != nil {
		return CreatePaymentResp{}, err
	}
	resp, err := http.Post(createPaymentURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return CreatePaymentResp{}, err
	}
	defer resp.Body.Close()

	var paymentResp CreatePaymentResp
	err = json.NewDecoder(resp.Body).Decode(&paymentResp)
	if err != nil {
		return CreatePaymentResp{}, err
	}

	return paymentResp, nil
}

func pay(userId int64, paymentId int64, goodsId int64) (PayResp, error) {
	payReq := PayReq{
		UserId:    userId,
		PaymentId: paymentId,
		GoodsId:   goodsId,
	}
	body, err := json.Marshal(payReq)
	if err != nil {
		return PayResp{}, err
	}
	resp, err := http.Post(payURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return PayResp{}, err
	}
	defer resp.Body.Close()

	var payResp PayResp
	err = json.NewDecoder(resp.Body).Decode(&payResp)
	if err != nil {
		return PayResp{}, err
	}

	return payResp, nil
}
