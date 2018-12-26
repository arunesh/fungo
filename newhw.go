package main

import "fmt"
import "math/rand"
import "time"

const s = "hello jupiter"

func main() {
    fmt.Println("A demo of slicing.")

    var p = []int{2, 4, 7, 13, 16, 24}

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(p)
    fmt.Println(s)

    printRand()
    fmt.Println(time.Now())
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
