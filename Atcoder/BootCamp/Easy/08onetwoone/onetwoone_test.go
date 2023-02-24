package main

import "testing"

func TestIsInt(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "expect: true",
			args: args{10.0000},
			want: true,
		},
		{
			name: "expect: false",
			args: args{4.000000001},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInt(tt.args.x); got != tt.want {
				t.Errorf("IsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
