package main

import "testing"

func TestCollected(t *testing.T) {
	type args struct {
		l []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "collect",
			args: args{[]bool{true, true, true, false, false, false, false, false, false}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Collected(tt.args.l); got != tt.want {
				t.Errorf("Collected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetKey(t *testing.T) {
	type args struct {
		s []int
		e int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "contain",
			args: args{[]int{3, 4, 1}, 1},
			want: 2,
		},
		{
			name: "not contain",
			args: args{[]int{3, 4}, 1},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetKey(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExists(t *testing.T) {
	type args struct {
		s []int
		e int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "exists",
			args: args{[]int{3, 4, 1}, 1},
			want: true,
		},
		{
			name: "not exists",
			args: args{[]int{3, 4}, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exists(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckT(t *testing.T) {
	type args struct {
		s []bool
		i []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name:       "fail",
			args:       args{[]bool{true, false, true}, []int{0, 2}},
			wantResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := CheckT(tt.args.s, tt.args.i); gotResult != tt.wantResult {
				t.Errorf("CheckT() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
