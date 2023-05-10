package main

import (
	"context"
	"fmt"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
)

var ctx = context.Background()
var cli = sui.NewSuiClient(constant.BvMainnetEndpoint)

func main() {
	SuiGetObject()
	//SuiXGetOwnedObjects()
	return
	SuiMultiGetObjects()
	SuiXGetDynamicField()
	SuiTryGetPastObject()
	SuiGetLoadedChildObjects()
}

func SuiGetObject() {
	rsp, err := cli.SuiGetObject(ctx, models.SuiGetObjectRequest{
		ObjectId: "0x4de3bf06892a322795846c2d3d364a8032d8ba2a2d54a1422d749dc0fa272386",
		Options: models.SuiObjectDataOptions{
			ShowContent:             true,
			ShowDisplay:             true,
			ShowType:                true,
			ShowBcs:                 true,
			ShowOwner:               true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)

}

func SuiXGetOwnedObjects() {
	suiObjectResponseQuery := models.SuiObjectResponseQuery{
		Filter: models.ObjectFilterByPackage{
			Package: "0xee496a0cc04d06a345982ba6697c90c619020de9e274408c7819f787ff66e1a1",
		},
		Options: models.SuiObjectDataOptions{
			ShowType:                true,
			ShowContent:             true,
			ShowBcs:                 true,
			ShowOwner:               true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
			ShowDisplay:             true,
		},
	}

	rsp, err := cli.SuiXGetOwnedObjects(ctx, models.SuiXGetOwnedObjectsRequest{
		Address: "0xb7f98d327f19f674347e1e40641408253142d6e7e5093a7c96eda8cdfd7d9bb5",
		Query:   suiObjectResponseQuery,
		Limit:   5,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiMultiGetObjects() {
	rsp, err := cli.SuiMultiGetObjects(ctx, models.SuiMultiGetObjectsRequest{
		ObjectIds: []string{"0xeeb964d1e640219c8ddb791cc8548f3242a3392b143ff47484a3753291cad898"},
		Options: models.SuiObjectDataOptions{
			ShowContent:             true,
			ShowDisplay:             true,
			ShowType:                true,
			ShowBcs:                 true,
			ShowOwner:               true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, object := range rsp {
		utils.PrettyPrint(*object)
	}
}

func SuiXGetDynamicField() {
	rsp, err := cli.SuiXGetDynamicField(ctx, models.SuiXGetDynamicFieldRequest{
		ObjectId: "0xeeb964d1e640219c8ddb791cc8548f3242a3392b143ff47484a3753291cad898",
		Limit:    5,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiTryGetPastObject() {
	rsp, err := cli.SuiTryGetPastObject(ctx, models.SuiTryGetPastObjectRequest{
		ObjectId: "0xeeb964d1e640219c8ddb791cc8548f3242a3392b143ff47484a3753291cad898",
		Version:  19636,
		Options: models.SuiObjectDataOptions{
			ShowContent:             true,
			ShowDisplay:             true,
			ShowType:                true,
			ShowBcs:                 true,
			ShowOwner:               true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiGetLoadedChildObjects() {
	rsp, err := cli.SuiGetLoadedChildObjects(ctx, models.SuiGetLoadedChildObjectsRequest{
		Digest: "DDvbPE1Ty138BEsu1238rRkpx4DMMDKbCJktt4H1cG6T",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}
