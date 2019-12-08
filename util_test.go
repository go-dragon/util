package util

import (
	"dragon/model"
	"fmt"
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

func TestStruct2Map(t *testing.T) {
	device := model.TDevice{
		DeviceId:       2,
		DeviceName:     "",
		DeviceType:     "",
		DeviceToken:    "",
		DeviceCode:     "",
		DeviceInfo:     "",
		DevicePosition: "",
		Longitude:      0,
		Latitude:       0,
		StoreCode:      "",
		ChannelInfo:    "",
		DeviceOwnerId:  0,
		DeviceStatus:   0,
		CreateTime:     "",
		UpdateTime:     "",
	}
	fmt.Println(StructJsonTagToMap(device))
}
