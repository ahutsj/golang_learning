package main

import (
    "fmt"
)

func main () {
    type stu struct {
        name string
        age  int
    }
    xiaoming := stu{
        name: "AAA",
        age:  10,
    }
    p := &xiaoming
    //p := xiaoming
    p.name = "BBB"
    fmt.Println(xiaoming.name)
    //结构体在内存中成员的分布是连续的，就像数组一样获取不到它的地址，但是可以获取结构体的成员地址
    fmt.Println(&xiaoming.name)
    fmt.Println(&p.name)
}
