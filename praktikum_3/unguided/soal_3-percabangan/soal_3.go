package main

import "fmt"

func main() {
    var bilangan int
    fmt.Print("Bilangan: ")
    fmt.Scanln(&bilangan)

    if bilangan <= 1 {
        fmt.Println("Bilangan harus lebih besar atau sama dengan 1")
        return
    }

	print("Faktor: ")
    hitungfaktor := 0

    for i := 1; i <= bilangan; i++ {
        if bilangan%i == 0 {
            print(i, " ")
            hitungfaktor++
        }
    }

    if hitungfaktor == 2 {
        println("\nPrima: true")
    } else {
        println("\nPrima: false")
    }
}