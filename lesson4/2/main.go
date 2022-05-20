/*
2. Написать программу, которая при получении в канал сигнала SIGTERM останавливается не
позднее, чем за одну секунду (установить таймаут)
*/

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancelSig := context.WithCancel(context.Background())
	chSigs := make(chan os.Signal, 1)
	signal.Notify(chSigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Waiting close signal")
			}

		}
	}()

	sig := <-chSigs

	cancelSig()

	fmt.Println("Received SIGTERM", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()
	<-ctx.Done()
}
