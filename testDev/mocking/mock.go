package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 依赖注入 模拟
const timeCount = 3
const finalWorld = "GO!"

func CountDown(w io.Writer, sleeper Sleeper) {
	for i := 0; i < timeCount; i++ {
		sleeper.Sleep()
	}
	for i := timeCount; i > 0; i-- {
		fmt.Fprintln(w, timeCount)
	}

	sleeper.Sleep()
	fmt.Fprint(w, finalWorld)
}

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// 为什么要创建这个struct
type ConfigurableSleeper struct {
	duration time.Duration
}

func (t *ConfigurableSleeper) Sleep() {
	time.Sleep(t.duration)
}

func main() {
	spies := &ConfigurableSleeper{1 * time.Second}

	CountDown(os.Stdout, spies)
}
