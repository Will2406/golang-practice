package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct {
}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct {
}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Payment using BankPayment %d", bankAccount)
}

type BankPaymentAdapter struct {
	bankPayment BankPayment
	bankAccount int
}

func (bpa BankPaymentAdapter) Pay() {
	bpa.bankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	bank := &BankPaymentAdapter{
		bankPayment: BankPayment{},
		bankAccount: 123213,
	}
	ProcessPayment(bank)
}
