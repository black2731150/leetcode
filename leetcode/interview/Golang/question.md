# 代码实现一些问题

## golang中如何实现producer-consumer模式

```golang
package main

import (
    "fmt"
    "sync"
    "time"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        ch <- i  // 生产数据并发送到通道
        time.Sleep(time.Millisecond * 100) // 模拟耗时的生产过程
    }
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := range ch {
        fmt.Printf("消费者收到数据: %d\n", i) // 从通道接收数据并消费
        time.Sleep(time.Millisecond * 150) // 模拟耗时的消费过程
    }
}

func main() {
    var wg sync.WaitGroup
    ch := make(chan int, 5) // 创建一个带缓冲的通道

    // 启动生产者协程
    wg.Add(1)
    go producer(ch, &wg)

    // 启动消费者协程
    wg.Add(1)
    go consumer(ch, &wg)

    // 等待生产者和消费者完成
    wg.Wait()
    close(ch) // 关闭通道
}
```

## 写两个goroutine。交替打印1和2。要求：不能sleep。不能用有长度的chan，chan最多一个。
- 1和2顺序没有要求的情况
```golang
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    ch := make(chan struct{})

    wg.Add(2)

    // 第一个 goroutine 打印 "1"
    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ { // 为了演示，仅重复5次
            <-ch // 等待信号
            fmt.Println("1")
            ch <- struct{}{} // 发送信号给第二个 goroutine
        }
    }()

    // 第二个 goroutine 打印 "2"
    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ { // 为了演示，仅重复5次
            fmt.Println("2")
            ch <- struct{}{} // 发送信号给第一个 goroutine
        }
    }()

    // 开始交替打印
    ch <- struct{}{}

    wg.Wait()
    close(ch)
}
```

- 1和2顺序有要求，必须先打印1
```golang
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    ch := make(chan struct{})

    wg.Add(2)

    // 第一个 goroutine 打印 "1"
    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ { // 为了演示，仅重复5次
            fmt.Println("1")
            ch <- struct{}{} // 先打印 "1"，然后发送信号给第二个 goroutine
            <-ch // 等待第二个 goroutine 打印 "2"
        }
    }()

    // 第二个 goroutine 打印 "2"
    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ { // 为了演示，仅重复5次
            <-ch // 等待第一个 goroutine 打印 "1"
            fmt.Println("2")
            ch <- struct{}{} // 发送信号给第一个 goroutine
        }
    }()

    wg.Wait()
    close(ch)
}
```
