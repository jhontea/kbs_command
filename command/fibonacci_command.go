package command

import (
	"github.com/jhontea/kbs_command/receiver"
)

// FibonacciCommand :nodoc
type FibonacciCommand struct {
	Device receiver.Device
}

// Execute :nodoc:
func (c *FibonacciCommand) Execute() {
	c.Device.Fibonacci()
}
