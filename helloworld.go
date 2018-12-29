package main

import "fmt"
import "strconv"
import "world"

/* print something in the main function */

type stack struct {
    top int
    data [10]int
}

func (p *stack) push(k int) {
    if p.top < 9 {
        p.data[p.top] = k
        p.top ++
    }
}

func (p *stack) pop() int {
    p.top--
    x := p.data[p.top]
    p.data[p.top] = 0
    return x
}

func main() {
//    var b string

    var x []int

    x = make([]int, 10)

    x[1]= 109
    fmt.Printf("X = %v\n", x)

    fmt.Println(" compareWith(5, 10) = ", compareWith(5, 10, returnValue))
    variadic(1, 5, 7, 9, 30, 4, 634)

    cd, _ := doesPanic(makePanic)
    fmt.Println("Did panic?", cd)

    fmt.Println("Average:", avg([]float64{0.6, 0.7, -0.9, 2.5}))
    fmt.Println("Average:", avg([]float64{}))

    newline(1, 4,6, 78, 2)
    p := plusX(5)
    fmt.Println("plus 5 = ", p(2))
    fmt.Println("array = ", mapf(p, 4, 5, 6, 8, 2))

    var s stack
    s.push(10)
    s.push(1000)
    fmt.Println("S = ", s)
    fmt.Println("Pop =", s.pop())
    fmt.Println("new S =", s)
    fmt.Println("Is even =", world.IsEven(4))

    c := make(chan int)
    go fib(10, c)
    fmt.Println("Fib:")
    for i := range c {
        fmt.Println(i)
    }
}

func mapf(x func (int) int, y...int) (r []int) {
    r = make([]int, len(y))
    for i, v := range y {
        r[i] = x(v)
    }
    return
}

func plusX(x int) (r func(y int) int) {
    r = func(y int) int {
        return y + x
    }
    return
}

func plusTwo() (x func(y int) int) {
    x = func(y int) int {
        return y + 2
    }
    return
}

func newline(x...int) {
    for i, y := range x {
        fmt.Printf("array[%d] = %d \n", i, y)
    }
}

func avg(x []float64) (avg float64) {
    avg = 0.0
    switch len(x) {
    case 0:
        avg = 0.0
    default:
    for i:= 0; i < len(x); i++ {
        avg += x[i]
    }
    avg = avg / float64(len(x))
}
return
}

func fib(n int, c chan int) {

    f := make([]int, n)
    f[0], f[1] = 1, 1
    for i := 2; i < n; i ++ {
        f[i] = f[i -1] + f[i - 2]
        c <- f[i]
    }
    close(c)
}

func returnValue(x int, y int) int {
    if  x > y  {
        return x
    } else {
        return y
    }
}

func compareWith(x int, y int, c func (int, int)int) int {
    return c(x, y)
}

func makePanic() {
    var x []int
    x[2] = 5
}

func doesPanic(f func()) (b bool, x int) {
    b = false
    defer func() {
        if x:= recover(); x != nil {
            b = true
            fmt.Println("PaniC!!")
        }
    }()
    f()
    return b, x
}

func replace(s string){
    c :=  []rune(s)
    c[0] = '1'
    fmt.Println(string(c))
}

func variadic(arg ... int) {
    for i, j := range arg {
        fmt.Printf("arg[%d] = %d \n", i, j)
    }
}

func stringArrayDemo() {
 const (
       mymsg = "Value of a ="
       tempInt = 22345
   )

    _, a := 15, 100
    myList := []string{"a1", "b2", "c3", "d4", "e5"}

    for k, val := range(myList) {
        fmt.Println("key ", k, " value", val)
        for k1, v1 := range(val) {
            fmt.Printf("key %d, val %c \n", k1, v1)
        }
    }
    replace(strconv.Itoa(tempInt))
    fmt.Println(mymsg, a)


}
