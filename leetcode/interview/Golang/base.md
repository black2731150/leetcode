# Golang 基础知识

## make 和 new 的区别
- new(T)用于给给定类型T分配一个内存空间，并返回指向T的类型的零值的指针
- make(T,len,cep)用户给切片，映射和通道分配内存，并返回其初始化的值，而非指针。其中T是数据类型，len是类型的长度，cep是类型的容量

## Golang中切片的扩容规则是怎么样的
    1. 倍增规则：对于长度小于1024的切片，当添加元素超出切片当前的内容时，Go会创建一个新切片，其容量通常是原来容量的2倍。
    2. 1.25倍增长：对于长度大于1024的切片，Go可能会选择每一次增加原容量的25%。
    3. 精确扩容：如果是通过append函数扩容请求明确了需要的容量，go将分配足够容纳所有新元素的空间。

## 单个package中，init函数与常量、全局变量的执行顺序
- 会先执行导包，然后是变量的初始化，然后是常量的初始化，最后是init函数的执行

## slice作为函数传值问题
- A函数调用B函数，B函数入参为slice，slice元素类型为string，如果在B函数中对传入的slice做了修改，那么调用完成后，在A函数中再次读取slice的值是否会有变化，为什么?
```golang
func testA() {
        s := make([]string, 0)
        for i := 0; i < 10; i++ {
                s = append(s, strconv.Itoa(i))
        }
        fmt.Println(s)
        testB(s)
        fmt.Println(s)
}
func testB(s []string) {
        for i := range s {
                s[i] = "new-" + s[i]
        }
}
```
- golang中函数的参数传递都是值传递，Slice是一个结构体，由指向数组的指针*ptr，切片的长度len，和切片的容量cap组成。由于是值的传递，所以会传递一样的数组指针，所以底层指向的是同一块内存，所以B函数中对切片的更改会影响A函数中的切片。

## Golang如何进行异常处理，优缺点是什么，如果发生异常了如何处理
- Golang中由error标准库，可以用这个库来处理异常
- 优点是很完善，几乎包含了所有可能发生的错误和应对方法，缺点是比较麻烦
- 一般用defer语句去处理错误，比如关闭打开的文件，指定出现错误时候的应对逻辑

## Go GC机制

1. 并发标记-清除（Concurrent Mark-Sweep）
Go的GC是一个并发的标记-清除垃圾回收器。这意味着它在执行垃圾回收时，应用程序的goroutines仍然可以运行，从而减少了停顿时间。
标记阶段（Marking Phase）：GC遍历所有的根对象（如全局变量、活跃的goroutine栈上的变量），然后递归地遍历它们引用的对象。在这个过程中，访问到的每个对象都会被标记为“可达”的，即这些对象正在被使用，不应该被回收。
清除阶段（Sweeping Phase）：GC遍历堆中的所有对象，回收那些在标记阶段未被标记为“可达”的对象所占用的内存。

2. 写屏障（Write Barrier）
为了让标记阶段和程序的正常执行并发进行，Go使用了写屏障技术。当程序在标记阶段修改对象引用时，写屏障确保这些修改被正确处理，以保持GC的正确性。这意味着即使在对象图发生变化时，GC也能正确地标记活跃的对象。

3. STW（Stop-The-World）
尽管Go的GC是并发的，但在标记的开始阶段和最终阶段，仍然需要短暂的STW（停止所有程序执行），以确保GC的安全性和一致性。这些STW的停顿通常很短。

4. GC触发
Go的GC可以通过几种方式触发：
定期触发：基于分配的内存量。当从上一次GC以来分配的内存达到一定阈值时，会自动触发GC。
手动触发：通过调用runtime.GC()函数手动触发GC。这会引起一个STW，然后执行完整的GC周期。