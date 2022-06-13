package task

import (
	"fmt"
	"time"
)

func Schedule() {
	// 每天7时0分触发更新
	t := time.NewTimer(SetTime(0, 51, 0))
	defer t.Stop()
	for {
		select {
		case <-t.C:
			t.Reset(time.Hour * 24)
			// 定时任务函数
			RunTask()
		}
	}
}

func SetTime(hour, min, second int) (d time.Duration) {
	now := time.Now()
	setTime := time.Date(now.Year(), now.Month(), now.Day(), hour, min, second, 0, now.Location())
	d = setTime.Sub(now)
	if d > 0 {
		return
	}
	return d + time.Hour*24
}

func RunTask() {
	fmt.Println("命中定时任务")
}
