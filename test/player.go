package main

import (
	"fmt"
	"github.com/yy/timewheel"
	"time"
)

type Player struct {
	id string

	// 房子数量
	houseNum int

	// 士兵数量
	soldierNum int
}

// 创建房屋
func (this *Player) ConstructHouse() {
	fmt.Printf("玩家:%v开始创建房屋,cur houseNum:%v\n", this.id, this.houseNum)

	// 添加定时器,2s之后创建完成
	timewheel.GetInstance().AddTimer(2*time.Second, "001", func() {
		this.houseNum += 50
		fmt.Printf("玩家:%v创建房屋完成,cur houseNum:%v\n", this.id, this.houseNum)
	})
}

// 创建士兵
func (this *Player) ConstructSoldier() {
	fmt.Printf("玩家:%v开始创建房屋,cur soldierNum:%v\n", this.id, this.soldierNum)

	// 添加定时器,5s之后创建完成
	timewheel.GetInstance().AddTimer(5*time.Second, "001", func() {
		this.soldierNum += 20
		fmt.Printf("玩家:%v创建房屋完成,cur soldierNum:%v\n", this.id, this.soldierNum)
	})
}

func NewPlaeyr(id_ string) *Player {
	return &Player{
		id: id_,
	}
}
