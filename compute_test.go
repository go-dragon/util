package util

import (
	"reflect"
	"testing"
)

func TestTwoDecimalPlaces(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{value: 3.35333}, want: 3.35},
		{name: "t2", args: args{value: 3.5588}, want: 3.55},
		{name: "t2", args: args{value: 3.5558}, want: 3.55},
		{name: "t2", args: args{value: 19.9}, want: 19.9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoDecimalPlaces(tt.args.value); got != tt.want {
				t.Errorf("Decimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindNumsInString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{"满100减50"}, want: []string{"100", "50"}},
		{name: "t2", args: args{"3210kdad32"}, want: []string{"3210", "32"}},
		{name: "t3", args: args{"0321kdad132"}, want: []string{"0321", "132"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNumsInString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindNumsInString() = %v, want %v", got, tt.want)
			}
		})
	}
}
