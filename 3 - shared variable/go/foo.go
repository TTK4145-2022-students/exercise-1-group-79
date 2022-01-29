// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0

func incrementing() {
	//TODO: increment i 1000000 times
<<<<<<< HEAD
	for i < 1000000 {
=======
	for i < 100000 {
>>>>>>> e6ebb58b85ac9ef50ecfc490777d22c2f4165f7a
		i++
	}
}

func decrementing() {
	//TODO: decrement i 1000000 times
<<<<<<< HEAD
	for i > -1000000 {
=======

	for i > -100000 {
>>>>>>> e6ebb58b85ac9ef50ecfc490777d22c2f4165f7a
		i--
	}
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	// TODO: Spawn both functions as goroutines
	go incrementing()
	go decrementing()
<<<<<<< HEAD
=======

>>>>>>> e6ebb58b85ac9ef50ecfc490777d22c2f4165f7a
	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
