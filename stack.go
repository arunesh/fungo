package stack

import "fmt"

type Stack struct {
    top int
    data [10]int
}

func (p *Stack) Push(k int) {
    if p.top < 9 {
        p.data[p.top] = k
        p.top ++
    }
}

func (p *Stack) Pop() int {
    p.top--
    x := p.data[p.top]
    p.data[p.top] = 0
    return x
}

func ExampleStack() {
    var s Stack
    s.Push(10)
    x := s.Pop()
    if x == 10 {
        fmt.Printf("works")
    }

    // Output:
    // works
}

