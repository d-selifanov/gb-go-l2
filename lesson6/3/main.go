/*
3. Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”
*/
package main

import (
	"fmt"
	"sync"
)

const cnt = 1000

func main() {
	var (
		counter int
		wg      sync.WaitGroup
	)
	wg.Add(cnt)
	for i := 0; i < cnt; i += 1 {
		go func() {
			defer wg.Done()
			counter += 1
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
