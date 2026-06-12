# 01 — 环境搭建 + 第一印象

## 安装

```bash
brew install go
```

确认版本：

```bash
go version
# go1.24.x darwin/arm64
```

## Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 邯郸! 等我回来.")
}
```

Java 对比：

| 维度 | Java | Go |
|------|------|----|
| 入口 | public static void main(String[] args) | func main() |
| 导入 | import java.util.* | import "fmt" |
| 分号 | 必须 | 基本不用（编译器自动加） |
| 编译 | javac → java | go build → ./binary |

最直观的感受：**Go 不需要类**。没有 public class HelloWorld 包一层，直接写 func main() 就行。

## 第一印象：像 Java 和 C 的结合体

- 大括号风格跟 Java 一样（但 if/for 的条件没有括号）
- 编译速度和 go build 的感觉像 C（但依赖管理比 C 友好）
- 没有构造方法、没有继承、没有泛型（1.18 之后有了）——少了很多 Java 的重量感

## 值得注意的点

1. **`package main` 是可执行程序的约定** —— 不像 Java 里面包路径反着域名写
2. **未使用的变量/包直接编译报错** —— Java 只是 warning，Go 直接不让过。一开始略烦，但确实逼你写干净代码
3. **`:=` 短变量声明** —— 不用写类型，编译器推断。这个比 Java 的 `var` 还简洁

```
x := 42          // 不用写 int x = 42
var y int = 42   // 完整写法（很少用）
```

## TODO

- [x] 安装 Go
- [x] 跑通 Hello World
- [x] 看完官方 Tour of Go 基础部分
