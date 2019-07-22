package mutex

import "sync"

func main() {

	// 条件变量
	var cond = sync.NewCond(&sync.Mutex{})
	cond.L.Lock()
	var maxClinks = 30
	gettedClients := 30
	for{
		if gettedClients == maxClinks {
			cond.Wait()
		}
		cond.L.Unlock()
	}
}
