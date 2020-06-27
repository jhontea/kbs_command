package command

import (
	"github.com/jhontea/kbs_command/receiver"
)

// SumCommand :nodoc
type SumCommand struct {
	device receiver.DeviceReceiverInterface
}

// NewSumCommand :nodoc:
func NewSumCommand(d receiver.DeviceReceiverInterface) Command {
	return &SumCommand{
		device: d,
	}
}

// Execute :nodoc:
func (c *SumCommand) Execute() {
	c.device.Sum()
}
