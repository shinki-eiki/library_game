package model

type Card interface {
	GetImage() string // 获取卡面文件位置
	Str() string      // 简介字符串
}

// interface Card{

// }
