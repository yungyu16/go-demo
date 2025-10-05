package main

import "time"

func main() {
	timer := time.NewTimer(1 * time.Second)
	if !timer.Stop() {
		<-timer.C
	}

}
