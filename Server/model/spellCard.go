package model

import "fmt"

/* 符卡 */
type SpellCard struct {
	Name  string
	Cost  [5]int `json:"-"`
	Score int
}

func (e *SpellCard) GetImage() string {
	return e.Name
	// return strconv.Itoa(e.Cost) + "_" + strconv.Itoa(e.Score)
}

func (e *SpellCard) Str() string {
	res := fmt.Sprintf("%s(%d)", e.Name, e.Score)
	return res
}

// 总的资源消耗数量
func (e *SpellCard) TotalCost() int {
	res := 0
	for _, v := range e.Cost {
		res += v
	}
	return res
}
