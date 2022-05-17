package task

import (
	"fmt"
	"math/rand"
	"time"
)

func RunTask() {
	var (
		runTime = map[int]bool{
			20: true,
			30: true,
		}
	)

	ticker := time.NewTicker(time.Second * 1)

	for {
		now := <-ticker.C
		nowSecond := now.Second()

		if runTime[nowSecond] {
			go func() {
				fmt.Printf("signing in %d", nowSecond)
			}()
		}
		if nowSecond == 59 {
			runTime = generateRandomSchedule()
			fmt.Printf("重定义计划")
		}
	}
}

func generateRandomSchedule() map[int]bool {
	first := 10 + rand.Intn(10)
	second := 40 + rand.Intn(15)
	return map[int]bool{
		first:  true,
		second: true,
	}
}
