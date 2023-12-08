package main

import "fmt"

// Parent 表示父类型
type Parent struct {
	ParentField int
}

// Child 表示子类型，嵌套了父类型
type Child struct {
	Parent     // 嵌套的父类型
	ChildField int
}

func main() {
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	a, ok0 := m["3"], false //false
	_, ok1 := m["4"]        //false
	_, ok2 := m["4"], false //false
	_, ok3 := m["4"], true  //true
	_, ok4 := m["3"]        //true

	fmt.Println(ok0)
	fmt.Println(ok1)
	fmt.Println(ok2)
	fmt.Println(ok3)
	fmt.Println(ok4)

	childObj := Child{
		Parent:     Parent{ParentField: 42}, // 初始化父类型字段
		ChildField: 10,                      // 初始化子类型字段
	}

	//两种方式都不成立
	var parentobj Parent

	//不支持隐式转换
	parentobj = childObj //x
	childObj = parentobj //x

	//这种方式的显示转换也不支持
	parentobj = Parent(childObj) //x
	childObj = Child(parentobj)  //x

	parentobj = childObj.Parent //√
	parentobj = childObj.Parent //√

}
