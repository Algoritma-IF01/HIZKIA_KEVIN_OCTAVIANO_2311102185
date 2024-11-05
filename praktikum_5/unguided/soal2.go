package main

import (
    "fmt"
    "math"
)

func soal_2() {
    const nMax = 100
    var n, x int

    for {
        fmt.Print("Masukkan jumlah elemen array (maksimum 100): ")
        fmt.Scan(&n)

        if n > nMax {
            fmt.Printf("Jumlah elemen tidak boleh lebih dari %d. Silakan coba lagi.\n", nMax)
        } else {
            break
        }
    }

    arr := make([]int, n)

    fmt.Println("Masukkan elemen array:")
    for i := 0; i < n; i++ {
        fmt.Printf("Elemen ke-%d: ", i)
        fmt.Scan(&arr[i])
    }

    fmt.Print("Isi array: ")
    for i := 0; i < n; i++ {
        fmt.Printf("%d ", arr[i])
    }
    fmt.Println()

    fmt.Print("Elemen bernilai ganjil: ")
    for i := 0; i < n; i++ {
        if arr[i]%2 != 0 {
            fmt.Printf("%d ", arr[i])
        }
    }
    fmt.Println()

    fmt.Print("Elemen bernilai genap: ")
    for i := 0; i < n; i++ {
        if arr[i]%2 == 0 {
            fmt.Printf("%d ", arr[i])
        }
    }
    fmt.Println()

    fmt.Print("Masukkan nilai x untuk indeks kelipatan: ")
    fmt.Scan(&x)
    fmt.Printf("Elemen pada indeks kelipatan %d: ", x)
    for i := 0; i < n; i++ {
        if i%x == 0 {
            fmt.Printf("%d ", arr[i])
        }
    }
    fmt.Println()

    fmt.Print("Masukkan indeks yang akan dihapus: ")
    fmt.Scan(&x)
    if x >= 0 && x < n {
        for i := x; i < n-1; i++ {
            arr[i] = arr[i+1]
        }
        n--

        fmt.Print("Isi array setelah penghapusan: ")
        for i := 0; i < n; i++ {
            fmt.Printf("%d ", arr[i])
        }
        fmt.Println()
    } else {
        fmt.Println("Indeks tidak valid!")
    }

    sum := 0
    for i := 0; i < n; i++ {
        sum += arr[i]
    }
    rataRata := float64(sum) / float64(n)
    fmt.Printf("Rata-rata: %.2f\n", rataRata)

    var simpanganBaku float64
    for i := 0; i < n; i++ {
        simpanganBaku += math.Pow(float64(arr[i]) - rataRata, 2)
    }
    simpanganBaku = math.Sqrt(simpanganBaku / float64(n))
    fmt.Printf("Simpangan baku: %.2f\n", simpanganBaku)

    fmt.Print("Masukkan bilangan untuk menghitung frekuensinya: ")
    fmt.Scan(&x)
    frek := 0
    for i := 0; i < n; i++ {
        if arr[i] == x {
            frek++
        }
    }
    fmt.Printf("Frekuensi bilangan %d: %d\n", x, frek)

    fmt.Println("Program selesai, Code By Kepin.")
}