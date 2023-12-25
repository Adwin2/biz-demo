package routes

import (
	"context"
	"testing"

	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/cart"
)

func TestRPCAddCart(t *testing.T) {
	rpc.InitClient()
	resp, err := rpc.CartClient.AddItem(context.TODO(), &cart.AddItemRequest{UserId: 1, Item: &cart.CartItem{ProductId: 6, Quantity: 2}})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
