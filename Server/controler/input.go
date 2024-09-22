package controler

import "patchouli_lib/model"

// 用来测试命令行输入的类
type Test_class struct {
	com *CommandLine
}

func (t *Test_class) AskElement() (res int) {
	res, _ = t.com.AskIntegerRange("Please input the index of element card you want:", 1, 5)
	return
}

func (t *Test_class) AskAction(n int) (res int) {
	res, _ = t.com.AskIntegerRange("Please input the action you want:", 1, n)
	return
}

func (t *Test_class) AskSpell() (res int) {
	res, _ = t.com.AskIntegerRange("Please input the index of spell card you want:", 1, 5)
	return
}

func (t *Test_class) AskHand(n int) (res int) {
	res, _ = t.com.AskIntegerRange("Please input the index of spell card you want:", 1, n)
	return
}

func (t *Test_class) AskHelper() (res int) {
	res, _ = t.com.AskIntegerRange("Please select the index of helper you want to remain:", 1, 2)
	return
}

func (t *Test_class) AskResponse() (res int) {
	return
}

func (t *Test_class) AskPlayer() (res int) {
	return
}

var CommandInput model.GameInput = &Test_class{}
