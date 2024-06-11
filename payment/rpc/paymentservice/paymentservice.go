// Code generated by goctl. DO NOT EDIT.
// Source: payment.proto

package paymentservice

import (
	"context"

	"myFlashSale/payment/rpc/pb/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreatePaymentRequest  = pb.CreatePaymentRequest
	CreatePaymentResponse = pb.CreatePaymentResponse
	PayReq                = pb.PayReq
	PayResp               = pb.PayResp

	PaymentService interface {
		CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
		Pay(ctx context.Context, in *PayReq, opts ...grpc.CallOption) (*PayResp, error)
	}

	defaultPaymentService struct {
		cli zrpc.Client
	}
)

func NewPaymentService(cli zrpc.Client) PaymentService {
	return &defaultPaymentService{
		cli: cli,
	}
}

func (m *defaultPaymentService) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	client := pb.NewPaymentServiceClient(m.cli.Conn())
	return client.CreatePayment(ctx, in, opts...)
}

func (m *defaultPaymentService) Pay(ctx context.Context, in *PayReq, opts ...grpc.CallOption) (*PayResp, error) {
	client := pb.NewPaymentServiceClient(m.cli.Conn())
	return client.Pay(ctx, in, opts...)
}
