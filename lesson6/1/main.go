/*
1. Написать программу, которая использует мьютекс для безопасного доступа к данным
из нескольких потоков. Выполните трассировку программы
*/
package main

import (
	"os"
	"runtime/trace"
	"sync"
)

const cnt = 1000

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var (
		arr [cnt]int
		mu  sync.Mutex
		wg  sync.WaitGroup
	)
	wg.Add(cnt)
	for i := 0; i < cnt; i += 1 {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			if i%2 == 0 {
				arr[i] = i
			} else {
				arr[i] = i + 1
			}
		}(i)
	}
	wg.Wait()
}
