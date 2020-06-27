package invoker

import (
	"github.com/jhontea/kbs_command/command"
)

// Invoker :nodoc:
type Invoker struct {
	sumCommand       command.Command
	multiplyCommand  command.Command
	primeCommand     command.Command
	fibonacciCommand command.Command
}

// NewInvoker :nodoc:
func NewInvoker(sum, multiply, prime, fibonacci command.Command) Invoker {
	return Invoker{
		sumCommand:       sum,
		multiplyCommand:  multiply,
		primeCommand:     prime,
		fibonacciCommand: fibonacci,
	}
}

// Sum :nodoc:
func (i *Invoker) Sum() {
	i.sumCommand.Execute()
}

// Multiply :nodoc:
func (i *Invoker) Multiply() {
	i.multiplyCommand.Execute()
}

// Prime :nodoc:
func (i *Invoker) Prime() {
	i.primeCommand.Execute()
}

// Fibonacci :nodoc:
func (i *Invoker) Fibonacci() {
	i.fibonacciCommand.Execute()
}
