package timewheel

import (
	"container/list"
	"time"
)

// 时间轮
type TimeWheel struct {
	// 时间轮间隔时间
	interval time.Duration

	// 槽数量
	slotNum int

	// 时间轮槽
	slots []*list.List

	// 当前时间轮运行到哪一个槽
	currentPos int

	// 停止定时器channel
	stopChannel chan bool
}

// 延时任务
type Task struct {
	// 延迟时间
	delay time.Duration

	// 时间轮需要转动几圈
	circle int

	// 定时器唯一标识
	identifer interface{}

	// 回调函数
	callback func()
}

// 创建时间轮
func New(interval time.Duration, slotNum int) *TimeWheel {
	if interval <= 0 || slotNum <= 0 {
		print("err: params invalid!\n")
		return nil
	}
	obj := &TimeWheel{
		interval:    interval,
		slotNum:     slotNum,
		slots:       make([]*list.List, slotNum),
		currentPos:  0,
		stopChannel: make(chan bool),
	}

	// 初始化槽，每个槽指向一个双向链表
	for i := 0; i < obj.slotNum; i++ {
		obj.slots[i] = list.New()
	}

	// 启动时间轮
	obj.run()

	return obj
}

// 启动时间轮
func (this *TimeWheel) run() {
	// 创建一个定时器
	ticker := time.NewTicker(this.interval)

	// 循环遍历
	go func() {
		for {
			select {
			case <-ticker.C:
				this.handleTick()
			case <-this.stopChannel:
				ticker.Stop()
				return
			}
		}
	}()
}

// 停止时间轮
func (this *TimeWheel) Stop() {
	this.stopChannel <- true
}

// 添加定时器
func (this *TimeWheel) AddTimer(delay time.Duration, identifer interface{}, callbackFunc func()) {
	if delay <= 0 || identifer == nil {
		return
	}

	// 添加到链表中
	pos, circle := this.getTimerInfo(delay)
	task := &Task{
		delay:     delay,
		circle:    circle,
		identifer: identifer,
		callback:  callbackFunc}

	this.slots[pos].PushBack(task)
}

// 定时处理
func (this *TimeWheel) handleTick() {
	curList := this.slots[this.currentPos]
	this.execute(curList)
	if this.currentPos == this.slotNum-1 {
		this.currentPos = 0
	} else {
		this.currentPos++
	}
}

// 遍历链表, 执行回调函数
func (this *TimeWheel) execute(l *list.List) {
	for e := l.Front(); e != nil; {
		task := e.Value.(*Task)
		if task.circle > 0 {
			task.circle--
			e = e.Next()
			continue
		}

		// 执行回调函数
		go task.callback()

		// 移除过期定时器
		next := e.Next()
		l.Remove(e)
		e = next
	}
}

// 获取定时器在槽中的位置, 时间轮需要转动的圈数
func (this *TimeWheel) getTimerInfo(d time.Duration) (pos int, circle int) {
	delaySeconds := int(d.Seconds())
	intervalSeconds := int(this.interval.Seconds())
	circle = int(delaySeconds / intervalSeconds / this.slotNum)
	pos = int(this.currentPos+delaySeconds/intervalSeconds) % this.slotNum

	return
}
