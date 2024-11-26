package main

import "fmt"

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func soal1_selectionsort() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat: ")
	fmt.Scan(&n)

	if n <= 0 || n >= 1000 {
		fmt.Println("Jumlah daerah kerabat tidak boleh kurang dari 1 atau lebih dari 1000")
		return
	}

	hasil := make([][]int, n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("\nMasukkan jumlah rumah kerabat di daerah %d: ", i+1)
		fmt.Scan(&m)

		if m <= 0 || m >= 1000000 {
			fmt.Println("Jumlah rumah kerabat tidak boleh kurang dari 1 atau lebih dari 1000000")
			return
		}

		fmt.Print("Masukkan nomor rumah kerabat: ")
		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		selectionSort(rumah)
		hasil[i] = rumah
	}
	fmt.Println("\nBerikut hasil data setelah diurutkan:")
	for i, terkecil := range hasil {
		fmt.Printf("Daerah %d: ", i+1)
		for j, num := range terkecil {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(num)
		}
		fmt.Println()
	}
}