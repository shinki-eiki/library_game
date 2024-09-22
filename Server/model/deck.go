package model

import (
	"fmt"
	"math/rand"
	"time"
)

type CardType interface {
	*Element | *SpellCard | *Helper
}

/* 卡片的集合 */
type Deck[T CardType] []T

// 打印展示每一张牌
func (deck Deck[T]) Show() {
	fmt.Println("Length of this deck :", len(deck))
	for _, v := range deck {
		// var c Card
		c := Card(v)
		fmt.Print(c.Str(), ",")
	}
}

/* 从卡堆最上面拿走一张牌，这种写法可以修改pop掉底层的切片 */
func (deck *Deck[T]) Pop() (res T) {
	d := *deck
	res = d[len(d)-1]
	*deck = d[:len(d)-1]
	return
}

func (slice Deck[T]) RandShuffle() {
	/* 将Deck打乱 */
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
