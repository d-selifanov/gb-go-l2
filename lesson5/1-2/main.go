/*
1. Напишите программу, которая запускает 𝑛 потоков и дожидается завершения их всех
2. Реализуйте функцию для разблокировки мьютекса с помощью defer
*/

package main

import (
	"fmt"
	"sync"
)

const n = 1000

func main() {
	var (
		counter int
		mutex   sync.Mutex
		wg      = sync.WaitGroup{}
	)
	wg.Add(n)
	for i := 0; i < n; i += 1 {
		go func(i int, mutex *sync.Mutex) {
			mutex.Lock()
			defer mutex.Unlock()
			counter += 1
			fmt.Println("Goroutine:", i+1)
			wg.Done()
		}(i, &mutex)
	}
	wg.Wait()
	fmt.Println("Total completed:", counter)
}
