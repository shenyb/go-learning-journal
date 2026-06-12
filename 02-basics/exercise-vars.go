package main

import "fmt"

func main() {
	// === 练习1：变量声明 ===
	// Java: String name = "申";
	// Go:   name := "申"
	//
	// 下面声明三个变量：你的名字(name)、年龄(age)、是否已婚(married)
	// 用 := 短声明
	// TODO: 在这里写代码
	name := "申"
	age := 35
	married := true


	// === 练习2：打印 ===
	// 用 fmt.Printf 打印出 "我叫XX，今年XX岁，已婚：true/false"
	fmt.Printf("我叫%s，今年%d岁，已婚：%t\n", name, age, married)

	// Go 的格式化占位符：%s=字符串 %d=整数 %t=布尔
	// TODO: 在这里写代码
	// === 练习3：if 判断 ===
	// 如果年龄大于30，打印"中年危机警告 🚨"
	// 否则打印"还年轻 💪"
	// TODO: 在这里写代码
	if age > 30 {
		fmt.Println("中年危机警告 🚨")
	} else {
		fmt.Println("还年轻 💪")
	}
}
