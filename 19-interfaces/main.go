package main

import "fmt"

/*
An Interface defines behaviour, not data
Unlike Java, where each class has to explicitly implement an interface
Go does NOT require this.
Go Interfaces Are IMPLICIT
If type has required methods: it automatically satisfies interface.
*/

type PaymentProcessor interface {
	Pay(accountNumber string, amount float32)
}

type Razorpay struct{}

func (r Razorpay) Pay(accountNumber string, amount float32) {
	fmt.Printf("Paying via Razorpay --> Account Number %v | Amount %v\n", accountNumber, amount)
}

type Stripe struct{}

func (s Stripe) Pay(accountNumber string, amount float32) {
	fmt.Printf("Paying via Stripe --> Account Number %v | Amount %v\n", accountNumber, amount)
}

type Paypal struct{}

func (s Paypal) Pay(accountNumber string, amount float32) {
	fmt.Printf("Paying via Paypal --> Account Number %v | Amount %v\n", accountNumber, amount)
}

type Payment struct {
	pgw           PaymentProcessor
	accountNumber string
	amount        float32
}

func NewPayment(pgw PaymentProcessor, accountNumber string, amount float32) *Payment {
	return &Payment{
		pgw:           pgw,
		accountNumber: accountNumber,
		amount:        amount,
	}
}

func (p Payment) makePayment() {
	p.pgw.Pay(p.accountNumber, p.amount)
}

func main() {
	razorpayGW := Razorpay{}
	stripeGW := Stripe{}
	paypalGW := Paypal{}
	accountNumber := "abc123"
	amount := float32(10.99)
	razorpayPayment := NewPayment(razorpayGW, accountNumber, amount)
	stripePayment := NewPayment(stripeGW, accountNumber, amount)
	paypalPayment := NewPayment(paypalGW, accountNumber, amount)

	razorpayPayment.makePayment()
	stripePayment.makePayment()
	paypalPayment.makePayment()

}
