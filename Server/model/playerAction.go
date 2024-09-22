package model

import (
	"fmt"
)

// 定义玩家可以执行的操作类型
type ActionType string

const (
	ActionGainSpell  ActionType = "gainSpell"
	ActionGainHelper ActionType = "gainhelper"
	ActionUseHelper  ActionType = "useHelper"
	ActionNeta       ActionType = "neta"

	ActionAttack ActionType = "attack"
	ActionDefend ActionType = "defend"
	ActionHeal   ActionType = "heal"
	// ActionMove    ActionType = "move"
	// ActionUseItem ActionType = "use_item"
	// ActionSpecial ActionType = "special"
)

// PlayerActions 结构体用于跟踪玩家在当前回合的动作状态
type PlayerActions struct {
	actionsMap map[ActionType]bool
	// mu         sync.Mutex // 互斥锁保证线程安全
}

// 打印当前集合数据
func (pa *PlayerActions) ShowMap() {
	fmt.Println(pa.actionsMap)
	for _, v := range pa.actionsMap {
		fmt.Println(v)
	}
}

// NewPlayerActions 创建一个新的 PlayerActions 实例
func NewPlayerActions() *PlayerActions {
	return &PlayerActions{
		actionsMap: make(map[ActionType]bool),
	}
}

// RecordAction 记录玩家执行的操作，同一操作在同一回合只能执行一次
func (pa *PlayerActions) RecordAction(action ActionType) error {
	if pa.actionsMap[action] {
		return fmt.Errorf("action %s has already been performed this turn", action)
	}

	pa.actionsMap[action] = true
	return nil
}

// HasAction 查询玩家是否已经执行过指定的操作
func (pa *PlayerActions) HasAction(action ActionType) bool {
	return pa.actionsMap[action]
}

// ResetActions 重置所有动作，准备下一轮游戏
func (pa *PlayerActions) ResetActions() {
	for k := range pa.actionsMap {
		pa.actionsMap[k] = false
	}
}
