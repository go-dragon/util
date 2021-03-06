package util

import (
	"fmt"
	"testing"
)

//初始化page和pageSize, 默认page为0, pageSize=10, 返回默认为字符串，方便之后查询拼接字符串
func TestInitPageAndPageSize(t *testing.T) {
	off, limit := InitPageAndPageSize("20", "10")
	if off != 200 || limit != 10 {
		t.Error()
		return
	}
	fmt.Println(InitPageAndPageSize("0.3", "10.2"))
	fmt.Println(InitPageAndPageSize("-2", "-3"))
	fmt.Println(InitPageAndPageSize("3", "dasd"))
}
