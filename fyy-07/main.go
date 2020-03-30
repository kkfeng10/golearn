package main

import (
	"fmt"
	"strings"
)

// 字符串

func main() {
	// 打印一个windows下的路径
	winPath := "D:\\Go\\src\\github.com\\learngo" // 需要将路径的反斜杠都进行转义
	fmt.Println(winPath)
	// 打印一个带""的路径
	winPath2 := "\"D:\\Go\\src\\github.com\\learngo\"" // 需要将("")加反斜杠进行转义
	fmt.Println(winPath2)
	// 打印一个带‘’的路径
	winPath3 := "'D:\\Go\\src\\github.com\\learngo'" // 不需要将(‘’)加反斜杠进行转义，直接在""中使用' 即可
	fmt.Println(winPath3)
	// 打印多行字符串
	s1 := `
人生就像一场戏 因为有缘才相聚
相扶到老不容易 是否更该去珍惜
	`
	fmt.Println(s1)
	// 原样输出，在里面打印的字符串都不需要进行转义,比如打印一个windows下的路径
	s2 := `
D:\Go\src\github.com\learngo
	`
	fmt.Println(s2)
	// 打印长度
	fmt.Println(len(s2))
	// 字符串拼接,+ 拼接
	hello := "hello"
	world := "world"
	fmt.Println(hello + " " + world)
	// 字符串占位符输出
	fmt.Printf("%s %s\n", hello, world)
	// fmt.Sprintf 拼接，返回变量，不打印
	helloWorld := fmt.Sprintf("%s %s", hello, world)
	fmt.Println(helloWorld)
	// strings.Split 分割
	res1 := strings.Split(s2, "\\") // 返回一个类似数组的结果
	fmt.Println(res1)
	// strings.Contains 是否包含,返回true / false
	fmt.Println(strings.Contains(s1, "因为"))
	// strings.HasPrefix 前缀判断,返回 true / false
	fmt.Println(strings.HasPrefix(s1, "\n人生")) // 注意，因为多行输出的时候，将首行进行了换行，所以判断时开头部分是\n
	// strings.HasSuffix 后缀判断,返回 true / false
	fmt.Println(strings.HasSuffix(helloWorld, "ld"))
	// strings.Index 子串出现的位置,从0开始
	s3 := "abcdefgbcd"
	fmt.Println(strings.Index(s3, "c"))
	// strings.LastIndex 子串最后一次出现的位置
	fmt.Println(strings.LastIndex(s3, "bc"))
	fmt.Println(strings.LastIndex(s3, "cb")) // 不存在的时候返回 -1
	// String.Join 以选定的字符将列表内容进行拼接
	fmt.Println(strings.Join(res1, "+"))
}
