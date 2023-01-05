package main

import "testing"

func TestPower(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "not quadratic",
			args: args{5, 1},
			want: 5,
		},
		{
			name: "quadratic",
			args: args{5, 3},
			want: 125,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Power(tt.args.a, tt.args.b)
		})
	}
}
