package day03

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestDayThreePartAProcessor_ProcessData(t *testing.T) {
	tests := []struct {
		name  string
		input [][]rune
		want  ProcessResponse
	}{
		{
			name: "one input",
			input: [][]rune{
				[]rune("101"),
			},
			want: ProcessResponse{
				Result: generateExpected("101"),
				Error:  nil,
			},
		},
		{
			name: "three inputs",
			input: [][]rune{
				[]rune("1010"),
				[]rune("1101"),
				[]rune("0101"),
			},
			want: ProcessResponse{
				Result: generateExpected("1101"),
				Error:  nil,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DayThreePartAProcessor{}
			ch := make(chan []rune, len(tt.input))
			for _, r := range tt.input {
				ch <- r
			}
			close(ch)

			gotCh := d.ProcessData(ch)
			got := <-gotCh

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateExpected(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "ten",
			input: "101",
			want:  10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateExpected(tt.input)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}

func generateExpected(s1 string) int {
	replacer := strings.NewReplacer("1", "0", "0", "1")
	s2 := replacer.Replace(s1)

	i1, err := strconv.ParseInt(s1, 2, 64)
	if err != nil {
		return 0
	}

	i2, err := strconv.ParseInt(s2, 2, 64)
	if err != nil {
		return 0
	}

	return int(i1 * i2)
}
