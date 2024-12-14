package main

import "fmt"

func perkalian(n, m int) int {
    if n == 0 {
        return 0
    }
    return m + perkalian(n-1, m)
}

func main() {
    var n, m int
    
    fmt.Scan(&n, &m)
    
    hasil := perkalian(n, m)
    fmt.Println(hasil)
}