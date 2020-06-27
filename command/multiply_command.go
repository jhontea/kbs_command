package command

import (
	"github.com/jhontea/kbs_command/receiver"
)

// MultiplyCommand :nodoc
type MultiplyCommand struct {
	device receiver.DeviceReceiverInterface
}

// NewMultiplyCommand :nodoc:
func NewMultiplyCommand(d receiver.DeviceReceiverInterface) Command {
	return &MultiplyCommand{
		device: d,
	}
}

// Execute :nodoc:
func (c *MultiplyCommand) Execute() {
	c.device.Multiply()
}
