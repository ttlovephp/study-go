package main

import (
	"fmt"
	"reflect"
	"study/lib"
)

func main() {
	//函数
	lib.Add(1, 1)
	//方法
	p := lib.Person{
		Name: "zhangsan",
		Age:  24,
	}

	pT := reflect.TypeOf(p)
	fmt.Println("-----pT-----", pT)
	fmt.Println(p.GetName())
	//值类型接受
	p.ModifyName1("maoyou1")
	fmt.Println(p.GetName())

	//引用型接受
	p.ModifyName2("maoyou2")
	fmt.Println(p.GetName())
	pT1 := reflect.TypeOf(p)
	fmt.Println("-----pT1-----", pT1)
	//引用型接受 显示指针调用
	(&p).ModifyName2("maoyou3")
	fmt.Println((&p).GetName())

	//什么情况用值类型，什么情况应该用引用类型呢？
	//指针类型接受者传递的是一个指向原指针的副本，指针的副本指向的还是原来的类型，所以修改时，还是会影响原来类型变量的值

	//在调用方法的时候，传递的接收者本质上都是副本，只不过一个是这个值的副本，一个是指向这个值指针的副本，
	//指针具有指向原有值的特性，所以修改了指针指向的值，也就修改了原油值，我们可以简单理解为值接受者使用的是值的副本来调用方法，
	//而指针接收者使用实际的值来调用方法

	//总结：
	//不管是使用值接收者，还是使用指针接收者，一定要搞清楚类型本质，对类型进行操作的时候，是要改变当前值，还是要创建新值进行返回，这些
	//就可以决定我们是使用值传递，还是指针传递，简单理解为如果你是对数据的拷贝就进行值传递，如果对数据的改变就用指针,但是针对大型变量
	//传递指针是卓有成效的，为了保证代码的一致性，尽量都统一使用指针类型

	// time.Tick接受的是chan 可以做定时循环跟，time sleep类似，但是实现比sleep优雅
	// var i int
	// for range time.Tick(1 * time.Second) {
	// 	fmt.Println("for-----i", i)
	// 	i++
	// }

	fmt.Printf("%v", p)  //{maoyou3 24}
	fmt.Printf("%+v", p) //{Name:maoyou3 Age:24}

	t := T{
		Name: "t1",
	}

	t.M1("tm1")

	fmt.Println("\nt---m1", t.Name)

	t.M2("tm2")

	fmt.Println("\nt---m2", t.Name)

}

// T ...
type T struct {
	Name string
}

// M1 ...
func (t T) M1(name string) {
	t.Name = name
}

// M2 ...
func (t *T) M2(name string) {
	t.Name = name
}

// S ...
type S interface {
	M1()
	M2()
}
