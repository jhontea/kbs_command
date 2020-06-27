package command

import (
	"github.com/jhontea/kbs_command/receiver"
)

// FibonacciCommand :nodoc
type FibonacciCommand struct {
	device receiver.DeviceReceiverInterface
}

// NewFibonacciCommand :nodoc:
func NewFibonacciCommand(d receiver.DeviceReceiverInterface) Command {
	return &FibonacciCommand{
		device: d,
	}
}

// Execute :nodoc:
func (c *FibonacciCommand) Execute() {
	c.device.Fibonacci()
}
