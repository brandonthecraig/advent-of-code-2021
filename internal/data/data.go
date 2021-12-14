package data

import (
	"os"
	"strings"
)

func GetRuneSliceDataStream(filename string, splitFunc func(r rune) bool) (<-chan []rune, error) {
	dataList, err := GetSplitData(filename, splitFunc)
	if err != nil {
		return nil, err
	}

	ch := make(chan []rune, len(dataList))
	go func() {
		defer println("Data Stream Generation Completed")
		defer close(ch)
		for _, s := range dataList {
			ch <- []rune(s)
		}
	}()
	return ch, nil
}

func GetStringDataStream(filename string, splitFunc func(r rune) bool) (<-chan string, error) {
	dataList, err := GetSplitData(filename, splitFunc)
	if err != nil {
		return nil, err
	}

	ch := make(chan string, len(dataList))
	go func() {
		defer println("Data Stream Generation Completed")
		defer close(ch)
		for _, s := range dataList {
			ch <- s
		}
	}()
	return ch, nil
}

func GetSplitData(fileName string, splitFunc func(r rune) bool) ([]string, error) {
	data, err := GetData(fileName)
	if err != nil {
		return nil, err
	}
	return strings.FieldsFunc(string(data), splitFunc), nil
}

func GetData(fileName string) (string, error) {
	data, err := os.ReadFile("internal/data/source/" + fileName)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
