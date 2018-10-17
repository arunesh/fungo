package main

import "fmt"
import "strconv"

/* print something in the main function */

func main() {
//    var b string

      var myarray [10]int
    myarray[0] = 5
    fmt.Println("myarray[0] =", myarray[0])
    slice := myarray[:9]
    slice2 := myarray[:5]
    fmt.Println("slice[0] =", slice[0])
    fmt.Println("slice2[0] =", slice2[0])

    slice2 = append(slice2, 3, 5, 7)
    myarray[0] = 1
    fmt.Println("slice[0] =", slice[0])
    fmt.Println("slice2[0] =", slice2[0])
    fmt.Println("slice2 =", slice2)
    fmt.Println("slice1 =", slice)
    fmt.Println("myarray =", myarray)

    fmt.Println(" compareWith(5, 10) = ", compareWith(5, 10, returnValue))
    variadic(1, 5, 7, 9, 30, 4, 634)

    cd, _ := doesPanic(makePanic)
    fmt.Println("Did panic?", cd)
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
