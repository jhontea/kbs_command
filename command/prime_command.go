package command

import (
	"github.com/jhontea/kbs_command/receiver"
)

// PrimeCommand :nodoc
type PrimeCommand struct {
	device receiver.DeviceReceiverInterface
}

// NewPrimeCommand :nodoc:
func NewPrimeCommand(d receiver.DeviceReceiverInterface) Command {
	return &PrimeCommand{
		device: d,
	}
}

// Execute :nodoc:
func (c *PrimeCommand) Execute() {
	c.device.Prime()
}
