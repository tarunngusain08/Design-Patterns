package Strategy

import "fmt"

// PaymentMethod is the strategy interface.
type PaymentMethod interface {
	Pay(amount float64) error
}

// CreditCard is one of the concrete strategies.
type CreditCard struct{}

func (cc *CreditCard) Pay(amount float64) error {
	fmt.Printf("Paying %.2f using credit card.\n", amount)
	return nil
}

// PayPal is another concrete strategy.
type PayPal struct{}

func (pp *PayPal) Pay(amount float64) error {
	fmt.Printf("Paying %.2f using PayPal.\n", amount)
	return nil
}

// PaymentProcessor is the context that uses the strategy.
type PaymentProcessor struct {
	Method PaymentMethod
}

func (pp *PaymentProcessor) ProcessPayment(amount float64) error {
	return pp.Method.Pay(amount)
}

func main() {
	// Set up the payment processor with a credit card payment method.
	paymentProcessor := &PaymentProcessor{
		Method: &CreditCard{},
	}

	// Process a payment using the credit card payment method.
	err := paymentProcessor.ProcessPayment(25.00)
	if err != nil {
		fmt.Printf("Error processing payment: %s\n", err.Error())
	}

	// Switch to the PayPal payment method.
	paymentProcessor.Method = &PayPal{}

	// Process a payment using the PayPal payment method.
	err = paymentProcessor.ProcessPayment(50.00)
	if err != nil {
		fmt.Printf("Error processing payment: %s\n", err.Error())
	}
}
