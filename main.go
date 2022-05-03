package main

import "gb-go-l2/lesson2/lesson2"

// Main function
func main() {
	lesson2.PanicAndRecover()
	lesson2.PanicAndRecoverWithTimestamp()
	//lesson2.CreateFiles()
	lesson2.PanicInParallelStream()
}
