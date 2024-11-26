package main

import "fmt"

func insertionSort(arr []int) {
    for i := range arr {
        for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
            arr[j], arr[j-1] = arr[j-1], arr[j]
        }
    }
}

func cekJarak(arr []int) bool {
    jarak := arr[1] - arr[0]
    for i := 1; i < len(arr)-1; i++ {
        if arr[i+1]-arr[i] != jarak {
            return false
        }
    }
    return true
}

func soal1_insertionsort() {
    var data []int
    var angka int

    fmt.Println("Masukkan angka-angka (akhiri dengan angka negatif):")
    for fmt.Scan(&angka); angka >= 0; fmt.Scan(&angka) {
        data = append(data, angka)
    }

    if len(data) == 0 {
        fmt.Println("Data kosong.")
        return
    }

    insertionSort(data)

    fmt.Print("hasil:")
    for _, v := range data {
        fmt.Printf(" %d", v)
    }
    fmt.Println()

    if cekJarak(data) {
        fmt.Printf("status: Data berjarak %d\n", data[1]-data[0])
    } else {
        fmt.Println("status: Data berjarak tidak tetap")
    }
}