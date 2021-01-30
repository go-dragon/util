package util

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHmacSha1(t *testing.T) {
	fmt.Println(HmacSha1("dasdaf", "123"))
}

func BenchmarkHmacSha1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HmacSha1("dasdaf", "123")
	}
}

func TestHmacMD5(t *testing.T) {
	fmt.Println(HmacMD5("dasdaf", "123"))
	fmt.Println(HmacMD5("dasdaf", "123"))
}

func BenchmarkHmacMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HmacMD5("dasdaf", "123")
	}
}

func TestOnlyCols(t *testing.T) {
	params := map[string]string{
		"p1": "hello",
		"p2": "hi",
		"p3": "good",
	}
	OnlyCols([]string{"p1", "p2", "p5"}, params)
	if params["p1"] != "hello" || params["p2"] != "hi" {
		t.Error("TestOnlyCols fail")
		return
	}
	fmt.Println(params)
}

func BenchmarkOnlyCols(b *testing.B) {
	params := map[string]string{
		"p1": "hello",
		"p2": "hi",
		"p3": "good",
	}
	OnlyCols([]string{"p1", "p2", "p5"}, params)
	for i := 0; i < b.N; i++ {
		OnlyCols([]string{"p1", "p2", "p5"}, params)
	}
}

func TestGetStructFields(t *testing.T) {
	var s = struct {
		DeviceName string `json:"device_name"`
		UserName   string `json:"user_name"`
	}{}
	fields := GetStructFields(s)
	fmt.Println(fields)
}

func TestGetStructJsonTags(t *testing.T) {
	var s = struct {
		DeviceName string `json:"device_name"`
		UserName   string `json:"user_name"`
	}{}
	tags := GetStructJsonTags(s)
	fmt.Println(tags)
}

func TestMapStringToInterface(t *testing.T) {
	type args struct {
		data map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{data: map[string]string{"age": "18", "username": "gg"}}, want: map[string]interface{}{"age": "18", "username": "gg"}},
		{name: "t2", args: args{data: map[string]string{"data": "flsow", "klo": "ppy"}}, want: map[string]interface{}{"data": "flsow", "klo": "ppy"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapStringToInterface(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapStringToInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapInterfaceToString(t *testing.T) {
	type args struct {
		data map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{data: map[string]interface{}{"age": "18", "username": "gg"}}, want: map[string]string{"age": "18", "username": "gg"}},
		{name: "t2", args: args{data: map[string]interface{}{"data": "flsow", "klo": "ppy"}}, want: map[string]string{"data": "flsow", "klo": "ppy"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapInterfaceToString(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapInterfaceToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
