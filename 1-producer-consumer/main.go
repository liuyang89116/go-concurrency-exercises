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

func producer(stream Stream, c chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}
		c <- tweet
	}
	close(c)
}

func consumer(c <-chan *Tweet) {
	for t := range c {
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

	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan *Tweet)

	// Producer
	go func() {
		producer(stream, c)
		wg.Done()
	}()

	// Consumer
	go func() {
		consumer(c)
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
