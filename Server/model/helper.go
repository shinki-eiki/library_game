package model

import (
	"fmt"
)

/* 帮手卡 */
type Helper struct {
	name    string
	text    string
	CanUse  bool
	score   int
	Ability func()
}

// var Hong = Helper{
// 	false,
// 	func() {},
// }

func (e *Helper) GetImage() string {
	return e.name + ":" + e.text
}

func (e *Helper) Str() string {
	res := fmt.Sprintf("%v(%v):%s", e.name, e.score, e.text)
	return res
}
