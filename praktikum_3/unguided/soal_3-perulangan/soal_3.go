package main

import "fmt"

func main() {
    var beratKantong1, beratKantong2 float64

    for {
        fmt.Print("Masukan berat belanjaan di kedua kantong: ")
        fmt.Scan(&beratKantong1, &beratKantong2)

        if beratKantong1 < 0 || beratKantong2 < 0 {
            break
        }

        totalBerat := beratKantong1 + beratKantong2
        if totalBerat > 150 {
            break
        }

        selisih := beratKantong1 - beratKantong2
        if selisih < 0 {
            selisih = -selisih
        }

        fmt.Println("Sepeda motor pak Andi akan oleng:", selisih >= 9)
    }

    fmt.Println("Proses selesai.")
}