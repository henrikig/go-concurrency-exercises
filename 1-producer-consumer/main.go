//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, c chan<- *Tweet, wg *sync.WaitGroup) (tweets []*Tweet) {
	defer wg.Done()
	for {
		tweet, err := stream.Next()

		if err == ErrEOF {
			close(c)
			return
		}

		c <- tweet
	}
}

func consumer(tweets <-chan *Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	wg := new(sync.WaitGroup)
	c := make(chan *Tweet)

	// Producer
	wg.Add(1)
	go producer(stream, c, wg)

	// Consumer
	wg.Add(1)
	go consumer(c, wg)

	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
