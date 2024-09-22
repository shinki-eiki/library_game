package model_test

import (
	"fmt"
	. "patchouli_lib/model"
	"testing"
)

// 手卡的增加及删除
func TestAppend(t *testing.T) {
	h := HandCard{make([]*Element, 0)}
	h.Show()
	h.Append(&Element{2, 3})
	h.Show()
	for i := 0; i < 5; i++ {
		h.Append(&Element{i, 1})
	}
	h.Show()

	e := h.Remove(0)
	if e != nil {
		t.Fatal("Index over.")
	}
	h.Show()

	e = h.Remove(3)

	if e != nil {
		fmt.Println("...Fail to remove.")
	}
	h.Show()
}

// 测试能否判断获取符卡以及资源消耗
func TestGain(t *testing.T) {
	h := HandCard{make([]*Element, 0)}
	for i := 0; i < 5; i++ {
		h.Append(&Element{i, 1})
	}
	h.Append(&Element{2, 3})
	h.Show()

	s := &SpellCard{"", [5]int{1, 1, 1, 1, 1}, 5}
	fmt.Println("Can get the spell card?", h.CanGain(s))
	fmt.Println(h.CanGain(&SpellCard{"", [5]int{1, 2, 1, 0, 0}, 3}))
	fmt.Println(s.Cost)

	h.PayFor(s)
	fmt.Println("After pay for the spell card...")
	h.Show()
}

// 测试能否判断获取帮手
func TestGainHelper(t *testing.T) {
	h := HandCard{make([]*Element, 0)}
	for i := 0; i < 5; i++ {
		h.Append(&Element{i, 1})
		fmt.Println(h.CanGainHelper())
	}
	h.Append(&Element{2, 3})
	h.Show()
}

// 测试打印为表
func TestTable(t *testing.T) {
	h := HandCard{make([]*Element, 0)}
	for i := 0; i < 5; i++ {
		h.Append(&Element{i, 1})
	}
	PropertyTable(h.Data)
}
