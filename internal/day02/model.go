package day02

type Service interface {
	formatter
	processor
}

type formatter interface {
	formatData(<-chan string) (<-chan ProcessRequest, error)
}

type processor interface {
	processData(<-chan ProcessRequest) (<-chan int, error)
}

type ProcessRequest struct {
	Axis   string
	Change int
}
