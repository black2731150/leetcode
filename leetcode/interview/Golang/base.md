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
