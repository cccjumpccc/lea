package main

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestFoo(t *testing.T) {
	// t.Fatal("not implemented")
	c := make(chan os.Signal, 3)
	signal.Notify(c)

	for {
		select {
		case sig := <-c:
			fmt.Println(sig)
		default:
			time.Sleep(10 * time.Second)
		}
	}
}
