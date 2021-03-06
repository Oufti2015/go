package main

import "fmt"

func zeroval(ival int) {
	fmt.Println("zeroval", " pointers:", &ival)
    ival = 0
}

func zeroptr(iptr *int) {
	fmt.Println("zeroptr", " pointers:", iptr)
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}