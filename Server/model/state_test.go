package model_test

import (
	"fmt"
	. "patchouli_lib/model"
	"testing"
)

func TestState(t *testing.T) {
	// 测试用pv的指针构造的state能不能修改原pv，可以
	pv := &PublicValue{}
	s := &State{PublicValue: pv} //CurrentState: 0,
	// s.Change()

	fmt.Println(pv.DrawCardNumber)
	fmt.Println(s.DrawCardNumber)
	s.DrawCardNumber = 2
	fmt.Println(pv.DrawCardNumber)
	fmt.Println(s.DrawCardNumber)
}
