package main

import (
	"fmt"
	"time"
)

var counter int

// increment function increments the counter variable by 1 every 500 milliseconds for 5 times and prints the value of counter variable.
func increment() {
	for i := 0; i < 5; i++ {
		counter++
		fmt.Println("increment:", counter)
		time.Sleep(time.Millisecond * 500)
	}
}

// decrement function decrements the counter variable by 1 every 500 milliseconds for 5 times and prints the value of counter variable.
func decrement() {
	for i := 0; i < 5; i++ {
		counter--
		fmt.Println("decrement:", counter)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// goroutines are used to execute the functions concurrently.
	go increment()
	go decrement()
	// sleep function is used to wait for 5 seconds before the program terminates.
	// if we don't use sleep function, the program will terminate before the goroutines are executed.
	time.Sleep(time.Second * 5)
}
