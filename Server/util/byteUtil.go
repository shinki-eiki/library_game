package util

import (
	"fmt"
)

// ConvertToBytes 将任意数量的整数或字符串参数转换为 byte 类型的数组
func ConvertToBytes(values ...interface{}) []byte {
	var bytes []byte

	for _, v := range values {
		switch v := v.(type) {
		case int:
			// 将整数转换为 byte 类型
			bytes = append(bytes, byte(v+48))
		case int32:
			// 将int32不用+48，就可以转换为 byte 类型
			bytes = append(bytes, byte(v))
		case string:
			// 将字符串转换为 byte 切片
			bytes = append(bytes, []byte(v)...)
		case byte: // 原来就是byte就不用管它了，但是添加还是得添加
			bytes = append(bytes, v)
		default:
			// 处理未知类型
			fmt.Printf("Unsupported type: %T\n", v)
		}
	}

	return bytes
}

// 将一个byte数组转变成相应的整数
// 由于传递消息格式的原因，第一个数忽略
// 例如0{1,2}转变为12，注意
func ConvertBytesToInt(arr []byte) int {
	res := 0
	for i := 1; i < len(arr); i++ {
		res *= 10
		res += int(arr[i])
	}
	return res
}
