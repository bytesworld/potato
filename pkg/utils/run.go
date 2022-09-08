package utils

import (
	"fmt"
	"time"
)

func runFuncDuration(duration time.Duration, customFunc func()) {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(duration)
		timeout <- true
	}()
	go func() {
		customFunc()
	}()
	<-timeout
}

func long()  {
	for true {
		fmt.Println(time.Now())
	}
}

func test()  {
	runFuncDuration(time.Second,long)
}