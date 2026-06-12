# 02 — 变量、打印、条件判断、基本类型、常量

## 练习文件

| 文件 | 内容 |
|------|------|
| [exercise-vars.go](./exercise-vars.go) | 变量声明、格式化打印、if 判断 |
| [exercise-types.go](./exercise-types.go) | 基本类型细分、零值、常量 & iota |

> 两个文件分别用 `go run exercise-vars.go` 和 `go run exercise-types.go` 执行。不要对目录整体 `go run .`（因为都有 `func main`）。

---

## 变量声明

```go
name := "申"          // 短声明（最常用）—— 类型由编译器推断
var age int = 35      // 完整声明（函数内少用）
var married bool      // 零值声明（没赋值就是 false）
```

**关键区别**：`:=` 只在函数内可用。包级别变量必须用 `var`。

---

## 基本数据类型（vs Java）

### 整数

| Go 类型 | 位数 | 范围 | 对比 Java |
|---------|------|------|-----------|
| `int`   | 平台相关（arm64=64） | 同 int64 | Java 固定 32 位 |
| `int8`  | 8   | -128~127 | `byte` |
| `int16` | 16  | -32768~32767 | `short` |
| `int32` | 32  | -2^31~2^31-1 | `int` |
| `int64` | 64  | -2^63~2^63-1 | `long` |
| `uint`  | 平台相关 | 无符号 | Java 无（除 char） |
| `byte`  | 8   | 0~255 | 无（`byte` 是有符号） |
| `rune`  | 32  | Unicode 码点 | `int` / `char` |

**关键差异**：Go 没有 `long` / `short` 这种别名，直接写 `int64` / `int16`。Java 的 `byte` 是有符号的，Go 的 `byte` 是无符号的（=uint8）。

### 浮点数

| Go 类型 | 精度 | 对应 Java |
|---------|------|-----------|
| `float32` | ~7 位小数 | `float` |
| `float64` | ~15 位小数 | `double` |

注意：`float32(math.Pi)` 打印到第 7 位就开始漂了（`3.1415927410`），`float64` 才准。

### 复数（Java 没有）

```go
var c complex64 = 1 + 2i
real(c)  // 实部 1.0
imag(c)  // 虚部 2.0
```

### 字符串

- Go 字符串是 **UTF-8 字节序列**，不是 UTF-16
- `len("邯郸")` = 6（一个汉字 3 字节），不是 2
- 遍历字符串要用 `[]rune(s)` 才能拿到字符

---

## 零值（Zero Values）

Go 最舒服的特性之一：声明不赋值也有默认值，**没有 null**。

```go
var i int          // 0
var f float64      // 0.0
var b bool         // false
var s string       // ""
var slice []int    // nil（但不像 Java 的 null 会 NPE——append 可以直接用）
var m map[int]bool // nil
var p *int         // nil
```

> Java 里这些全部是 null/NPE 风险点。Go 的做法是：**你声明了就能用**，不用为初始化操心。

---

## 常量 & iota

### 普通常量

```go
const pi = 3.14159         // 未指定类型（灵活）
const typed int = 100      // 指定类型（严格）
```

**类型化 vs 非类型化**：`const untyped = 100` 可以赋给 `int32`、`int64`、`float64`。类型化常量不行——`const typed int = 100` 不能隐式转 `int32`。这在 Java 里是做不到的（int 就是 int）。

### iota（自增枚举器）

基本用法：

```go
const (
    Spring = iota  // 0
    Summer         // 1
    Autumn         // 2
    Winter         // 3
)
```

跳过值：

```go
const (
    _ = iota       // 0 被丢弃
    Monday         // 1
    ...
    Sunday         // 7
)
```

### 位运算枚举（经典模式）

```go
const (
    Read    = 1 << iota  // 1 (0001)
    Write                // 2 (0010)
    Execute              // 4 (0100)
)

perm := Read | Execute   // 5 (0101)
perm & Read != 0         // true——有读取权限
perm & Write != 0        // false——无写入权限
```

> 这个比 Java 的 `EnumSet` 轻量多了。Go 没有 Java 的 enum 类型，iota 就是做枚举的主流方式。

---

## 格式化打印

| 占位符 | 含义 | 对比 Java |
|--------|------|-----------|
| `%s` | 字符串 | `%s` |
| `%d` | 十进制整数 | `%d` |
| `%t` | 布尔值 true/false | 无直接对应 |
| `%v` | 任意类型的默认格式 | 类似 `String.valueOf()` |
| `%+v` | 结构体带字段名 | 无 |
| `%T` | 类型名 | 无 |
| `%#v` | Go 语法表示 | 无 |

```go
fmt.Printf("我叫%s，今年%d岁，已婚：%t\n", name, age, married)
```

---

## if 语句

```go
if age > 30 {
    fmt.Println("中年危机警告 🚨")
} else {
    fmt.Println("还年轻 💪")
}
```

**和 Java 的差别**：

1. **条件没括号** —— `if age > 30` 不是 `if (age > 30)`
2. **大括号必须换行** —— 不能写成 `if age > 30\n{`
3. **if 前可以写短语句** —— `if score := 85; score >= 60 { ... }`

---

## Java 对比总表

| 维度 | Java | Go |
|------|------|----|
| 变量声明 | `String name = "申";` | `name := "申"` |
| 未使用变量 | Warning | **编译错误** |
| 默认值 | null（对象）/ 0（数字） | 零值（永远可用） |
| if 条件 | `if (x > 0) ` | `if x > 0` |
| 条件前声明 | 不支持 | `if x := get(); x > 0` |
| 整数类型 | int(32), long(64), short(16), byte(8有符号) | int(平台), int8/16/32/64, byte(8无符号) |
| 复数 | 无 | `complex64` / `complex128` |
| 字符串编码 | UTF-16 | UTF-8 |
| 枚举 | `enum` 关键字 | `iota`（更灵活） |
| 布尔占位符 | `%b`（是位运算） | `%t` |

## 值得注意的点

1. **未使用变量 = 编译失败** —— 一开始有点烦，但逼你写干净代码
2. **`:=` 是声明+赋值** —— 对同一个变量重复 `:=` 会报错（除非至少有一个新变量）
3. **零值机制** —— 声明即安全，没有 NPE 焦虑。这是 Java 转 Go 最舒爽的体验之一
4. **int 是平台相关** —— arm64（M 芯片 Mac）上是 64 位，不像 Java 的 int 固定 32 位
5. **rune ≠ char** —— Go 没有 `char` 类型。`rune` 是 Unicode 码点（int32），`byte` 是 ASCII（uint8）
6. **iota 从 0 开始** —— 如果要跳过 0，用 `_ = iota`

## TODO

- [x] 短声明 `:=`
- [x] 变量类型推断
- [x] `fmt.Printf` 占位符
- [x] if 条件判断
- [x] 基本数据类型细分（整数、浮点、布尔、字符串、复数）
- [x] 零值机制
- [x] 常量 & iota
- [ ] 下一步：数组、切片、map
