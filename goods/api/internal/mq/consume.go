package mq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"myFlashSale/common/types"
	"myFlashSale/goods/api/internal/logic"
	"myFlashSale/goods/api/internal/svc"
)

func ConsumeCheckStock(svcCtx *svc.ServiceContext, ch *amqp.Channel, queueName string) {
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	go func() {
		for d := range msgs {
			var req types.StockReq
			err := json.Unmarshal(d.Body, &req)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}

			// 调用本地的CheckStock逻辑
			checkStockLogic := logic.NewCheckStockLogic(context.Background(), svcCtx)
			stockReq := &types.StockReq{}
			resp, err := checkStockLogic.CheckStock(stockReq)
			if err != nil {
				log.Printf("Error checking stock: %s", err)
				continue
			}
			log.Printf("Stock available : %d", resp.ExistStock)
		}
	}()
}
