package main

import "fmt"

// order struct
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time //nanosec precision
// }

// func newOrder(id string, amount float32, status string) *order {
// 	//initial setup
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }
// func (o *order) changestatus(status string) {
// 	o.status = status

// }
// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main() {

	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// muOrder := newOrder("1", 30.50, "received")
	// fmt.Println(muOrder.amount)

	// myorder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }
	// myorder.createdAt = time.Now()
	// fmt.Println(myorder.status)

	// fmt.Println("Order struct", myorder)
}
