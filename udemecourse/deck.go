package main

import "fmt"

//创建一个deck类型的切片

type deck []string

// 新增一个newDeck函数，返回一个deck类型的切片
func newDeck() deck {
	//创建一个切片，包含52张扑克牌
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	//遍历花色和点数，组合成一副牌
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// 声明一个函数，作用是遍历打印切片的内容
func print(d deck) {
	for index, deck := range d {
		fmt.Println(index, deck)
	}
}

//模拟发牌 声明一个函数，作用是将切片打乱
/**
 	这里使用了go的一个特性：
	范围：
		d[:handSize], d[handSize:]
		表示从切片d的第0个元素到handSize-1个元素，和从handSize到最后一个元素
		也就是将切片d分成两部分，前handSize个元素和后面的元素
		返回两个切片
		这里的handSize是一个参数，表示发牌的数量
		这里的d是一个切片，表示一副牌
	@desc 返回多个参数值类型的例子
*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
