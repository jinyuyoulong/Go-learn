package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Person Person
	ID     int
	Name   string
	Age    int
}

type Person struct {
	Name  string
	Sizes []string
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

	u := User{Person{
		"xiao", []string{"1", "2"},
	},
		1, "OK", 22}
	baseFunc(u)

	// // Set(&u)
	// setFunction(&u)
	// fmt.Println(u)
}

func baseFunc(o interface{}) {
	// 反射的基本操作
	// 1. 获取数据类型
	oType := reflect.TypeOf(o)
	// 1.1 获取对象属性
	f := oType.Field(0)
	// 1.1.1 获取属性 Name
	propertyName := f.Name
	fmt.Printf("%s\n", propertyName)

	// 2. 获取对象的值
	vType := reflect.ValueOf(o)

	// 2.1 获取属性值的 自定义数据类型
	valueType := vType.FieldByName(propertyName)
	fmt.Println(valueType.Interface())

	// 2.2 获取属性值
	propertyValue := vType.Field(0).Interface()
	fmt.Printf("field: %T, %v\n", propertyValue, propertyValue)
	// fmt.Printf("%v", Person(propertyValue).Sizes)

	// 2.3 获取子属性值 取 第0个属性的第一个属性值
	// personValue := vType.FieldByIndex([]int{0, 1}).Interface()
	person := vType.Field(0).Interface()

	// 3 类型断言，转换类型 interface{} ——> Person
	mperson := person.(Person)

	fmt.Printf("%T, %v", mperson, mperson.Sizes[1])
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
