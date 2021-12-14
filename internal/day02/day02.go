package day02

import (
	"advent_of_code/internal/data"
	"fmt"
	"log"
)

func Process(service Service) {
	dataStream, err := data.GetStringDataStream("input02.txt", split)
	if err != nil {
		log.Fatalln(err)
	}

	processRequestStream, err := service.formatData(dataStream)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := service.processData(processRequestStream)
	if err != nil {
		log.Fatalln(err)
	}

	x := <-result
	fmt.Println(x)
}

func split(r rune) bool {
	return r == '\n' || r == ' '
}
