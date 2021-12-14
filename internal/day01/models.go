package day01

type service interface {
	calculate(input <-chan int, output chan<- int)
}
