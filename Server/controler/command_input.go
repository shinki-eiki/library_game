/*
CommandLine结构中用来询问用户输入的方法
*/

package controler

import (
	"fmt"
)

// 询问一个整数
func (com *CommandLine) AskInteger(info string) (res int, err error) {
	fmt.Println(info)
	for {
		_, err := fmt.Scanln(&res)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
		fmt.Println("PLease input again:")
	}
	return
}

// 询问一个范围[low-1,up]内的整数，允许输入low-1是为了满足例如退出的操作，由调用者处理返回值
func (com *CommandLine) AskIntegerRange(info string, low, up int) (res int, err error) {
	fmt.Printf("%s, range from %v to %v.Input %v to exit.", info, low-1, up, low-1)
	for {
		_, err := fmt.Scanln(&res)
		if err != nil {
			// fmt.Println(res)
			fmt.Println(err)
		} else if res < low-1 || res > up {
			fmt.Println(res, "is not in the range!")
		} else {
			break
		}
		fmt.Println("PLease input again:")
	}
	return
}
