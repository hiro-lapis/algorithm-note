package main

import "testing"

func TestIsPrime(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "100003 expect: true",
			args: args{100003},
			want: true,
		}, {
			name: "100004 expect: true",
			args: args{100004},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.n); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
