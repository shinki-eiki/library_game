package model

import (
	"fmt"
	"patchouli_lib/util"

	// "encoding/json"
	// "fmt"
	"time"
)

// 操作及其对应说明的集合
var Operations_EXPLAIN = map[string]string{
	"UseHelper":  "使用帮手",
	"GainSpell":  "获得符卡",
	"GainHelper": "获得帮手",
	"Neta":       "弃一抽一",
}

// 一个游戏的所有内存变量的集合
type Game struct {
	Players   []*Player            `json:"players"`
	CP        *Player              `json:"-"` // 当前玩家
	Element   CardDeck[*Element]   `json:"element"`
	Spellcard CardDeck[*SpellCard] `json:"spell"`
	Helpers   CardDeck[*Helper]    `json:"-"`

	Round       int       `json:"-"`       // 游戏轮数
	Current     int       `json:"current"` // 当前玩家编号
	TotalPlayer int       `json:"-"`       // 玩家总数
	BeginTime   time.Time `json:"-"`       // 游戏开始时间

	State *State       `json:"-"`
	Pv    *PublicValue `json:"-"`
	// Operations    map[string]bool `json:"-"` // 操作记录
	Operations *PlayerActions `json:"-"`

	// Inputer  GameInput `json:"-"` // 获取输入
	// Outputer GameOuput `json:"-"` // 发送游戏事件
	ToHub    chan []byte `json:"-"`
	FromHub  chan []byte `json:"-"`
	Selected chan int    `json:"-"` // 发送玩家在事件中的选择

	Begined bool `json:"-"`
}

// 单纯用来传输json信息的结构体
type GameInfo struct {
	*Game
	LocalPlayer int
}

// 创建游戏，同时设置其管道
func NewGame(names []string, to chan []byte, from chan []byte) (g *Game) {
	g = &Game{}
	g.ToHub = to
	g.FromHub = from
	fmt.Printf("[Game]ToHub:%v,FromHub:%v\n", to, from)
	g.Selected = make(chan int)
	// 初始化玩家
	g.Players = NewPlayers(names)

	// 初始化公共卡堆
	g.Pv = &PublicValue{}
	g.Pv.Init()
	g.Element = ElementCardDeckInit()
	g.Spellcard = SpellCardDeckInit()
	g.Helpers = HelperCardDeckInit()
	g.State = &State{0, g.Pv}

	// 初始化非集合的游戏变量
	g.Round = 0
	g.Current = 0
	g.Begined = false
	g.BeginTime = time.Now()
	g.TotalPlayer = len(names)
	g.Operations = NewPlayerActions()
	g.CP = g.Players[g.Current]

	// 分发手牌
	for _, pl := range g.Players {
		for i := 0; i < 3; i++ {
			// pl.GetElement(g.Element.Pop())
			g.PlayerDrawElement(pl)
		}
	}
	return
}

// 每回合的第一步：获取元素
func (g *Game) SelectElement() {
	// 露米娅效果
	for i := 0; i < g.Pv.DrawCardNumber; i++ {
		g.ToHub <- []byte{'-', byte(g.Current), 'a', '2'}
		mes := <-g.FromHub // mes is like a02
		idx := 0
		for i := 1; i < len(mes); i++ {
			idx *= 10
			idx += int(mes[i] - 48)
		}
		fmt.Println("Get Element", idx, g.Element.Display[idx])
		g.PlayerGainElement(g.CP, idx)
	}
}

// 不断接受Hub的消息并处理
func (g *Game) MessageHandle() {
	var mes []byte
	for {
		fmt.Println("[Game] Waiting...")
		mes = <-g.FromHub
		fmt.Println("[Game] receive from Hub : ", mes)
		switch mes[0] {
		case '0': // 当前玩家获取元素
			g.PlayerGainElement(g.CP, int(mes[1])-48)
		case '1': // 当前玩家获取符卡
			g.ToGainSpell(g.CP, int(mes[1])-48)
		case '5': // 当前玩家结束回合
			g.TurnEnd()
		case 'a': // 处理前端所选择的索引，送到专门的管道中
			sel := int(mes[1] - 48)
			if len(mes) == 3 {
				sel *= 10
				sel += int(mes[2] - 48)
			}
			g.Selected <- sel
		default:
			fmt.Println("No this method!")
		}
		fmt.Println("[Game] Handling finish.")
	}
}

// 游戏开始，让游戏开始接受前端的信息，然后让首位玩家开始选择
func (g *Game) Begin() {
	// 传递给每位玩家其编号
	for idx := range g.Players {
		g.ToHub <- util.ConvertToBytes('-', byte(idx), 'c', idx)
	}

	// 让首位玩家选择某一元素
	g.SelectElement()
	if !g.Begined {
		go g.MessageHandle()
	}
}

// 玩家p抽取牌堆顶一张元素
func (g *Game) PlayerDrawElement(p *Player) {
	ele := g.Element.Pop()
	p.GetElement(ele)
	g.ToHub <- util.ConvertToBytes('d', ele.Property, ele.Index)
}

// 玩家p获取第idx张元素
func (g *Game) PlayerGainElement(p *Player, idx int) {
	// g.Element.Show()
	p.GetElement(g.Element.GetDisplay(idx))
	top := g.Element.Display[idx]
	mes := util.ConvertToBytes(0, idx, top.Property, top.Index)
	g.ToHub <- mes
}

//===========需要记录的操作集合

// 判断玩家p能否获得了第idx张符卡，然后执行获取操作
func (g *Game) ToGainSpell(p *Player, idx int) (res int) {
	if g.Operations.HasAction(ActionGainSpell) {
		fmt.Println()
		return
	}
	tar := g.Spellcard.Display[idx]
	need := tar.TotalCost()

	// 是否需要额外丢弃手牌，如果需要，则判断是否还有多出的一张
	// 同时判断是否满足获取的条件
	if (g.Pv.NeedDiscardWhenGain && len(p.Hand.Data) <= need) || !p.Hand.CanGain(tar) {
		return -1
	}
	fmt.Println(g.Current, "player :Satisify.")

	// 获取丢弃的手牌，并逐张通知丢弃
	dis := g.CP.Hand.PayFor(tar)
	for _, v := range dis {
		dis := util.ConvertToBytes('b', v)
		fmt.Println("will discard ", dis)
		g.ToHub <- dis
	}

	// 玩家获取符卡，然后获取获得符卡后的新分数，判断是否要结束游戏
	newScore := p.GetSepllCard(g.Spellcard.GetDisplay(idx))
	if newScore >= 10 {
		fmt.Println("Game over.")
	}

	// 需要额外丢弃的操作流程
	if g.Pv.NeedDiscardWhenGain {
		g.ToHub <- []byte{'-', byte(g.Current), 'a', '1'}
		sel := <-g.FromHub
		discard := util.ConvertBytesToInt(sel)
		fmt.Println("will discard", discard)
		g.ToHub <- util.ConvertToBytes('b', discard)
	}

	//获取后处理，向所有人发送获取了对应符卡的信息
	newState := g.State.Change()
	newSpell := g.Spellcard.Display[idx]
	mes := util.ConvertToBytes(1, idx, newSpell.Name, newSpell.Score, newState)
	g.ToHub <- mes
	g.Operations.RecordAction(ActionGainSpell)
	return
}

// 判断玩家p能否获得帮手
func (g *Game) ToGainHelper(p *Player) bool {
	if !p.Hand.CanGainHelper() {
		return false
	}
	g.PlayerGainHelper(p)
	return true
}

// 玩家p获得了一张帮手的处理
func (g *Game) PlayerGainHelper(p *Player) {
	p.GetHelper(g.Helpers.Pop())
	g.Operations.RecordAction(ActionGainHelper)
}

// 玩家p使用帮手
func (g *Game) PlayerUseHelper(p *Player) {
	// 帮手或者场地效果限制不能使用
	if !g.Pv.CanUseHelper || !p.Helper.CanUse {
		return
	}

	p.UseHelper(g)
	if p == g.CP {
		g.Operations.RecordAction(ActionUseHelper)
	}
}

// 整顿状态效果的操作流程，参数discard和idx分别代表丢弃的手牌和将要获得的元素卡的序号
func (g *Game) Neta(discard int, idx int) {
	g.ToHub <- util.ConvertToBytes('b', discard)
	g.CP.Hand.Remove(discard)
	g.PlayerGainElement(g.CP, idx)
}

// 回合开始的处理
func (g *Game) TurnBegin() {
	g.Operations.ResetActions()
	g.CP = g.Players[g.Current]
	g.SelectElement()
}

// 回合结束的处理
func (g *Game) TurnEnd() {
	p := g.CP
	// 将新获得的帮手设置为可用，然后看看状态是否让抽卡
	h := p.Helper
	if h != nil && !h.CanUse {
		h.CanUse = true
		if g.Pv.DrawAfterGetHepler {
			g.PlayerDrawElement(p)
		}
	}

	// ，并询问
	g.Current++
	if g.Current == g.TotalPlayer {
		g.Current = 0
		g.Round++
	}

	g.ToHub <- util.ConvertToBytes(5)
	g.TurnBegin()
}

// 游戏结束的处理
func (g *Game) Over() {

}

// 废弃代码
// 可行的操作数组
// func (g *Game) WhatCanDo() {
// 	for key,v := range g.Operations {
// 		if v{
// 		}
// 	}
// }

// 自定义game的json编码方法
// func (g *Game) MarshalJSON() (encode []byte, err error) {
// 	res := make(map[string]string)
// 	// res := make(map[string][]byte)
// 	var cur []byte
// 	cur, err = json.Marshal(g.Element.Display)
// 	fmt.Println("Json string", string(cur))
// 	res["element"] = string(cur)
// 	// res["element"] = cur
// 	encode, err = json.Marshal(res)
// 	return
// }

// func (g *Game) Endoce() (res map[string]any) {
// 	res = make(map[string]any)
// 	s, err := json.Marshal(g)
// 	fmt.Println(s, err)
// 	// json.MarshalIndent()
// 	return
// }
