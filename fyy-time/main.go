package main

// time 时间函数

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.UTC())
	fmt.Println(now.Year())                                              // 年
	fmt.Println(now.YearDay())                                           // 一年当中的第多少天[1-365] 或者 [1-366]
	specT, _ := time.Parse("2006-01-02 15:04:05", "2019-12-10 11:13:33") // 自定义时间，返回time.Time 结构体
	fmt.Println(now.Year())
	fmt.Println(specT.YearDay())
	fmt.Println(specT.Month())                           // 月
	fmt.Println(specT.Day())                             // 日
	fmt.Println(specT.Hour())                            // 时
	fmt.Println(specT.Minute())                          // 分
	fmt.Println(specT.Second())                          // 秒
	fmt.Println(specT.Add(time.Duration(1 * time.Hour))) // 在自定义时间上 增加一个小时的时间
	fmt.Println(specT.Unix())                            // unix 时间戳
	fmt.Println(specT.UTC())
	// st := time.Tick(time.Duration(1 * time.Second)) // 定时器，1秒触发一次动作
	// for k := range st {
	// 	fmt.Println(k)
	// }
	// time.Sleep(time.Duration(3 * time.Second))  // sleep 挂起
	// fmt.Println("3秒后执行的我")
	bf := specT.Before(now) // 判断时间先后顺序
	fmt.Println(bf)
	subRes := specT.Sub(now)
	fmt.Println(subRes.Minutes())
}
