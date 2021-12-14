package day03

import (
	"advent_of_code/internal/data"
	"fmt"
	"log"
)

func Calculate(service Service) {
	dataStream, err := data.GetRuneSliceDataStream("input03.txt", data.NewLineSplit)
	if err != nil {
		log.Fatal(err)
	}

	respChan := service.ProcessData(dataStream)

	response := <-respChan

	if response.Error != nil {
		log.Fatalln(response.Error)
	}

	fmt.Println(response.Result)

	// convert data stream
}
