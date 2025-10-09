package main

import (
	"fmt"
	"time"
)

type PaymentProcessor interface {
	ProcessPayment(amount float64) (string, error)
	Refund(transactionId string) error
	GetBalance() float64
}

type PayPalProcessor struct {
	ApiKey             string
	Sandbox            bool
	Balance            float64
	TransactionCounter int
}


type StripeProcessor struct {
    SecretKey   string
    Balance     float64
    Currency    string
}

func NewPaypalProcessor(apikey string, sanbox bool) PaymentProcessor {
	return &PayPalProcessor{
		ApiKey:  apikey,
		Sandbox: sanbox,
		Balance: 0.0,
	}
}


func (p *PayPalProcessor) ConnectToPaypal() string {
	if p.Sandbox {
		return "Connected to Paypal Sandbox"
	} else {
		return "Connected to PayPal Production"
	}
}

func (p *PayPalProcessor) GenerateTransactionId() string {
	p.TransactionCounter++
	return fmt.Sprintf("PayPal-%d", p.TransactionCounter)
}

func (p *PayPalProcessor) ProcessPayment(amount float64) (string, error){
	fmt.Printf("%s\n", p.ConnectToPaypal())

	if amount <=0 {
		return "", fmt.Errorf("invalid amount: %.2f", amount)
	}

	transactionId := p.GenerateTransactionId()
	p.Balance+=amount
	fmt.Printf("Processing $%.2f via PayPal. Transaction: %s\n", amount, transactionId)

	return transactionId, nil
}

func (p *PayPalProcessor) GetBalance() float64 {
    return p.Balance
}

func (p *PayPalProcessor) Refund(transactionId string) error{
	fmt.Printf("%s\n", p.ConnectToPaypal())

	transactionId = p.GenerateTransactionId()
	fmt.Printf("Refunding via Paypal: %s\n", transactionId)

	refundAmount := 50.0
	p.Balance-=refundAmount

    return nil

}


func NewStripeProcessor(secretKey string) PaymentProcessor {
    return &StripeProcessor{
        SecretKey: secretKey,
        Balance: 0.0,
        Currency: "USD",
    }
}

func (s *StripeProcessor) AuthenticateStripe() bool {
    return len(s.SecretKey) > 0
}

func (s *StripeProcessor) GetBalance() float64{
	if !s.AuthenticateStripe() {
        fmt.Printf("stripe authentication failed")
    }
	return s.Balance
}

func (s *StripeProcessor) ProcessPayment(amount float64) (string, error) {
    if !s.AuthenticateStripe() {
        return "", fmt.Errorf("stripe authentication failed")
    }
    
    if amount <= 0 {
        return "", fmt.Errorf("invalid amount: %.2f", amount)
    }
    
    transactionId := fmt.Sprintf("stripe-%d", time.Now().Unix())
    s.Balance += amount
    
    fmt.Printf("Processing $%.2f via Stripe. Transaction: %s\n", amount, transactionId)
    return transactionId, nil
}

func (s *StripeProcessor) Refund(transactionId string) error {
    if !s.AuthenticateStripe() {
        return fmt.Errorf("stripe authentication failed")
    }
    
    fmt.Printf("Refunding via Stripe: %s\n", transactionId)
    refundAmount := 50.0
    s.Balance -= refundAmount
    
    return nil
}

func ProcessOrder(processor PaymentProcessor, amount float64) {
	transactionId, err := processor.ProcessPayment(amount)
    if err != nil {
        fmt.Printf("Payment failed: %v\n", err)
        return
    }
    
    fmt.Printf("Payment successful! Transaction ID: %s\n", transactionId)
    fmt.Printf("Current balance: $%.2f\n", processor.GetBalance())
}

func main() {
	
    var payPal PaymentProcessor = NewPaypalProcessor("paypal_key_123", true)
    var stripe PaymentProcessor = NewStripeProcessor("stripe_secret_456")

    ProcessOrder(payPal, 100.0)
    ProcessOrder(stripe, 75.50)
    
    payPal.Refund("PayPal-1")
    fmt.Printf("Balance after refund: $%.2f\n", payPal.GetBalance())
}