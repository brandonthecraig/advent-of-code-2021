package day02

import (
	"reflect"
	"strings"
	"testing"
)

func Test_split(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "no breaks",
			input: "abc",
			want:  []string{"abc"},
		},
		{
			name:  "space breaks",
			input: "a b c ",
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "space and newline breaks",
			input: "a b\nc ",
			want:  []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strings.FieldsFunc(tt.input, split)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want %v\ngot  %v", tt.want, got)

			}
		})
	}
}
