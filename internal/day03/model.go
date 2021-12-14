package day03

type Service interface {
	Processor
}

type Formatter interface {
	FormatData(<-chan string) <-chan []rune
}

type Processor interface {
	ProcessData(<-chan []rune) <-chan ProcessResponse
}

type PartOneService struct {
	DayThreeFormatter
	DayThreePartAProcessor
}

//
//type ProcessRequest struct {
//	Request   []rune
//}

type ProcessResponse struct {
	Result int
	Error  error
}
