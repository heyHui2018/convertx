package convertx

import (
	"testing"
)

/*
运行单元测试： go test -v convert_test.go convert.go
运行指定单元测试：go test -v -run TestSnake2Hump convert_test.go convert.go
*/
func TestSnake2Hump(t *testing.T) {
	s := "hello_world"
	res1 := Snake2Hump(s, false)
	res2 := Snake2Hump(s, true)
	t.Logf("res1 = %v,res2 = %v", res1, res2)
}

func TestHump2Snake(t *testing.T) {
	s := "HelloWorld"
	res := Hump2Snake(s)
	t.Logf("res = %v", res)
}

func TestSeniorSplit(t *testing.T) {
	s := ",a,,,,b,,c,"
	res := SeniorSplit(s, ",")
	t.Logf("res = %v", res)
}
