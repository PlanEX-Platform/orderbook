package pkg

import (
	"github.com/satori/go.uuid"
	"math/big"
	"time"
)

type Order struct {

	// true - ордер на продажу, false - ордер на покупку
	isSale bool

	// id пользователя
	userId uuid.UUID

	// объём
	volume big.Float

	// цена
	price big.Float

	// время
	time time.Time

	// статус сделки
	status OrderStatus
}

type OrderStatus int
const (
	OS_NEW OrderStatus = iota
	OS_OPEN
	OS_PARTIAL
	OS_FILLED
	OS_CANCELLED
)
