package main

import "fmt"

type Counter struct {
	value int
}

func (c Counter) Value() int {
	return c.value
}

func (c *Counter) Increment() {
	c.value++
}

type Wallet struct {
	balance float64
}

func (w Wallet) Balance() float64 {
	return w.balance
}

func (w *Wallet) Deposit(amount float64) {
	w.balance += amount
}

func main() {
	c := Counter{}
	fmt.Println("counter:", c.Value())
	c.Increment()
	fmt.Println("counter:", c.Value())

	w := &Wallet{}
	w.Deposit(50)
	w.Deposit(25)
	fmt.Println("wallet:", w.Balance())
}
