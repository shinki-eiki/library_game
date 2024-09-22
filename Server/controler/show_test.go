package controler

import (
	"fmt"
	"patchouli_lib/model"
	"testing"
)

var OUPUTER = CommandLine{}

//	func TestGain(t *testing.T) {
//		h := model.HandCard{Data: make([]*model.Element, 0)}
//		for i := 0; i < 5; i++ {
//			h.Append(&model.Element{Property: i, Index: 1})
//		}
//		h.Append(&model.Element{Property: 2, Index: 3})
//		h.Show()
//	}

// 测试符卡展示的函数
func TestShowSpellCard(t *testing.T) {
	sp := model.SpellCardDeckInit()
	// sp.Show()
	OUPUTER.showSpellDeck(sp.Display)
}

// 测试元素卡展示的函数
func TestShowElementCard(t *testing.T) {
	sp := model.ElementCardDeckInit()
	// sp.Show()
	OUPUTER.showElementDeck(sp.Display)
}

// 测试玩家展示的函数
func TestShowPlayers(t *testing.T) {
	p := model.NewPlayers([]string{"Reimu", "Kochiya"})
	h := &model.HandCard{Data: make([]*model.Element, 0)}
	for i := 0; i < 5; i++ {
		h.Append(&model.Element{Property: i, Index: 1})
	}
	h.Append(&model.Element{Property: 2, Index: 3})
	p[1].Hand = h
	OUPUTER.showPlayer(p)
}

func TestShowOnePlayers(t *testing.T) {
	p := model.NewPlayers([]string{"Reimu", "Kochiya"})
	h := &model.HandCard{Data: make([]*model.Element, 0)}
	for i := 0; i < 5; i++ {
		h.Append(&model.Element{Property: i, Index: 1})
	}
	h.Append(&model.Element{Property: 2, Index: 3})
	p[1].Hand = h
	OUPUTER.showOnePlayer(p[1], true)
}

// 测试一下go里的数组有什么特性
func TestArray(t *testing.T) {
	arr := [5]int{}
	temp := arr
	arr[2] -= 23
	// arr := make([]int, 5)
	// arr := [5]int{}
	// 有长度的数组，如果不初始化，会采用默认值
	fmt.Println(arr)
	fmt.Println(temp)
	// arr[2] = 114
	// fmt.Println(arr)
}

// 测试从命令行读取数据
func TestInput(t *testing.T) {
	var a int
	fmt.Println("Before reading", a)
	fmt.Println(">>>Please Input the number")
	// fmt.Scanf("%d", &a)
	// fmt.Sscanf("114514\n", "%d", &a)
	a, _ = fmt.Sscanf("114514\n", "%d", a)
	fmt.Println("After reading", a)
}

func TestNewGame(t *testing.T) {
	g := model.NewGame([]string{"Reimu", "Kochiya"}, nil, nil)
	OUPUTER.showAll(g)
}
