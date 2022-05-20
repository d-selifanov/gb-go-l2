/*
!!! Взял задание из методички, оно там отличается от того, что было на видео.
Но в прикрепленной ссылке для презентации почему-то тоже самое что методичка...
https://gbcdn.mrgcdn.ru/uploads/asset/3791895/attachment/12f01675785079709a6132d8e29054db.pdf

1. С помощью пула воркеров написать программу, которая запускает 1000 горутин, каждая из
которых увеличивает число на 1. Дождаться завершения всех горутин и убедиться, что при
каждом запуске программы итоговое число равно 1000.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count = 0
	var chWorkers = make(chan int, 1000)
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		chWorkers <- 1

		go func(m *sync.Mutex) {
			m.Lock()
			count += <-chWorkers
			m.Unlock()
		}(&m)
	}

	time.Sleep(time.Second * 1)

	fmt.Printf("Count: %v\n", count)
}
