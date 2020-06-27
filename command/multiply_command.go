package command

import (
	"github.com/jhontea/kbs_command/receiver"
)

// MultiplyCommand :nodoc
type MultiplyCommand struct {
	Device receiver.Device
}

// Execute :nodoc:
func (c *MultiplyCommand) Execute() {
	c.Device.Multiply()
}
