/*
2. Написать многопоточную программу, в которой будет использоваться явный вызов
планировщика. Выполните трассировку программы
*/

package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	runtime.GOMAXPROCS(2)
	runtime.Gosched()

	start := time.Now()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println("Total time: " + elapsedTime.String())

	time.Sleep(2 * time.Second)
}
