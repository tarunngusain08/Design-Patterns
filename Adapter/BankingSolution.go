package main

import (
	"fmt"
)

// LegacyBank is a legacy system that needs to be adapted.
type LegacyBank struct{}

// Withdraw method of LegacyBank system.
func (b *LegacyBank) Withdraw(amount float32) {
	fmt.Printf("Withdrawing amount %0.2f from LegacyBank system...\n", amount)
}

// NewBank is a new banking system that uses a different method for withdrawing money.
type NewBank struct{}

// PullOut method of NewBank system.
func (b *NewBank) PullOut(amount float32) {
	fmt.Printf("Pulling out amount %0.2f from NewBank system...\n", amount)
}

// BankAdapter is an adapter that converts the NewBank system's PullOut method into the Withdraw method that the client is expecting.
type BankAdapter struct {
	NewBank *NewBank
}

// Withdraw method of BankAdapter.
func (a *BankAdapter) Withdraw(amount float32) {
	a.NewBank.PullOut(amount)
}

func main() {
	// Create a LegacyBank object.
	legacyBank := &LegacyBank{}

	// Withdraw from the LegacyBank system.
	legacyBank.Withdraw(100.50)

	// Create a NewBank object.
	newBank := &NewBank{}

	// Create a BankAdapter that uses the NewBank system.
	adapter := &BankAdapter{NewBank: newBank}

	// Withdraw using the adapter (which calls the NewBank system's PullOut method).
	adapter.Withdraw(200.75)
}

/*
In this example, we have a LegacyBank system that has a Withdraw method, and a NewBank system that has a PullOut method. 
The client expects to use the Withdraw method, so we create an adapter called BankAdapter that converts the NewBank system's 
PullOut method into the Withdraw method. We then use the adapter to withdraw money from the NewBank system using the 
Withdraw method that the client is expecting.
*/
