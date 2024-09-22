package util

import (
	"fmt"
	"testing"
)

// 测试将其他类型的多个数据转变成byte数组的函数
func TestBytes(t *testing.T) {
	result := ConvertToBytes(65, 66, 67, "Hello", 99, "World", byte(100))
	fmt.Printf("%q\n", result)
}

// 测试int、string和byte的相互转化规则
func TestIntToByte(t *testing.T) {
	var a int = 0
	fmt.Println(byte(a))
	res := ConvertToBytes(0)
	fmt.Println(res)

	var s string = "a01"
	res = ConvertToBytes(s)
	fmt.Println(res)
}

// 测试byte转变成int的规则
func TestByteToInt(t *testing.T) {
	var b byte = 0
	fmt.Print(int(b))
	var s string = "a98s"
	fmt.Print([]byte(s))
}

func TestConvertBytesToInt(t *testing.T) {
	var bs []byte = []byte{'a', 1, 2, 3}
	res := ConvertBytesToInt(bs)
	fmt.Println(res)
}
