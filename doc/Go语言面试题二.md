##### 1. 写出下面代码的输出内容。

```go
package main

import "fmt"

func main() {
    deferCall()
}

func deferCall() {
    defer func() { fmt.Println("打印前") }()
    defer func() { fmt.Println("打印中") }()
    defer func() { fmt.Println("打印后") }()

    panic("触发异常")
}
```

##### 2. 下面代码有什么问题？说明原因

```go
package main

import "fmt"

type student struct {
    Name string
    Age  int
}

func parseStudent() map[string]*student {
    m := make(map[string]*student)
    students := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 28},
        {Name: "wu", Age: 32},
        {Name: "wen", Age: 33},
    }

    for _, stu := range students {
        m[stu.Name] = &stu
    }
    return m
}


```

##### 3. go 语言全局变量的定义方式是怎么样的？

```go
ar a string = "a"

var b = "b"

c := "c"

var d string
d = "d"

func main() {
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
}
```

##### 4. go 语言中的结构体是如何定义的？

```go
// 使用struct可以定义结构体，具体可以参考如下的方式：
type Person struct {
    name string
    age int
}

func main() {
    p1 := Person{"Alice", 20}

    var p2 Person
    p2.name  "Panda"
    p2.age = 24

    fmt.Println(p1)
    fmt.Println(p2)
}
```

##### 5. go 语言通过指针访问成员变量的方式有几种？

```go
// 1. 直接使用 ``:=`` 可以获取变量的地址
// 2. 用&xxxx来获取地址
func main() {
    person := Person{"Alice", 20}

    p1 := &person

    fmt.Println(person)
    fmt.Println(&person)
    fmt.Println(&p1)
    fmt.Println(*p1)

    fmt.Println(person.name)
    fmt.Println(&person.name)
    fmt.Println(p1.name)
    fmt.Println(&p1.name)
}

type Person struct {
    name string
    age  int
}
```

##### 6. go 语言格式化输出的方式有哪些？

```go
    - %T 类型
    - %t 布尔
    - %d 10进制整数
    - %x 16进制整数
    - %f     浮点数
    - %s     字符串
    person := Person{"Alice", 20}
    fmt.Printf("%T\n", person)

    flag := true
    fmt.Printf("%t\n", flag)

    number10 := 99
    fmt.Printf("%d\n", number10)

    //十进制100为16进制的64
    number16:= 0x64
    fmt.Printf("%X\n", number16)
    fmt.Printf("%d\n", number16)

    number0:= 0.123
    fmt.Printf("%f\n", number0)

    str:= "hello world"
    fmt.Printf("%s\n", str)
```

##### 7. go 语言中的接口作用是什么？一个接口如果实现了一个接口的所有函数，那么？

```go
    // go 语言接口的作用是，可以实现OO面向对象的特性，从语法上看，Interface定义了一个或一组method(s)，这些method(s)只有函数签名，没有具体的实现代码（有没有联想起C++中的虚函数？）。下面是一个运用的实例：
    type MyInterface interface{
       Print()
   }

   func TestFunc(x MyInterface) {}
   type MyStruct struct {}
   func (me MyStruct) Print() {}

   func main() {
       var me MyStruct
       TestFunc(me)
   }
```

##### 8. go 语言中 init 函数有什么特性？能够在一个包里面写多个init吗？

```

```

##### 9. go 语言如何定义多参数函数， 调用其的方式有哪些？

```

```

##### 10. go 语言中是如何进行类型转换的？

```

```

##### 11. go 语言中引用类型有哪些？

```go
1. slice
2. map
3. channel
Golang的引用类型包括 slice、map 和 channel。它们有复杂的内部结构，除了申请内存外，还需要初始化相关属性。
```

##### 12. go 语言中引用的作用是什么？

```
能够让外部变量直接操作某块内存地址。 内置函数 new 计算类型大小，为其分配零值内存，返回指针。而 make 会被编译器翻译成具体的创建函数，由其分配内存和初始化成员结构，返回对象而非指针。

引用类型：
变量存储的是一个地址，这个地址所存储的值，
在内存上通常是将其分配在堆上面，
在程序中通过GC来进行回收。
获取指针类型的所指向的值，
可以使用： “*”取值符号。
比如 var *p int, 使用*p获取p指向的值
指针、slice、map、chan都是引用类型。
```

##### 13. go 语言main函数的特点有什么？

```
main函数不能带参数
main函数不能定义返回值
main函数所在的包必须为main
main函数中可以使用flag包来获取和解析命令行参数
```

##### 14. slice切片是如何初始化的？

实现切片初始化的方式，大致我学习到的可以分为两种：

- 使用make进行切片的初始化
- 使用数组直接定义初始化

具体可以参考下面的代码：

```go
func main() {
    s1 := make([]int, 0)
    s2 := make([]int, 6,10)
    s3 := []int{1, 2, 3, 4, 5}

    fmt.Println(s1)
    fmt.Println(s2)
    fmt.Println(s3)

    fmt.Println(len(s2))
    fmt.Println(cap(s2))
}
```

##### 15. go 语言中函数的定义方式有哪些？请举例说明。

go语言当中函数的定义，大致我概括一下可以分为如下几类：

- 不带参数的函数定义
- 带参数的函数定义
- 返回值有/无标识的函数定义
- 多参数函数定义
- 类函数定义

具体可以参考如下的代码：

```

```

##### 16. go 两个接口之间可以存在什么关系？

- 如果两个接口有相同的方法列表，那么他们就是等价的，可以相互赋值。
- 如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A。
- 接口查询是否成功，要在运行期才能够确定。

##### 17. go 当中同步锁有什么特点？作用是什么

- 当一个goroutine(协程)获得了Mutex后，其他gorouline(协程)就只能乖乖的等待，除非该gorouline释放了该Mutex
- RWMutex在 读锁 占用的情况下，会阻止写，但不阻止读
- RWMutex在 写锁 占用情况下，会阻止任何其他goroutine（无论读和写）进来，整个锁相当于由该goroutine独占

同步锁的作用是保证资源在使用时的独有性，不会因为并发而导致数据错乱，保证系统的稳定性。

##### 18. go 语言当中channel(通道)有什么特点，需要注意什么？

- 如果给一个 nil 的 channel 发送数据，会造成永远阻塞
- 如果从一个 nil 的 channel 中接收数据，也会造成永久爱阻塞
- 给一个已经关闭的 channel 发送数据， 会引起 pannic
- 从一个已经关闭的 channel 接收数据， 如果缓冲区中为空，则返回一个零值

##### 19. go 语言当中channel缓冲有什么特点？

无缓冲的 channel是同步的，而有缓冲的channel是非同步的。

##### 20. go 语言中cap函数可以作用于那些内容？

cap函数在讲引用的问题中已经提到，可以作用于的类型有：

- array(数组)
- slice(切片)
- channel(通道)

##### 21.  go convey是什么？一般用来做什么？

- go convey是一个支持golang的单元测试框架
- go convey能够自动监控文件修改并启动测试，并可以将测试结果实时输出到Web界面
- go convey提供了丰富的断言简化测试用例的编写

##### 22. go 语言中类型断言是什么？其作用是什么？举例说明。

##### 23. go 语言当中，切片是如何删除元素的？

##### 24. go 语言当中，如果对json进行重命名？

##### 25. go 语言当中，是如何实现类似继承的操作的？

##### 26. go 语言当中， 使用 `for range`迭代`map`是每次顺序是一样的吗？为什么？举例说明

```go
// 使用 for range迭代map时，每次迭代的顺序可能不一样，因为map的得带是随机的
func main() {
    m := make(map[string]int)
    m["string"] = 1

    m["int"] = 1
    m["float"] = 2
    m["bool"] = 3
    m["byte"] = 4


    for k, v := range m {
        // 打印的顺序会出现不一样的情况
        fmt.Println(k, ",", v)
    }
}
```

##### 27. go 语言中基本的数据类型有哪些？

| 序号 | 类型和描述                                                   |
| ---- | ------------------------------------------------------------ |
| 1    | **布尔型**：布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true |
| 2    | **数字类型**：整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。 |
| 3    | **字符串类型**：字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本 |
| 4    | **派生类型**：1. 指针类型（Pointer） 2. 数组类型 3. 结构化类型(struct) 4.Channel 类型 5. 函数类型 6. 切片类型 7. 接口类型（interface） 8. Map 类型 |

##### 28. go 语言中switch是如何运用的？有什么特殊的地方？

```go
//go 当中switch语句和其他语言类似，只是有一个特殊的地方，switch后面可以不跟表达式
func main() {
    i := rand.Intn(2)

    switch i {
    case 0:
        fmt.Println("get 0")
    case 1:
        fmt.Println("get 1")
    }
}

// switch后面可以不跟表达式
func main() {
    i := rand.Intn(2)

    switch {
    case i == 0:
        fmt.Println("get 0")
    case i == 1:
        fmt.Println("get 1")
    }
}
```

##### 29. go 语言结构体在序列化时，非导出变量（以小写字母开头的变量名）在解码的时候会出现什么问题？为什么？

结构体在序列化的时候非导出变量（以小写字母开头的变量名）不会被encode，所以在decode时这些非导出变量的值为其类型的零值。

##### 30. go 语言当中 new 和 make 有什么区别吗？

- `new`的作用是初始化一个纸箱类型的指针
- `new`函数是内建函数，函数定义：

```go
func new(Type) *Type
```

- 使用`new`函数来分配空间
- 传递给`new`函数的是一个类型，而不是一个值
- 返回值是指向这个新非配的地址的指针