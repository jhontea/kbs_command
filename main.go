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

	var a, b, n int

	// invoke sum a + b
	a = 1
	b = 2
	device1.SetA(a)
	device1.SetB(b)
	invoke.Sum()

	// invoke multiply a * b
	a = 1
	b = 3
	device1.SetA(a)
	device1.SetB(b)
	invoke.Multiply()

	// invoke first n prime
	n = 4
	device1.SetN(n)
	invoke.Prime()

	// invoke first n fibonacci
	n = 5
	device1.SetN(n)
	invoke.Fibonacci()
}
