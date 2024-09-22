package model

import (
	"encoding/json"
	"errors"
	"fmt"
)

// 手牌
type HandCard struct {
	Data []*Element
}

func (d *HandCard) MarshalJSON() (encode []byte, err error) {
	return json.Marshal(d.Data)
}

/* 判断资源是否足够获取卡牌 */
func (h *HandCard) CanGain(s *SpellCard) bool {
	cnt := s.Cost
	for _, v := range h.Data {
		cnt[v.Property]--
	}
	for _, v := range cnt {
		if v > 0 {
			return false
		}
	}
	return true
}

// 判断能否获得帮手
func (h *HandCard) CanGainHelper() bool {
	cnt := 0
	for _, v := range h.ArraySum() {
		if v != 0 {
			cnt++
		}
	}
	return cnt > 1
}

/* 假设资源足够的情况下，为一张符卡消耗资源 */
func (h *HandCard) PayFor(s *SpellCard) (rem []int) {
	cnt := s.Cost
	already := 0
	rem = make([]int, 0)
	// 找到最靠前的满足条件的元素卡，得到一个索引数组，
	// 注意这个数组是处理过的，可以直接用来切片的
	for idx, v := range h.Data {
		p := v.Property
		if cnt[p] != 0 {
			rem = append(rem, idx-already)
			already++
			cnt[p]--
		}
	}

	for _, v := range rem {
		h.Remove(v)
	}
	h.Show()
	return
}

// 丢弃指定索引的卡牌
func (h *HandCard) Remove(i int) (nil error) {
	if len(h.Data) <= i {
		return errors.New("index over")
	}
	h.Data = append(h.Data[:i], h.Data[i+1:]...)
	return
}

// 添加一张手牌
func (h *HandCard) Append(e *Element) {
	h.Data = append(h.Data, e)
}

// 手牌统计数组
func (h *HandCard) ArraySum() [5]int {
	sum := [5]int{0, 0, 0, 0, 0}
	for _, v := range h.Data {
		sum[v.Property]++
	}
	return sum
}

// 展示手牌细节
func (h *HandCard) Show() {
	fmt.Print("Current hand:")
	if len(h.Data) == 0 {
		fmt.Println("Empty.")
		return
	}
	for _, v := range h.Data {
		fmt.Print(v.Str(), ",")
	}
	fmt.Println()
}

// 打印列举五种元素的数量
func PropertyTable(es []*Element) string {
	fmt.Println("[  金 木 水 火 土  ]")
	sum := [5]int{0, 0, 0, 0, 0}
	for _, v := range es {
		sum[v.Property]++
	}
	fmt.Print("[  ")
	for _, v := range sum {
		fmt.Print(v, "  ")
	}
	fmt.Print("]")
	return ""
}

// 包括增删改查？
//
//
//
