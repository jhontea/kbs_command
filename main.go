package main

import (
	"github.com/jhontea/kbs_command/command"
	"github.com/jhontea/kbs_command/invoker"
	"github.com/jhontea/kbs_command/receiver"
)

func main() {
	// register new device
	device1 := receiver.NewDevice1Receiver()

	// register command
	sumCommand := command.NewSumCommand(device1)
	multiplyCommand := command.NewMultiplyCommand(device1)
	primeCommand := command.NewPrimeCommand(device1)
	fibonacciCommand := command.NewFibonacciCommand(device1)

	// set command for invoker
	invoke := invoker.NewInvoker(sumCommand, multiplyCommand, primeCommand, fibonacciCommand)
	invoke.Sum()
	invoke.Multiply()
	invoke.Prime()
	invoke.Fibonacci()
}
