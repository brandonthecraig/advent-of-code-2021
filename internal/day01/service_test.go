package day01

import "testing"

func TestPartTwoService_calculate(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "no input",
			input: nil,
			want:  0,
		},
		{
			name:  "one value",
			input: []int{1},
			want:  0,
		},
		{
			name:  "two values",
			input: []int{1, 2},
			want:  0,
		},
		{
			name:  "three values",
			input: []int{1, 2, 3},
			want:  0,
		},
		{
			name:  "four values, one group increase",
			input: []int{1, 2, 2, 2},
			want:  1,
		},
		{
			name:  "four values, no group increase",
			input: []int{2, 2, 2, 2},
			want:  0,
		},
		{
			name:  "five values, two group increases",
			input: []int{2, 2, 2, 3, 4},
			want:  2,
		},
		{
			name:  "five values, one group increases",
			input: []int{2, 2, 2, 3, 0},
			want:  1,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan int, len(tt.input))
			out := make(chan int)

			p := PartTwoService{}
			go p.calculate(ch, out)
			for _, i := range tt.input {
				ch <- i
			}
			close(ch)

			got := <-out
			if tt.want != got {
				t.Errorf("got %v, want %v", got, tt.want)
			}

		})
	}
}
