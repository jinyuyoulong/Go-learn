package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func (u *User) Hello(name string) string {
	fmt.Printf("hello %s ，my name is %s\n", name, u.Name)
	return u.Name
}
func main() {
	// 使用反射 修改变量的值
	// x := 123
	// v := reflect.ValueOf(&x) // 需要传一个指针
	// v.Elem().SetInt(999)
	// fmt.Println(x)

	u := User{1, "OK", 22}
	// Set(&u)
	setFunction(&u)

	fmt.Println(u)
}

// 修改变量的值
func Set(o interface{}) {
	v := reflect.ValueOf(o)
	// 检查 v 的类型，以及是否支持set 方法
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}

	// f := v.FieldByName("Name")
	f := v.FieldByName("Name1")

	if !f.IsValid() { // 是否找到要修改的属性变量
		fmt.Println("bad!")
		return
	}
	// 根据反射类型，调用对应的set方法修改变量的值
	if f.Kind() == reflect.String {
		f.SetString("baybay")
	}
}

// 动态调用方法
func setFunction(o interface{}) {
	v := reflect.ValueOf(o)
	mv := v.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("joe")}
	username := mv.Call(args) // 调用方法，并得到返回值
	fmt.Println(username)
}
