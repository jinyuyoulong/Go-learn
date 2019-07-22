package main

import (
	"fmt"
	"github.com/jakecoffman/cron"
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
	serviceCron = cron.New()
}

func main() {
	task1 := Task{Id: "name1"}
	task2 := Task{Id: "name2"}

	spec1 := "*/2 * * * * *"
	spec2 := "*/3 * * * * *"

	serviceCron.AddFunc(spec1, func1, task1.Id)
	serviceCron.AddJob(spec2, task2, task2.Id)

	serviceCron.Start()

	//打印-即将执行的任务
	go getNextJobs()

	// 阻塞协程 不退出
	select {}
}

func getNextJobs() {
	fmt.Println("next jobs start...")
	entries := getEntries(2)
	for k, v := range entries {
		job := v.Job
		nextTime := v.Next.Unix()
		nextType := fmt.Sprintf("%T", job)
		fmt.Println(nextType)
		if nextType == "main.Task" { // main.Task 报名.变量类型名(结构体名，方法名)
			fmt.Println(k, "job type ", job.(Task).Id, "===>", nextTime)
		} else { // cron.FuncJob
			fmt.Println(k, "job type ", "func", "===>", nextTime)
		}
	}
	fmt.Println("next job end")
}

func getEntries(size int) []*cron.Entry {

	result := serviceCron.Entries()
	if len(result) > size {
		result = result[:size]
	}
	return result
}

func func1() {
	fmt.Println("func1 name1 running")
}
