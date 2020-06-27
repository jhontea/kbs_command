package command

import (
	"github.com/jhontea/kbs_command/receiver"
)

// PrimeCommand :nodoc
type PrimeCommand struct {
	Device receiver.Device
}

// Execute :nodoc:
func (c *PrimeCommand) Execute() {
	c.Device.Prime()
}
