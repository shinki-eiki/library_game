package model_test

import (
	"fmt"
	. "patchouli_lib/model"
	"strconv"
	"strings"
	"testing"
)

func TestDemo(t *testing.T) {
}

// 元素卡的创建
func TestElement(t *testing.T) {
	e := Element{1, 2}
	fmt.Println(e)
	e.GetImage()
	// t.Log((e))
}

// 符卡的创建
func TestSpellCard(t *testing.T) {
	e := &SpellCard{"Test", [5]int{1, 2, 3, 4, 5}, 2}
	// fmt.Println(e)
	t.Log((e))
}

/* 泛型元素卡堆的测试函数 */
func TestGenericity(t *testing.T) {
	var eleDeck Deck[*Element] = []*Element{{2, 3}, {4, 5}}
	fmt.Println(len(eleDeck))
	res := &eleDeck
	r := res.Pop()
	fmt.Println(r.Str())
	fmt.Println(len(eleDeck))

	r = res.Pop()
	fmt.Println(r.Str())

}

/* 泛型符卡堆的测试函数 */
func TestGenericitySpell(t *testing.T) {

	var eleDeck Deck[*SpellCard] = []*SpellCard{}
	fmt.Println(len(eleDeck))
	res := &eleDeck
	r := res.Pop()
	fmt.Println(r.Str())
	fmt.Println(len(eleDeck))

	r = res.Pop()
	fmt.Println(r.Str())
}

func TestElementInit(t *testing.T) {
	// 测试元素卡的初始化
	d := ElementCardDeckInit()
	d.Show()

	for i := 0; i < 3; i++ {
		d.Pop()
	}
	d.Show()

	d.Shuffle()
	d.Show()
}

func TestSpellInit(t *testing.T) {
	// 测试元素卡的初始化
	d := SpellCardDeckInit()
	d.Show()

	fmt.Println("----------Pop to empty")
	for len(d.Data) != 0 {
		d.Discard = append(d.Discard, d.Pop())
	}
	d.Show()

	fmt.Println("----------shuffle")
	d.Shuffle()
	d.Show()
}

// 判断能否获取目标符卡的函数测试
func TestCanGain(t *testing.T) {
	// d := Deck{make([]Card, 0)}
	// e := &SpellCard{"Test", [5]int{1, 2, 3, 4, 5}, 2}

}

func TestPointer(t *testing.T) {
	// 用来测试go的指针相关
	a := make([]Card, 0)
	c := &Element{}
	a = append(a, c)
	fmt.Println(a)
}

// 字符串转化相关
func TestString(t *testing.T) {
	config := "01002,10020,20001,12000,01200,00012,20100,00120,02010,00201,00121,21001,12100,01102,10012,11020,20110,02011,01210,10201,11111"
	res := strings.Split(config, ",")
	cost := [5]int{0, 0, 0, 0, 0}
	sc := make([]*SpellCard, 0)

	for idx, v := range res[:] {
		for i := 0; i < 5; i++ {
			// fmt.Printf("%v-%T-", v, v[i])
			a, _ := strconv.Atoi(string(v[i]))
			cost[i] = a
		}

		score := (idx) / 10
		sc = append(sc, &SpellCard{
			string(v),
			cost,
			score,
		})
		fmt.Println(sc[idx])

		for i := 0; i < 5; i++ {
			cost[i] = 0
		}
	}

	// fmt.Println(sc[1], sc[14])
}
