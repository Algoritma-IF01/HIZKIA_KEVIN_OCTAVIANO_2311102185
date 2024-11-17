package main

import "fmt"

const nMax = 1000

func validasiInput(x, y int) bool {
	return x > 0 && y > 0 && x <= nMax && y <= nMax
}

func inputBeratIkan(x int, beratIkan *[nMax]float64) {
	fmt.Printf("Masukkan berat masing-masing ikan sebanyak %d kali:\n", x)
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}
}

func hitungBeratWadah(x, y int, beratIkan *[nMax]float64, totalWadah *[nMax]float64) (jumlahWadah int, totalBerat float64) {
	jumlahWadah = 0
	totalBerat = 0.0

	for i := 0; i < x; {
		beratWadah := 0.0
		count := 0

		for count < y && i < x {
			beratWadah += beratIkan[i]
			i++
			count++
		}

		totalWadah[jumlahWadah] = beratWadah
		totalBerat += beratWadah
		jumlahWadah++
	}
	return
}

func cetakTotalBeratWadah(jumlahWadah int, totalWadah *[nMax]float64) {
	fmt.Println("Total berat setiap wadah")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Berat wadah %d : %.2f\n", i+1, totalWadah[i])
	}
}

func cetakRataRataPerWadah(jumlahWadah, y int, totalWadah *[nMax]float64) {
	fmt.Println("Rata rata berat ikan di setiap wadah")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Rata rata berat wadah %d : %.2f\n", i+1, totalWadah[i]/float64(y))
	}
}

func cetakRataRataSeluruhWadah(jumlahWadah int, totalBerat float64) {
	if jumlahWadah > 0 {
		rataRataSeluruhWadah := totalBerat / float64(jumlahWadah)
		fmt.Printf("Rata rata seluruh wadah : %.2f\n", rataRataSeluruhWadah)
	}
}

func soal_2() {
	var x, y int
	var beratIkan [nMax]float64
	var totalWadah [nMax]float64

	fmt.Println("Masukkan jumlah ikan dan kapasitas per wadah:")
	fmt.Scan(&x, &y)

	if !validasiInput(x, y) {
		fmt.Println("Input tidak valid. Jumlah ikan dan kapasitas per wadah harus antara 1 hingga", nMax)
		return
	}

	inputBeratIkan(x, &beratIkan)

	jumlahWadah, totalBerat := hitungBeratWadah(x, y, &beratIkan, &totalWadah)

	cetakTotalBeratWadah(jumlahWadah, &totalWadah)
	cetakRataRataPerWadah(jumlahWadah, y, &totalWadah)
	cetakRataRataSeluruhWadah(jumlahWadah, totalBerat)
}