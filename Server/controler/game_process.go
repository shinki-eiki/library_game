package controler

import (
	"patchouli_lib/model"
	"strconv"
)

// 游戏的主要循环
func MainLoop() {
	// 初始化游戏内存，输入和输出接口变量
	game := model.NewGame([]string{"Reimu", "Kochiya"}, nil, nil)
	out := &CommandLine{}
	input := out
	// out.showAll(game)
	op := []string{
		"UseHelper",
		"GainSpell",
		"GainHelper",
	}
	var cp *model.Player

	// 进行游戏
	for { //每回合的操作
		game.TurnBegin()
		cp = game.CP
		out.showAll(game)

		// 开始获取元素卡
		// fmt.Println(game.Pv.DrawCardNumber)
		for i := 0; i < game.Pv.DrawCardNumber; i++ {
			if game.Pv.DrawWithSelect {
				out.showOnePlayer(cp, true)
				out.showElementDeck(game.Element.Display)
				temp, _ := input.AskIntegerRange(
					"Select one of the element card", 1, 5)
				game.PlayerGainElement(cp, temp-1)
			} else {
				game.PlayerDrawElement(cp)
			}
		}

		// 然后是剩下的操作
		gaming := true
		for gaming {
			// 打印可用操作
			out.showAll(game)
			rec := [6]bool{} // default is false
			out.newMessage("===You can do the things below:")
			for idx, v := range op {
				if game.Operations.HasAction(model.ActionType(v)) {
					out.newMessage(strconv.Itoa(idx), ":", model.Operations_EXPLAIN[v], "")
					rec[idx] = true
				}
			}
			out.newMessage("5 : End this turn.")

			// 选择操作
			sel, err := input.AskInteger("选择你要进行的操作：")
			if err != nil {
				out.newMessage("Wrong!", err.Error())
				continue
			}
			if sel != 5 && !rec[sel] {
				continue
			}

			// 根据输入进行调用操作
			switch sel {
			case 5:
				gaming = false
			case 0: // use helper
				gaming = false
			case 1: // to gain spell
				{
					out.showSpellDeck(game.Spellcard.Display)
					spellIdx, _ := input.AskIntegerRange(
						"Select one of the spell card,input 0 to escape", 1, 5)
					if spellIdx == 0 {
						continue
					}

					tar := game.Spellcard.Display[spellIdx-1]
					temp := game.ToGainSpell(game.CP, spellIdx-1)
					if temp == -1 {
						continue
					}

					out.newMessage(cp.Name, "获得了", tar.Name)
					if temp == 1 { // 强制丢弃一张卡牌
						if len(cp.Hand.Data) == 1 {
							temp = 0
						} else {
							out.showOnePlayer(cp, true)
							temp, _ = out.AskIntegerRange(
								"You must throw a card", 1, len(cp.Hand.Data)-1)
						}
						cp.Hand.Remove(temp)
					}
				}
			case 2: // to gain helper
				{
					yes := game.ToGainHelper(cp)
					if yes {
						out.newMessage("You gain a helper.")
					} else {
						out.newMessage("Not enough resource.")
					}
				}
			case 3: // koakuma neta
				{
					gaming = false
				}
			}
		}

		game.TurnEnd()
		out.newMessage("******************************\nNext player:")
	}

	// 结束游戏
	// game.Over()
}
