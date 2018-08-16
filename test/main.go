package main

func main() {
	// 构造player
	playerObj1 := NewPlaeyr("num1")
	playerObj2 := NewPlaeyr("num2")

	// 玩家创建房屋
	playerObj1.ConstructHouse()

	// 玩家创建士兵
	playerObj2.ConstructSoldier()

	select {}
}
