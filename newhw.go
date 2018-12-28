package main

import "fmt"
import "math/rand"
import "time"

const s = "hello jupiter"


func initSeq() func () int {
    i := 0
    return func () int {
        i++
        return i
    }
}

func main() {
    nextSeq := initSeq()
    fmt.Println("next = ", nextSeq())
    fmt.Println("next = ", nextSeq())
    fmt.Println("next = ", nextSeq())

    nextSeq2 := initSeq()
    fmt.Println("next2 = ", nextSeq2())
    fmt.Println("next2 = ", nextSeq2())
    fmt.Println("next = ", nextSeq())

    fmt.Println("A demo of slicing.")

    var p = []int{2, 4, 7, 13, 16, 24}

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(p)
    fmt.Println(s)

    sum := 0
    for _, v := range p {
        sum = sum + v
    }
    fmt.Printf("sum = %d\n", sum)

    printRand()
    fmt.Println(time.Now())

    newSlice := make([]string, 1, 100)
    fmt.Printf("Size of string: %d, capacity: %d\n", len(newSlice), cap(newSlice))

    someFunc := printRand

    someFunc()

    fmt.Println("A ")
    fmt.Println(p)

    m := make(map[string]int)
    m["a"] = 1
    m["b"] = 2

    for k, v := range m {
        fmt.Println("key :", k, " value:", string(v))
        fmt.Printf("value: %d\n:", v)
    }

    for i, c := range "Arunesh Mishra" {
        fmt.Println("i = ", i, " c = ", c)
    }
    fmt.Println("Function Sum =", doSum(2, 5, 6))
}

func printRand() {

    i := rand.Intn(100)

    fmt.Printf("i = %d\n", i)

    switch i%3 {
    case 0:
        fmt.Println("Divisible by 3")
    case 1:
        fmt.Println("Remainder 1, so even.")
    case 2:
        fmt.Println("Remainder 2, so odd.")

    default:
        fmt.Println("This cannot happen")

    }

    return
}

func doSum(vals ...int) int {
    fmt.Println("Number of values:", len(vals))
    sum := 0
    for _, v := range vals {
        sum += v
    }
    return sum
}
