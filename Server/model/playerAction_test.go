package model_test

import (
	"fmt"
	. "patchouli_lib/model"
	"testing"
)

func TestEmptyMap(t *testing.T) {
	// 如果获取一个不存在的键的对应值，将返回默认值
	// 所以需要用if-ok的形式去判断
	m := make(map[int]int)
	fmt.Println(m[114])
	fmt.Println(m[0])
	if v, ok := m[514]; ok {
		fmt.Printf("value is exist,it's %v\n", v)
	} else {
		fmt.Printf("value is not exist\n")
	}
}

func TestPlayerAction(t *testing.T) {
	playerActions := NewPlayerActions()
	playerActions.ShowMap()

	// 尝试记录不同的动作
	err := playerActions.RecordAction(ActionAttack)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Attack recorded.")
	}

	// 再次尝试记录相同的动作，应该失败
	err = playerActions.RecordAction(ActionAttack)
	if err != nil {
		fmt.Println(err)
	}

	// 记录其他动作
	err = playerActions.RecordAction(ActionDefend)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Defend recorded.")
	}

	// 查询动作是否被记录
	fmt.Println("Has attacked:", playerActions.HasAction(ActionAttack))
	fmt.Println("Has healed:", playerActions.HasAction(ActionHeal))

	// 重置所有动作
	playerActions.ResetActions()

	// 再次查询，应该都返回 false
	fmt.Println("Has attacked after reset:", playerActions.HasAction(ActionAttack))
	fmt.Println("Has defended after reset:", playerActions.HasAction(ActionDefend))
}
