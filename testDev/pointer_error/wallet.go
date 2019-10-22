package main

import "fmt"

// 使用类型别名具有可描述性
type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposite(amount Bitcoin) {
	w.balance = amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

// 让我们实现 Bitcoin 的 Stringer 方法
// 这个接口是在 fmt 包中定义的。当使用 %s 打印格式化的字符串时，你可以定义此类型的打印方式。
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
