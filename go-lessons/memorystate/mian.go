package main

// 测试Go 初始堆内存的大小
import "runtime"

var stat runtime.MemStats

func main() {
	runtime.ReadMemStats(&stat)
	println(stat.HeapSys)
}
