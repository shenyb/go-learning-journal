package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	// === 练习1：基本数据类型细分类 ===

	// 整数：有符号 vs 无符号
	// Java 的 int 永远是 32 位，Go 的 int 看平台
	var a int = 42
	var b int8 = 127          // -128 ~ 127
	var c int16 = 32767       // -32768 ~ 32767
	var d int32 = 2147483647  // -2^31 ~ 2^31-1
	var e int64 = 9223372036854775807

	fmt.Println("=== 整数类型 ===")
	fmt.Printf("int   大小=%d字节  值=%d\n", unsafe.Sizeof(a), a)
	fmt.Printf("int8  大小=%d字节  值=%d\n", unsafe.Sizeof(b), b)
	fmt.Printf("int16 大小=%d字节  值=%d\n", unsafe.Sizeof(c), c)
	fmt.Printf("int32 大小=%d字节  值=%d\n", unsafe.Sizeof(d), d)
	fmt.Printf("int64 大小=%d字节  值=%d\n", unsafe.Sizeof(e), e)

	// 无符号整数
	var u uint = 99
	fmt.Printf("uint  大小=%d字节  值=%d\n", unsafe.Sizeof(u), u)

	// byte = uint8, rune = int32（Unicode 码点）
	var ch byte = 'A'          // ASCII 字符
	var rn rune = '爪'         // Unicode 字符
	fmt.Printf("\n=== byte & rune ===\n")
	fmt.Printf("byte='%c' (%d), rune='%c' (U+%04X)\n", ch, ch, rn, rn)

	// 浮点数
	var f32 float32 = math.Pi
	var f64 float64 = math.Pi
	fmt.Printf("\n=== 浮点数 ===\n")
	fmt.Printf("float32: %.10f\n", f32)
	fmt.Printf("float64: %.10f\n", f64)

	// 复数（Java 没有）
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 3 + 4i
	fmt.Printf("\n=== 复数 ===\n")
	fmt.Printf("complex64:  %v (实部=%.1f, 虚部=%.1f)\n", c64, real(c64), imag(c64))
	fmt.Printf("complex128: %v\n", c128)

	// 布尔
	var flag bool = true
	fmt.Printf("\n=== 布尔 ===\n")
	fmt.Printf("bool: %t (大小=%d字节)\n", flag, unsafe.Sizeof(flag))

	// 字符串
	var s string = "邯郸" // UTF-8，不是 UTF-16
	// 长度是字节数不是字符数
	fmt.Printf("\n=== 字符串 ===\n")
	fmt.Printf("string: %q\n", s)
	fmt.Printf("len=%d 字节", len(s))
	fmt.Printf(" 字符数=%d (用 []rune 遍历才是字符)\n", len([]rune(s)))

	// === 零值演示 ===
	// Go 声明不赋值也有默认值，没有 null
	fmt.Printf("\n=== 零值（Zero Values）===\n")
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	var zeroSlice []int
	var zeroMap map[string]int
	var zeroPtr *int
	fmt.Printf("int:     %d\n", zeroInt)
	fmt.Printf("float64: %f\n", zeroFloat)
	fmt.Printf("bool:    %t\n", zeroBool)
	fmt.Printf("string:  %q\n", zeroString)
	fmt.Printf("slice:   %v (is nil? %t)\n", zeroSlice, zeroSlice == nil)
	fmt.Printf("map:     %v (is nil? %t)\n", zeroMap, zeroMap == nil)
	fmt.Printf("pointer: %v (is nil? %t)\n", zeroPtr, zeroPtr == nil)

	// Java 里这些全部是 null/NPE 风险点

	// === 常量 & iota ===
	fmt.Printf("\n=== 常量与 iota ===\n")

	// 普通常量
	const pi = 3.1415926
	const greeting = "你好，Go"
	fmt.Printf("pi = %v, 类型为 %T\n", pi, pi)

	// iota：自增枚举器（类似 Java enum 的 ordinal，但更灵活）
	const (
		Spring = iota // 0
		Summer        // 1
		Autumn        // 2
		Winter        // 3
	)
	fmt.Printf("Spring=%d, Summer=%d, Autumn=%d, Winter=%d\n", Spring, Summer, Autumn, Winter)

	// iota 跳过值
	const (
		_     = iota      // 0 被丢弃
		Monday            // 1
		Tuesday           // 2
		Wednesday         // 3
		Thursday          // 4
		Friday            // 5
		Saturday          // 6
		Sunday            // 7
	)
	fmt.Printf("Monday=%d, Friday=%d, Sunday=%d\n", Monday, Friday, Sunday)

	// iota 与位运算（经典用法：权限标记）
	const (
		Read   = 1 << iota // 1 (0001)
		Write              // 2 (0010)
		Execute            // 4 (0100)
	)
	perm := Read | Execute // 权限组合
	fmt.Printf("Read=%d, Write=%d, Execute=%d\n", Read, Write, Execute)
	fmt.Printf("perm=Read|Execute = %d (二进制 %04b)\n", perm, perm)
	fmt.Printf("有 Read 权限?  %t\n", perm&Read != 0)
	fmt.Printf("有 Write 权限? %t\n", perm&Write != 0)

	// 类型化常量 vs 非类型化常量
	const typedInt int = 100          // 类型固定
	const untypedInt = 100           // 灵活类型，可以赋给 int/int64/float64
	var x int32 = untypedInt         // 编译通过（untyped 可以隐式转换）
	// var y int32 = typedInt        // 编译错误：不能将 int 隐式转为 int32
	fmt.Printf("\n类型化 vs 非类型化：untyped 可以赋给 int32: %d\n", x)

	// 练习：定义星期的常量（用 iota）
	// TODO: 用 iota 定义 Weekday 常量，从 Sunday=0 到 Saturday=6
	// fmt.Printf("Sunday=%d, Monday=%d, Tuesday=%d\n", ...)
}
