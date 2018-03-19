package pkg

import (
	"github.com/huandu/skiplist"
)

type Orderbook struct {
	orders skiplist.SkipList
}

func addOrder(orderbook *Orderbook, order Order)  {
	orderbook.orders.Set(order.price, order)
}

func (orderbook *Orderbook) getValue(key interface{}) (interface{}, bool)   {
	return orderbook.orders.GetValue(key)
}

func compare(orderbook *Orderbook, order Order) {
	// compare and change status
	firstOrderWithSamePrice, exists := orderbook.getValue(20)
	if exists {
		Order (firstOrderWithSamePrice)
	}
}
