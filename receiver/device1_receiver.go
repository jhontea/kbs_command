package receiver

import "fmt"

// Device1Receiver :nodoc:
type Device1Receiver struct {
	a int
	b int
	n int
}

// NewDevice1Receiver :nodoc:
func NewDevice1Receiver() DeviceReceiverInterface {
	return &Device1Receiver{}
}

// SetA set input first number store to a
func (d *Device1Receiver) SetA(a int) {
	d.a = a
}

// SetB set input second number store to b
func (d *Device1Receiver) SetB(b int) {
	d.b = b
}

// SetN set input number store to n
func (d *Device1Receiver) SetN(n int) {
	d.n = n
}

// Sum :nodoc:
func (d *Device1Receiver) Sum() int {
	result := d.a + d.b
	fmt.Printf("Sum %d + %d = %d\n", d.a, d.b, result)

	return result
}

// Multiply :nodoc:
func (d *Device1Receiver) Multiply() int {
	result := d.a * d.b
	fmt.Printf("Multiply %d * %d = %d\n", d.a, d.b, result)

	return result
}

// Prime :nodoc:
func (d *Device1Receiver) Prime() []int {
	results := []int{}
	flag := false
	i := 0

	if d.n < 1 {
		fmt.Printf("n < 1")
		return results
	}

	for len(results) < d.n {
		if i == 1 || i == 0 {
			i++
			continue
		}

		flag = true

		for j := 2; j < i/2+1; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}

		if flag {
			results = append(results, i)
		}

		i++
	}

	fmt.Printf("First %d prime: %v\n", d.n, results)

	return results
}

// Fibonacci :nodoc:
func (d *Device1Receiver) Fibonacci() []int {
	results := []int{}
	f1 := 0
	f2 := 1

	if d.n < 1 {
		fmt.Printf("n < 1")
		return results
	}

	for i := 0; i < d.n; i++ {
		results = append(results, f1)

		next := f1 + f2
		f1 = f2
		f2 = next
	}

	fmt.Printf("First %d fibonacci: %v\n", d.n, results)

	return results
}
