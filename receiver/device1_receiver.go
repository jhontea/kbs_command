package receiver

import "fmt"

// Device1Receiver :nodoc:
type Device1Receiver struct {
}

// NewDevice1Receiver :nodoc:
func NewDevice1Receiver() DeviceReceiverInterface {
	return &Device1Receiver{}
}

// Sum :nodoc:
func (p *Device1Receiver) Sum() {
	fmt.Println("Sum ... ")
}

// Multiply :nodoc:
func (p *Device1Receiver) Multiply() {
	fmt.Println("Multiply ... ")
}

// Prime :nodoc:
func (p *Device1Receiver) Prime() {
	fmt.Println("Prime ... ")
}

// Fibonacci :nodoc:
func (p *Device1Receiver) Fibonacci() {
	fmt.Println("Fibonacci ... ")
}
