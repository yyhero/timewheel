// 实现单例模式的时间轮 （推荐使用）
package timewheel

import (
	"sync"
	"time"
)

var (
	// 时间轮间隔时间
	interval time.Duration = 1 * time.Second

	// 一个时间轮拥有的槽位个数
	slotNum int = 500

	// 时间轮对象
	tw *TimeWheel

	// 时间轮锁
	once sync.Once
)

// 获取单例对象
func GetInstance() *TimeWheel {
	once.Do(func() {
		tw = New(interval, slotNum)
	})

	return tw
}
