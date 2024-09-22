package model

import (
	"math/rand"
)

// const (
// 	normal = iota
// 	neta
// 	tired
// 	ill
// 	gatekeeper
// 	help
// )

var STATE_EXPLAIN = map[int]string{
	0: "平稳无事：没事。",
	1: "小恶魔「整理整顿」：在自己的回合内，可以舍弃自己持有的一张属性卡，并从场上取得一张。一回合只能进行一次。",
	2: "分身：当自己的回合开始，要取得属性卡的时候，改为取得２张。",
	3: "生病：每当玩家入手符卡的时候，必须追加舍弃一张属性卡。（种类随意）",
	4: "门卫：不能使用帮手卡。（但是可以入手）",
	5: "助人为乐：在回合內取得帮手卡的玩家，当回合结束时，从属性卡牌库顶抽一张属性卡入手。",
}

/* 游戏的全局状态 */
type State struct {
	CurrentState int
	*PublicValue // 游戏中变量集合的指针
}

// 切换当前状态外的状态，随机变化，返回新状态值
func (s *State) Change() int {
	ns := rand.Intn(6)
	for ns == s.CurrentState {
		ns = rand.Intn(6)
	}
	s.Eliminate()
	s.Flash(ns)
	return ns
}

// 更换状态前，撤销前一状态，但不修改状态的值
// 注意更换状态必须要设置界面，就挺麻烦的
// 都需要更换状态，此外只有状态4需要更改状态框
func (s *State) Eliminate() {
	switch s.CurrentState {
	case 0:
	case 1:
	case 2:
		s.DrawCardNumber = 1
	case 3:
		s.NeedDiscardWhenGain = false
	case 4:
		// 需要修改帮手框的可用与否
		s.CanUseHelper = true
	case 5:
		s.DrawAfterGetHepler = false
	}
}

// 更换状态后，设置新状态，修改状态值
func (s *State) Flash(ns int) {
	(*s).CurrentState = ns
	switch ns {
	case 0:
	case 1:
	case 2:
		s.DrawCardNumber = 2
	case 3:
		s.NeedDiscardWhenGain = true
	case 4:
		// 需要修改帮手框的可用与否
		s.CanUseHelper = false
	case 5:
		s.DrawAfterGetHepler = true
	}
}

// 返回表示当前状态的字符串
func (s *State) CurrentSTring() string {
	return STATE_EXPLAIN[s.CurrentState]
}
