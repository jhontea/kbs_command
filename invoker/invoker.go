package invoker

import (
	"github.com/jhontea/kbs_command/command"
)

// Invoker :nodoc:
type Invoker struct {
	Command command.Command
}

// Do :nodoc:
func (i *Invoker) Do() {
	i.Command.Execute()
}
