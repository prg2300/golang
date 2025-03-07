package main

import "fmt"

//no inbuit enum (type/const)
//enumerated types

type OrderStatus int

const (
	Received  OrderStatus = iota //predeclare identifier which increments
	Confirmed             = "confirmeed"
	Prepared
	Delivered
)

func changeOrderStatus(status string) {
	fmt.Println("changing status", status)
}

func main() {
	changeOrderStatus(Confirmed)
}
