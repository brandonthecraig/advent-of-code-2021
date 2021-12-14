package day03

import "fmt"

type DayThreePartAProcessor struct {
}

func (d DayThreePartAProcessor) ProcessData(runeFeed <-chan []rune) <-chan ProcessResponse {
	ch := make(chan ProcessResponse, 1)

	go func() {
		defer close(ch)
		for runes := range runeFeed {
			fmt.Println(string(runes))
		}
		ch <- ProcessResponse{
			Result: 1,
			Error:  nil,
		}
	}()
	return ch
}
