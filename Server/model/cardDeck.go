/* 卡堆抽象类，包括元素，符卡，帮手均可用 */

package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

/* 卡堆，包括卡组，弃牌堆，以及展示的卡片 */
type CardDeck[T CardType] struct {
	Data    Deck[T] `json:"-"`       //`卡组数据`
	Display Deck[T] `json:"display"` // 展示中的卡牌
	Discard Deck[T] `json:"-"`       // 已丢弃的卡牌
}

func (d *CardDeck[T]) MarshalJSON() (encode []byte, err error) {
	return json.Marshal(d.Display)
}

// 初始化，指定各数组的容量
// 然后用指定的卡牌数组cards填充，每种n张，
// display：是否给到展示区
// 为了方便，展示卡牌的数组仍设为切片类型
func (d *CardDeck[T]) Init(cards []T, n int, display bool) {
	// 初始化各个切片
	d.Data = make(Deck[T], n*len(cards))
	d.Display = make(Deck[T], 5)
	d.Discard = make(Deck[T], 0)

	// 填充卡堆
	idx := 0
	for i := 0; i < n; i++ {
		for _, v := range cards {
			d.Data[idx] = v
			idx++
		}
	}

	// 放5张到展示区
	d.Data.RandShuffle()
	if display {
		for i := 0; i < 5; i++ {
			d.Display[i] = d.Pop()
		}
	}
}

// 添加一张卡牌到卡组
func (d *CardDeck[T]) Append(t T) {
	d.Data = append(d.Data, t)
}

// 弹出并返回牌堆顶
func (d *CardDeck[T]) Pop() (res T) {
	if len(d.Data) == 0 {
		d.Shuffle()
	}
	res = d.Data[len(d.Data)-1]
	d.Data = d.Data[:len(d.Data)-1]
	// fmt.Println("Pop :", res)
	return
}

// 获取并返回展示区的一张牌，然后抽一张卡替代
func (d *CardDeck[T]) GetDisplay(idx int) (res T) {
	res = d.Display[idx]
	d.Display[idx] = d.Pop()
	return
}

// 洗牌，将已丢弃的牌放回空牌组，注意仅限于在抽牌堆为空时使用
func (d *CardDeck[T]) Shuffle() {
	d.Data = d.Discard[:]
	d.Discard = make(Deck[T], 0)
	d.Data.RandShuffle()
}

// 展示相关信息
func (d *CardDeck[T]) Show() {
	fmt.Printf("Current Deck' length : %d\n", len(d.Data))
	// fmt.Printf("Top of Deck  : %v\n", d.Data[len(d.Data)-1])
	// fmt.Printf("Deck[-2] : %v\n", d.Data[len(d.Data)-1])
	fmt.Println("Current dispaly : ")
	d.Display.Show()
	fmt.Printf("\nCurrent Discard' length : %d\n", len(d.Discard))
}

/* 初始化元素卡牌堆，每种元素各十张，各五个卡面 */
func ElementCardDeckInit() CardDeck[*Element] {
	ele := make([]*Element, 25)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			ele[i*5+j] = &Element{i, j}
		}
	}

	d := CardDeck[*Element]{}
	d.Init(ele, 2, true)
	// d.Show()
	return d
}

// 符卡卡堆的初始化
func SpellCardDeckInit() (scd CardDeck[*SpellCard]) {
	// 各张卡的cost
	config := "01002,10020,20001,12000,01200,00012,20100,00120,02010,00201,00121,21001,12100,01102,10012,11020,20110,02011,01210,10201,11111"
	res := strings.Split(config, ",")
	cost := [5]int{0, 0, 0, 0, 0}
	sc := make([]*SpellCard, 0)

	// 定义各张符卡
	for idx, v := range res[:] {
		for i := 0; i < 5; i++ {
			// fmt.Printf("%v-%T-", v, v[i])
			a, _ := strconv.Atoi(string(v[i]))
			cost[i] = a
		}

		name := strconv.Itoa(idx)
		if len(name) == 1 { // 补全两位
			name = "0" + name
		}
		score := (idx) / 10

		sc = append(sc, &SpellCard{
			name,
			cost,
			score + 2,
		})
		// fmt.Println(sc[idx])

		for i := 0; i < 5; i++ {
			cost[i] = 0
		}
	}

	sc[20].Score = 5
	scd.Init(sc, 1, true)
	return
}

/* 初始化帮手卡牌堆，*/
func HelperCardDeckInit() CardDeck[*Helper] {
	ele := make([]*Helper, 1)
	ele[0] = &Helper{
		name:   "rumia",
		text:   "",
		CanUse: false,
		score:  0,
	}
	d := CardDeck[*Helper]{}
	d.Init(ele, 20, false)
	// d.Show()
	return d
}
