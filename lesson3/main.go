package main

import (
	less3mod2 "github.com/d-selifanov/less3mod/v2"
)

// Main function
func main() {
	less3mod2.PanicAndRecover()
	less3mod2.PanicAndRecoverWithTimestamp()
	less3mod2.PanicInParallelStream()
}
