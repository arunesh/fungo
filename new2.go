package main

import "fmt"

type person struct {
    name string
    age int
}

func (p person) prettyPrint() {
    fmt.Println("Name =", p.name)
    fmt.Println("Age =", p.age)
}

func (p person) incAge(x int) person {
    p.age += x
    return p
}

func (s person) String() string {

    return fmt.Sprintf("S Person Name: %s, Age: %d\n", s.name, s.age)
}

func sums(val ...int) int {

    fmt.Println("number of values:", len(val))

    sum := 0
    for _, v := range val {
        fmt.Println("val = ", v)
        sum += v
    }
    fmt.Println("Sum = ", sum)

    var personA = person{"Jack Hill", 39}
    var ptr *person

    ptr = &personA

    fmt.Println("Person A ", personA)
    fmt.Println("Person ptr ", ptr)
    fmt.Println("after inc", personA.incAge(10))
    fmt.Println("Person A ", personA)
    return sum
}


func describe(i interface{} ) {

    fmt.Printf("( %v, %T)\n", i, i)
}

func main() {

    fmt.Println("Still learning GO")

    sums(1, 4, 5, 9 , 19, -25)
}
