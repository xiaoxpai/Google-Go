package main

import "fmt"

func main() {
	dillan := person{
		firstName: "Dillan",
		lastName:  "Smith",
		contact: contactInfo{
			email:   "dillan@xiaoxpai.cn",
			zipCode: 100000,
		},
	}
	fmt.Printf("%+v", dillan)
	//{firstName:Dillan lastName:Smith contact:{email:dillan@xiaoxpai.cn zipCode:100000}}
}

/*
  切片用例
	cards := newDeck()
	//发牌：deal（一副牌，发5张牌）
	hand, remainingDeck := deal(cards, 5)

	print(hand)
	print(remainingDeck)

*/

/* 嵌套一个struct结构体用例 */
// 联系方式
type contactInfo struct {
	email   string
	zipCode int
}

// 联系人
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
