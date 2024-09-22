// 包括游戏进行所需要的输入和输出接口
package model

// 1.选择一张元素牌获取
// 2.选择一项行动，包括：获取帮手，获取符卡，使用状态效果，使用帮手这四种之一
// 然后就是上面的拓展，
// 3.选择一张帮手获取
// 4.选择一张符卡获取
// 5.选择一张手卡
// 6.使用帮手，当只有一张帮手时，无需选择
// 考虑到
// 7.是否发动帮手的触发效果
// 8.从两张帮手之间选一张使用或丢弃
// 9.选择一名（其他）玩家

// 游戏需要输入时的相应方法
type GameInput interface {
	// 询问一张元素牌的编号（1-5之间）
	AskElement() int
	// 询问下一步动作
	AskAction(n int) int
	// 询问一张符卡的编号（1-5之间）
	AskSpell() int
	// 询问一张手卡的编号
	AskHand(n int) int
	// 询问两张帮手之间使用哪一张
	AskHelper() int
	// 询问是否响应使用帮手
	AskResponse() int
	// 选择一位玩家
	AskPlayer() int
}

// 游戏事件发生时的相应动作
type GameOuput interface {
	// 1. 将第idx张元素牌改为target
	ChangeElement(idx int, target *Element) int

	// 2. 将第idx张符卡改为target
	ChangeSpell(idx int, target *SpellCard) int

	// 3. 默认当前玩家获得一张手卡
	appendHand(e *Element) int

	// 4. 第index位玩家失去第idx张手牌
	LoseHandCard(index int, idx int) int

	// 5.当前玩家发生改变
	ChangePlayer() int

	// 6.全局状态发生变化
	ChangeState()

	// 7. 第idx位玩家失去失去其帮手
	UseHelper(idx int) int

	// ChangeAction(n int) int
	// ChangeResponse() int
}
