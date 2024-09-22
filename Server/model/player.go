package model

import "fmt"

type Player struct {
	Name   string       // 玩家昵称
	Hand   *HandCard    //`json:"-"` //手卡列表
	Helper *Helper      `json:"-"` //帮手卡
	Spells []*SpellCard `json:"-"`
	Score  int          `json:"-"` //当前分数
}

// 添加一张手牌
// func (p *Player) HandAppend(e *Element) {
// 	p.Hand.Append(e)
// }

// 丢弃一张手牌
func (p *Player) HandRemove(i int) {
	p.Hand.Remove(i)
}

// 获得一张元素牌，并加入手卡
func (p *Player) GetElement(e *Element) {
	p.Hand.Append(e)
}

// 玩家获得目标符卡
func (p *Player) GetSepllCard(s *SpellCard) int {
	// p.Hand.PayFor(s)
	p.Spells = append(p.Spells, s)
	p.Score += s.Score
	return p.CountScore()
}

// 获得帮手
func (p *Player) GetHelper(h *Helper) int {
	if p.Helper == nil {
		p.Helper = h
	} else { // 如果已有帮手卡，返回-1让上面处理
		return -1
	}
	return p.CountScore()
}

// 使用帮手
func (p *Player) UseHelper(g *Game) int {
	if !p.Helper.CanUse {
		return -1
	}
	// 使用帮手

	return p.CountScore()
}

// 计算分数然后返回
func (p *Player) CountScore() int {
	if p.Helper != nil {
		return p.Score + p.Helper.score
	}
	return p.Score
}

// 回合结束的处理
func (p *Player) End() {
	h := p.Helper
	if h != nil {
		h.CanUse = true
	}
}

// 回合开始的处理，处理一些持续效果
func (p *Player) Begin() {
	h := p.Helper
	if h != nil {
		h.CanUse = true
	}
}

// 根据人数创建玩家对象数组，然后随机选先
func NewPlayers(names []string) (res []*Player) {
	for _, v := range names {
		res = append(res, &Player{v, &HandCard{}, nil, make([]*SpellCard, 0), 0})
	}
	return
}

// 打印玩家信息
func (p *Player) Show() {
	fmt.Println(">>>Player", p.Name)
	// fmt.Println("Hand :")
	p.Hand.Show()
	fmt.Println("Has helper?", p.Helper != nil)
	fmt.Println("His score is ", p.Score)

}
