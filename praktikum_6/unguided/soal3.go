package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func soal_3() {
	var n int
	var arrBerat arrBalita
	var bMin, bMax float64

	for {
		fmt.Print("Masukan banyak data berat balita: ")
		fmt.Scan(&n)
		if n > 0 && n <= 100 {
			break
		}
		fmt.Println("Jumlah data harus antara 1 dan 100. Silakan masukkan ulang.")
	}

	for i := 0; i < n; i++ {
		for {
			fmt.Printf("Masukan berat balita ke-%d: ", i+1)
			fmt.Scan(&arrBerat[i])
			if arrBerat[i] > 0 {
				break
			}
			fmt.Println("Berat balita harus positif. Silakan masukkan ulang.")
		}
	}

	hitungMinMax(arrBerat, n, &bMin, &bMax)

	rataRata := rerata(arrBerat, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}