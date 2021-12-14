package day02

import (
	"log"
	"strconv"
)

type dayTwoPartAProcessor struct{}

func (p dayTwoPartAProcessor) processData(requests <-chan ProcessRequest) (<-chan int, error) {
	ch := make(chan int, 1)
	go func() {
		var horizontalChange int
		var verticalChange int
		for r := range requests {
			switch r.Axis {
			case horizontal:
				horizontalChange = horizontalChange + r.Change
			case vertical:
				verticalChange = verticalChange + r.Change
			}
		}
		ch <- horizontalChange * verticalChange
	}()
	return ch, nil
}

type dayTwoPartBProcessor struct{}

func (p dayTwoPartBProcessor) processData(requests <-chan ProcessRequest) (<-chan int, error) {
	ch := make(chan int, 1)
	go func() {
		var horizontalChange int
		var verticalChange int
		var aim int
		for r := range requests {
			switch r.Axis {
			case horizontal:
				horizontalChange = horizontalChange + r.Change
				verticalChange = verticalChange + r.Change*aim
			case vertical:
				aim = aim + r.Change
			default:
				log.Fatal("unexpected input")
			}
		}
		ch <- horizontalChange * verticalChange
	}()
	return ch, nil
}

// code smell with error handling here, look up how to do it better
func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}
