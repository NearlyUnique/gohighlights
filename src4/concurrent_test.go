package conc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_consume_channel(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const maxWork = 40
	chWork := make(chan int)
	chResults := make(chan string)

	// create 10 go routines
	for i := 0; i < 10; i++ {
		go func(id int) {
			for work := range chWork {
				// pretend work
				duration := time.Duration(rand.Intn(100)) * time.Millisecond
				time.Sleep(duration)
				workResult := fmt.Sprintf("[%d] processed %2d in %v", id, work, duration)
				// send the results
				chResults <- workResult
			}
		}(i)
	}
	// pump some work in
	go func() {
		for i := 0; i < maxWork; i++ {
			chWork <- i
		}
	}()

	// read the results
	for i := 0; i < maxWork; i++ {
		fmt.Println(<-chResults)
	}
	// tidy up
	close(chWork)
}
