package main
import "testing"

func TestSum(t *testing.T)  {
	sum := Sum(4)
	if sum != 5 {
		t.Errorf("测试失败")
		return
	}
	t.Log("测试成功")
}