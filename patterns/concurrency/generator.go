package concurrency

import "fmt"

// writer sends values to the channel and closes it.
func writer(n int, ch chan<- int) {
	// handle closing the channel to notify that reader
	// and prevent goroutine leak
	go func() {
		defer close(ch)
		for i := range n {
			ch <- i
		}
	}()
}

// reader reads values from the channel until it's closed.
// if the channel is never closed, the range loop blocks forever
// waiting for more values, which causes a deadlock in this example.
//
// running reader in a separate goroutine would lead to a goroutine leak
// because the goroutine never exits and continues consuming resources.
func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func PatternGenerator() {
	ch := make(chan int)
	writer(10, ch)

	reader(ch)
	fmt.Println("exited")
}
