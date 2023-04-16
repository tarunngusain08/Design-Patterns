package Facade

import "fmt"

type Order struct {
	item     string
	quantity int
}

type StockChecker struct {
	availableStock map[string]int
}

func (s *StockChecker) checkStock(item string, quantity int) bool {
	if stock, ok := s.availableStock[item]; ok {
		return stock >= quantity
	}
	return false
}

type OrderProcessor struct{}

func (p *OrderProcessor) processOrder(order Order) {
	fmt.Printf("Processing order for %d units of %s\n", order.quantity, order.item)
}

type PaymentProcessor struct{}

func (p *PaymentProcessor) processPayment(amount float64) bool {
	fmt.Printf("Processing payment for $%.2f\n", amount)
	return true
}

type ShippingProcessor struct{}

func (p *ShippingProcessor) shipOrder(order Order) {
	fmt.Printf("Shipping order for %d units of %s\n", order.quantity, order.item)
}

type OrderFacade struct {
	stockChecker      *StockChecker
	orderProcessor    *OrderProcessor
	paymentProcessor  *PaymentProcessor
	shippingProcessor *ShippingProcessor
}

func NewOrderFacade() *OrderFacade {
	stockChecker := &StockChecker{
		availableStock: map[string]int{
			"item1": 5,
			"item2": 10,
		},
	}
	return &OrderFacade{
		stockChecker:      stockChecker,
		orderProcessor:    &OrderProcessor{},
		paymentProcessor:  &PaymentProcessor{},
		shippingProcessor: &ShippingProcessor{},
	}
}

func (f *OrderFacade) PlaceOrder(order Order) error {
	if !f.stockChecker.checkStock(order.item, order.quantity) {
		return fmt.Errorf("not enough stock for %d units of %s", order.quantity, order.item)
	}
	f.orderProcessor.processOrder(order)
	if !f.paymentProcessor.processPayment(10.99) {
		return fmt.Errorf("payment failed")
	}
	f.shippingProcessor.shipOrder(order)
	return nil
}

func main() {
	facade := NewOrderFacade()

	order := Order{item: "item1", quantity: 2}
	err := facade.PlaceOrder(order)
	if err != nil {
		fmt.Println(err)
	}

	order = Order{item: "item2", quantity: 15}
	err = facade.PlaceOrder(order)
	if err != nil {
		fmt.Println(err)
	}
}

/*
In this example, we have an OrderFacade that provides a simplified interface for placing an order. 
The OrderFacade encapsulates several subsystems, such as the stock checker, order processor, 
payment processor, and shipping processor, and provides a single method PlaceOrder for placing an order.

The PlaceOrder method first checks if there is enough stock for the order by calling the checkStock 
method on the StockChecker. If there is enough stock, it then processes the order by calling the 
processOrder method on the OrderProcessor, processes the payment by calling the processPayment 
method on the PaymentProcessor, and ships the order by calling the shipOrder method on the ShippingProcessor.

The client code only needs to create an instance of the OrderFacade and call the PlaceOrder method
with the order details. The OrderFacade takes care of the rest, hiding the complexity of the underlying 
subsystems from the client.
*/
