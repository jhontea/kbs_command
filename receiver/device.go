package receiver

// DeviceReceiverInterface :nodoc:
type DeviceReceiverInterface interface {
	SetA(a int)
	SetB(b int)
	SetN(n int)
	Sum() int
	Multiply() int
	Prime() []int
	Fibonacci() []int
}
