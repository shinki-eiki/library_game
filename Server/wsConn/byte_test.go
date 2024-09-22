// 测试一下websocket中的数据传递

package wsConn_test

import (
	"fmt"
	"testing"
)

func TestBytes(t *testing.T) {
	var a byte = 97
	// string(byte b)return a string according to the ascll
	fmt.Println(a, int(a), string(a))
	fmt.Printf("%v", a)
	// byte is also can be index of array
	// arr := []int{1, 2, 3}
	// fmt.Print(arr[a])
}

// 将byte类型转成int，是直接适用ascll码作为值的
func TestByteToInt(t *testing.T) {
	var b byte = 48
	// var a int=114
	fmt.Println(int(b))
	fmt.Println(b)
	switch b {
	case '0':
		fmt.Println("0")

		// case 48:
		// 	fmt.Println("48")
	}

}
