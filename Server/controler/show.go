/*
CommandLine结构中用来展示游戏界面的方法
*/

package controler

import (
	"fmt"
	"patchouli_lib/model"
	"strconv"
	"strings"
)

// 命令行的输出接口
type CommandLine struct {
}

func NewCommandLine() *CommandLine {
	return &CommandLine{}
}

// var COMMANDLINE = &CommandLine{}

// 将一个5元数组转化为制表符分隔的字符串
func PropertyArrayToString(sum [5]int) string {
	res := make([]string, 5)
	for i := 0; i < 5; i++ {
		res[i] = strconv.Itoa(sum[i])
	}
	return strings.Join(res, "\t")
}

// 5种元素的统计的字符串表示
func PropertyTable(es []*model.Element) string {
	sum := [5]int{0, 0, 0, 0, 0}
	for _, v := range es {
		sum[v.Property]++
	}
	return PropertyArrayToString(sum)
}

// 展示一名玩家，布尔变量代表是否单独显示，若是，将输出分隔符
func (com *CommandLine) showOnePlayer(v *model.Player, sep bool) {
	if sep {
		fmt.Println("---Current Players ")
		fmt.Println("昵称\t金\t木\t水\t火\t土\t帮手\t分数")
	}
	hasHelp := "无"
	if v.Helper != nil {
		hasHelp = "有"
	}
	cnt := PropertyTable(v.Hand.Data)
	// cnt := PropertyTable(v.Hand.Data)
	fmt.Printf("%v\t%v\t%v\t %v\n", v.Name, cnt, hasHelp, v.Score)
	// if sep {
	// 	fmt.Println("------------------------ ")
	// }
}

// 展示所有玩家
func (com *CommandLine) showPlayer(ps []*model.Player) {
	fmt.Println("---Players ")
	fmt.Println("昵称\t金\t木\t水\t火\t土\t帮手\t分数")
	for _, v := range ps {
		com.showOnePlayer(v, false)
	}
	// fmt.Println("------------------------ ")
	// model.PropertyTable()
}

// 展示符卡堆
func (com *CommandLine) showSpellDeck(sd []*model.SpellCard) {
	fmt.Println("---Spell Card ")
	fmt.Println("名称\t金\t木\t水\t火\t土\t分数")
	for _, v := range sd {
		s := PropertyArrayToString(v.Cost)
		fmt.Printf("%v\t%v\t%v\n", v.Name, s, v.Score)
	}
	// fmt.Println("------------------------ ")
}

// 打印展示中的元素卡
func (com *CommandLine) showElementDeck(sd []*model.Element) {
	fmt.Println("---Element Card ")
	fmt.Println("1\t2\t3\t4\t5")
	// fmt.Println("名称\t金\t木\t水\t火\t土\t分数")
	for _, v := range sd {
		fmt.Printf("%v\t", model.PROPERTYMAP[v.Property])
	}
	// fmt.Println("------------------------ ")
	fmt.Println()
}

// 打印所有游戏内容
func (com *CommandLine) showAll(g *model.Game) {
	com.showSpellDeck(g.Spellcard.Display)
	com.showElementDeck(g.Element.Display)
	com.showPlayer(g.Players)
	com.showOnePlayer(g.Players[g.Current], true)
	// com.

	fmt.Println("-------元素卡\t符卡\t帮手卡")
	fmt.Printf("剩余\t%d\t%d\t%d\n", len(g.Element.Data), len(g.Spellcard.Data), len(g.Helpers.Data))
	fmt.Printf("丢弃的\t%d\t%d\t%d\n", len(g.Element.Discard), len(g.Spellcard.Discard), len(g.Helpers.Discard))
	fmt.Println("----当前状态:\n", g.State.CurrentSTring())

}

// 简单地打印到消息框
func (com *CommandLine) newMessage(s ...any) {
	for _, v := range s {
		fmt.Print(v)
	}
	fmt.Println()
}
