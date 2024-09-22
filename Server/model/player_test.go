package model_test

import (
	. "patchouli_lib/model"
	"testing"
)

// 新建玩家
func TestInitPlayers(t *testing.T) {
	ns := []string{"Reimu", "Marisa", "Alice", "Youmu"}
	ps := NewPlayers(ns)
	for _, v := range ps {
		// fmt.Print(v)
		v.Show()
	}
}
