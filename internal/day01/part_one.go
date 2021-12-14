package day01

import (
	"advent_of_code/internal/data"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Calculate(service service) {
	d, err := data.GetData("input01.txt", data.NewLineSplit)
	if err != nil {
		log.Fatalln(err)
	}

	dataList := strings.Split(d, "\n")

	ch := make(chan int, len(dataList))
	output := make(chan int)

	go service.calculate(ch, output)

	err = convertStringsToInts(dataList, ch)
	if err != nil {
		log.Fatalln(err)
	}
	close(ch)
	numIncreased := <-output
	fmt.Println(numIncreased)
}

func convertStringsToInts(strs []string, ch chan<- int) error {
	for _, s := range strs {
		if s == "" {
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		ch <- i
	}
	return nil
}
