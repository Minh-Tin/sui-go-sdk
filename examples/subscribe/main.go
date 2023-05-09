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
	var ctx = context.Background()
	var cli = sui.NewSuiWebsocketClient(constant.WssBvMainnetEndpoint)

	receiveMsgCh := make(chan models.SuiEventResponse, 10)
	err := cli.SubscribeEvent(ctx, models.SuiXSubscribeEventsRequest{
		//message := `{"jsonrpc":"2.0", "id": "sub_` + uuid.New().String() + `", "method": "suix_subscribeEvent", "params": [{"MoveEventType": "0xd5dd28cc24009752905689b2ba2bf90bfc8de4549b9123f93519bb8ba9bf9981::marketplace::ListingEvent"},{"MoveEventType": "0xd5dd28cc24009752905689b2ba2bf90bfc8de4549b9123f93519bb8ba9bf9981::marketplace::ChangePriceEvent"}]}`
		SuiEventFilter: map[string]interface{}{
			"Or": []interface{}{
				map[string]interface{}{
					//"EventType": "MoveEvent",
					//"Package":       "0xd5dd28cc24009752905689b2ba2bf90bfc8de4549b9123f93519bb8ba9bf9981",
					"MoveEventType": "0xd5dd28cc24009752905689b2ba2bf90bfc8de4549b9123f93519bb8ba9bf9981::marketplace::ListingEvent",
					//"MoveModule": "marketplace",
				},
				map[string]interface{}{
					//"EventType": "MoveEvent",
					//"Package":       "0xd5dd28cc24009752905689b2ba2bf90bfc8de4549b9123f93519bb8ba9bf9981",
					"MoveEventType": "0xd5dd28cc24009752905689b2ba2bf90bfc8de4549b9123f93519bb8ba9bf9981::marketplace::ChangePriceEvent",
					//"MoveModule": "marketplace",
				},
			},
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
	var ctx = context.Background()
	var cli = sui.NewSuiWebsocketClient(constant.WssBvTestnetEndpoint)

	receiveMsgCh := make(chan models.SuiEffects, 10)
	err := cli.SubscribeTransaction(ctx, models.SuiXSubscribeTransactionsRequest{
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
