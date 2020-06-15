package main

import "fmt"

// 空接口 interface{} , 不要为了定义接口而定义接口
// 类型断言，需要知道空接口中的值是什么

// func show(a interface{}) {
// 	fmt.Printf("type is %T, value is %v\n", a, a)
// }

// func main() {
// 	var m map[string]interface{}
// 	m = make(map[string]interface{}, 10)
// 	m["name"] = "string"
// 	m["age"] = 17
// 	m["married"] = false
// 	m["array"] = [...]string{"a", "b", "c"}

// 	fmt.Printf("%v\n", m)
// 	show(m)
// }

// 类型断言
func guess(x interface{}) {
	str, ok := x.(string)
	if !ok {
		fmt.Println("不是字符串")
	} else {
		fmt.Println(str)
	}
}

// 类型断言2,利用switch case 的方式进行类型判断，x.(type) 来获取类型
func guess2(x interface{}) {
	fmt.Printf("x type is %T\n", x)
	// 在这个switch 语句中，是对type 进行了case 选择操作，t 就是x.(type)的返回值，如 guess()当中的str
	switch t := x.(type) {
	case string:
		fmt.Println("type is string ,value is: ", t)
	case int:
		fmt.Println("type is int ,value is: ", t)
	case bool:
		fmt.Println("type is bool ,value is: ", t)
	}
}

func main() {
	//guess(123)
	guess2(false)
}
