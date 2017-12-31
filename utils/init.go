package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var once bool

func WaitForExit(twice bool, closeCh ...chan int) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(
		sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	for {
		s := <-sigChan
		switch s {
		case syscall.SIGINT, syscall.SIGTERM:
			if twice && !once {
				once = true
				for _, v := range closeCh {
					select {
					case <-v:
					default:
						close(v)
					}
				}
				fmt.Println("Send ^C to force exit.")
			} else {
				os.Exit(0)
			}
		}
	}
}
