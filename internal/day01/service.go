package day01

import "sync"

type PartOneService struct {
}

func (p PartOneService) calculate(input <-chan int, output chan<- int) {
	var lastInt int
	var numIncreased int
	once := sync.Once{}

	for i := range input {
		once.Do(func() {
			lastInt = i
		})
		if i > lastInt {
			numIncreased++
		}
		lastInt = i
	}

	output <- numIncreased
}

type PartTwoService struct {
}

func (p PartTwoService) calculate(input <-chan int, output chan<- int) {
	var lastTotal int
	var numIncreased int
	var index int
	var count int
	var startupComplete bool
	inputs := make([]int, 3)

	for i := range input {
		index = (index + 1) % 3
		inputs[index] = i
		if !startupComplete {
			count++
			lastTotal = lastTotal + i
			if count == 3 {
				startupComplete = true
			}
			continue
		}

		var total int
		for _, value := range inputs {
			total = total + value
		}

		if total > lastTotal {
			numIncreased++
		}
		lastTotal = total
	}

	output <- numIncreased
}
