package model

// 公共变量
type PublicValue struct {
	DrawCardNumber      int  // 抽卡数量，对应抽两张的状态
	CanUseHelper        bool // 能否使用帮手，对应状态效果
	NeedDiscardWhenGain bool
	DrawAfterGetHepler  bool
	DrawWithSelect      bool // 是否挑选元素卡上手，对应露米娅效果
}

func (pv *PublicValue) Init() {
	pv.DrawCardNumber = 1  // 抽卡数量，对应抽两张的状态
	pv.CanUseHelper = true // 能否使用帮手，对应状态效果
	pv.NeedDiscardWhenGain = false
	pv.DrawAfterGetHepler = false
	pv.DrawWithSelect = true // 是否挑选元素卡上手，对应露米娅效果
}

// 全局玩家，元素，符卡，帮手的变量
// var (
// 	Players       []*Player
// 	ElementDeck   CardDeck[*Element]
// 	SpellcardDeck CardDeck[*SpellCard]
// 	HelperDeck    CardDeck[*Helper]
// )
