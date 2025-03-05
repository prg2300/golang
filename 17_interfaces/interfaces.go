package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

//open close principle
func (p payment) makePayment(amount float32) {
	//	razorpayPaymentGw := razorpay{}       //another method for stripe pay
	//stripePayment := stripe{}
	//razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe", amount)
// }
type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose")
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

//make code scalable and organised

func (p paypal) refund(amount float32, account string) {

}

func main() {
	// stripePaymentGw := stripe{} //payment ke andr gateway pass krna
	//razorpayPaymentGw := razorpay{}
	//
	paypalGw := paypal{}
	// fakeGw := fakepayment{}
	newpayment := payment{
		gateway: paypalGw,
	}
	newpayment.makePayment(100)
}
