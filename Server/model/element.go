package model

import (
	"fmt"
	"strconv"
)

const (
// Metal="金"
)

var PROPERTYMAP = [5]string{"金", "木", "水", "火", "土"}

/* 元素卡 */
type Element struct {
	Property int //Element type
	Index    int //Image index
}

func (e *Element) GetImage() string {
	return strconv.Itoa(e.Property) + "_" + strconv.Itoa(e.Index)
}

// 输出 ： 属性（卡面索引）
func (e *Element) Str() (res string) {
	res = fmt.Sprintf("%v-%v", PROPERTYMAP[e.Property], e.Index)
	// res = fmt.Sprintf("%d(%d)", e.Property, e.Index)
	return res
}

// func NewElement(p,i int)  {
// 	e:=&Element{p,i}
// 	return e
// }
