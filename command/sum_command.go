package command

import (
	"github.com/jhontea/kbs_command/receiver"
)

// SumCommand :nodoc
type SumCommand struct {
	Device receiver.Device
}

// Execute :nodoc:
func (c *SumCommand) Execute() {
	c.Device.Sum()
}
