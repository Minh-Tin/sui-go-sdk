package main

import (
	"context"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
)

func main() {
	go SubscribeEvent()
	//go SubscribeTransaction()
	select {}
}

func SubscribeEvent() {
	var ctx, cancel = context.WithCancel(context.Background())
	var cli = sui.NewSuiWebsocketClient("ws://192.168.1.40:9990")

	receiveMsgCh := make(chan models.SuiEventResponse, 10)
	err := cli.SubscribeEvent(ctx, cancel, models.SuiXSubscribeEventsRequest{
		SuiEventFilter: map[string]interface{}{
			"All": []string{},
			//"Or": []interface{}{
			//	map[string]interface{}{
			//		"MoveEventType": "0xd5dd28cc24009752905689b2ba2bf90bfc8de4549b9123f93519bb8ba9bf9981::marketplace::ListingEvent",
			//	},
			//	map[string]interface{}{
			//		"MoveEventType": "0xd5dd28cc24009752905689b2ba2bf90bfc8de4549b9123f93519bb8ba9bf9981::marketplace::ChangePriceEvent",
			//	},
			//},
		},
	}, receiveMsgCh)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case msg := <-receiveMsgCh:
			utils.PrettyPrint(msg)
		case <-ctx.Done():
			return
		}
	}
}

func SubscribeTransaction() {
	var ctx, cancel = context.WithCancel(context.Background())
	var cli = sui.NewSuiWebsocketClient(constant.WssBvTestnetEndpoint)

	receiveMsgCh := make(chan models.SuiEffects, 10)
	err := cli.SubscribeTransaction(ctx, cancel, models.SuiXSubscribeTransactionsRequest{
		TransactionFilter: models.TransactionFilterByFromAddress{
			FromAddress: "0x0000000000000000000000000000000000000000000000000000000000000000",
		},
	}, receiveMsgCh)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case msg := <-receiveMsgCh:
			utils.PrettyPrint(msg)
		case <-ctx.Done():
			return
		}
	}
}
