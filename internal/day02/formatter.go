package day02

import (
	"fmt"
	"log"
)

type dayTwoFormatter struct {
}

func (f dayTwoFormatter) formatData(strings <-chan string) (<-chan ProcessRequest, error) {
	ch := make(chan ProcessRequest, len(strings)/2)
	go func() {
		defer close(ch)
		defer fmt.Println("Data Formatting Completed")

		for s := range strings {
			switch s {
			case forward:
				ch <- ProcessRequest{
					Axis:   horizontal,
					Change: stringToInt(<-strings),
				}
			case down:
				ch <- ProcessRequest{
					Axis:   vertical,
					Change: stringToInt(<-strings),
				}
			case up:
				ch <- ProcessRequest{
					Axis:   vertical,
					Change: -stringToInt(<-strings),
				}
			default:
				log.Fatalf("unexpected value of %v", s)
			}
		}
	}()

	return ch, nil
}
