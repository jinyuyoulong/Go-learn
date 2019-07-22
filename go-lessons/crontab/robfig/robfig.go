package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
)

type Task struct {
	Id string
}

// 实现cron.Job{}接口
func (t Task) Run() {
	fmt.Println("task job ", t.Id, " running...")
}

var serviceCron *cron.Cron

func init() {
	//serviceCron = cron.New()
}

func main() {

	demo()
}
func mytest() {
	task2 := Task{Id: "name2"}
	//serviceCron.AddJob("@every 1s",task2)

	//spec1 := "0 30 * * * *"
	//spec2 := "*/3 * * * * *"

	//entryID, _ := serviceCron.AddFunc("0 30 * * * *", func1)
	//fmt.Println("entry id ",entryID)
	//
	//jobID, _ := serviceCron.AddJob(spec2, task2)
	//fmt.Println("job id ",jobID)

	//进入| 说明| 相当于
	//----- | ----------- -------------
	//@yearly（或@annually）| 每年一次，午夜，1月1日| 0 0 0 1 1 *
	//@monthly | 每月运行一次，午夜，月初| 0 0 0 1 * *
	//@weekly | 每周一次，午睡/周日午夜| 0 0 0 * * 0
	//@daily（或@midnight）| 每天运行一次，午夜| 0 0 0 * * *
	//@hourly | 每小时运行一次，小时开始| 0 0 * * * *
	//@every <duration> 自定义间隔时间

	serviceCron = cron.New()
	// "* * * * * *" add faild
	serviceCron.AddFunc("* * * * * *", func() {
		fmt.Println("Every hour on the half hour")
	})
	//serviceCron.AddFunc("@hourly",      func() { fmt.Println("Every hour") })
	//serviceCron.AddFunc("@daily", func() { fmt.Println("Every day") })
	//serviceCron.AddFunc("@every 1s", func() { fmt.Println("Every 秒 job run") })
	spec0 := "*/1 * * * * *"
	serviceCron.AddJob(spec0, task2)

	serviceCron.Start()

	//inspect(serviceCron.Entries())

	// 阻塞协程 不退出
	select {}
}
func demo() {
	log.Println("Starting...")

	c := cron.New()
	//h := Task{"I Love You!"}
	//// 添加定时任务
	//c.AddJob("*/2 * * * * * ", h)
	//// 添加定时任务
	//c.AddFunc("*/5 * * * * * ", func() {
	//        log.Println("hello word")
	//    })

	paser := cron.Parser{}
	s, err := paser.Parse("*/3 * * * * *")
	if err != nil {
		log.Println("Parse error")
	}

	h2 := Task{"I Hate You!"}
	c.Schedule(s, h2)
	// 其中任务
	c.Start()
	// 关闭任务
	defer c.Stop()
	select {}
}

// 检查cron作业条目的下一次和上一次运行时间。
func inspect(entries []cron.Entry) {
	//打印-即将执行的任务
	go getNextJobs()
}

func getNextJobs() {
	fmt.Println("next jobs start...")
	entries := serviceCron.Entries() //getEntries(2)
	fmt.Println(entries)
	for k, v := range entries {
		job := v.Job
		nextTime := v.Next.Unix()
		nextType := fmt.Sprintf("%T", job)
		fmt.Println(nextType)
		if nextType == "main.Task" { // main.Task 报名.变量类型名(结构体名，方法名)
			fmt.Println(k, "job type ", job.(Task).Id, "===>", nextTime)
		} else if nextType == "cron.FuncJob" {
			fmt.Println(k, "job type ", "func", "===>", nextTime)
		} else {
			fmt.Println("不知道什么类型的job")
		}
	}
	fmt.Println("next job end")
}

func getEntries(size int) []cron.Entry {

	result := serviceCron.Entries()
	if len(result) > size {
		result = result[:size]
	}
	return result
}

func func1() {
	fmt.Println("func1 name1 running")
}
