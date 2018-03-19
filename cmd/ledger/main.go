package main

import (
	"log"
	"planex/ledger/pkg"
	"github.com/satori/go.uuid"
	"math/big"
	"time"
	"github.com/huandu/skiplist"
)

func main() {
	log.Printf("hello, world!")
	order := pkg.Order{true, uuid.FromStringOrNil(""), big.Float{},
	big.Float{}, time.Now(), pkg.OS_NEW}

	log.Println(order)
	pkg.Orderbook{skiplist.New(skiplist.Int)}
}
