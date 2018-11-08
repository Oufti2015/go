package main

import "fmt"
import "os"

func vals() (int, int) {
    return 3, 7
}

func splitPrice(price float32) (dollars, cents int) {
    dollars = int(price)
    cents = int((price - float32(dollars)) * 100.0)
    return
}

func getFileSize() (file_size int64, had_error bool) {
    f, err := os.Open("/tmp/dat")
    if err != nil {
        return 0, true
    }
    defer func() {
        if err := f.Close(); err != nil {
            had_error = true
        }
    }()

    fi, err := f.Stat()
    if err != nil {
        return 0, true
    }
    return fi.Size(), false
}

func main() {

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    fmt.Println(c)

    dollars, cents := splitPrice(12.42)
    fmt.Println(dollars, cents)

    file_size, had_error := getFileSize()
    fmt.Println(file_size, had_error)
}