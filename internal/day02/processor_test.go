package day02

import (
	"reflect"
	"testing"
)

func Test_dayTwoPartBProcessor_processData(t *testing.T) {
	tests := []struct {
		name    string
		input   []ProcessRequest
		want    int
		wantErr bool
	}{
		{
			name:    "no entries",
			input:   nil,
			want:    0,
			wantErr: false,
		},
		{
			name: "horizontal 2, vertical 2",
			input: []ProcessRequest{{
				Axis:   horizontal,
				Change: 2,
			}, {
				Axis:   vertical,
				Change: 2,
			}},
			want:    0,
			wantErr: false,
		},
		{
			name: "vertical 2, horizontal 2",
			input: []ProcessRequest{{
				Axis:   vertical,
				Change: 2,
			},
				{
					Axis:   horizontal,
					Change: 2,
				}},
			want:    8,
			wantErr: false,
		},
		//{
		//	name: "example from requirements",
		//	input: []ProcessRequest{
		//		{
		//			Axis:
		//		},
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan ProcessRequest, len(tt.input))
			for _, req := range tt.input {
				ch <- req
			}
			close(ch)

			p := dayTwoPartBProcessor{}
			gotChannel, err := p.processData(ch)
			if (err != nil) != tt.wantErr {
				t.Errorf("processData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := <-gotChannel

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
