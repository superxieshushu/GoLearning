
### 待解答的问题

- string 字符串 “+” 连接操作为什么比 strings.join() 方法性能差 ？

- map 的迭代顺序不确定(随机) 、为什么这么设计 ？

- 当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调
  用处，直到另一个goroutine往这个channel里写入、或者接收值，这样两个goroutine才会继续
  执行channel操作之后的逻辑 ？

- os.Stdout 、fmt.Fprintf 、lissajous程序里我们输出的是一个文件等都实现了 io.Writer 接口