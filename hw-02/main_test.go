package main

import "testing"

func TestRepeat(t *testing.T) {
	type args struct {
		character string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name:"1", args:args{"a"},want:"aaaab"},
		{name:"2", args:args{"b"},want:"bbbbb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.character); got != tt.want {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}