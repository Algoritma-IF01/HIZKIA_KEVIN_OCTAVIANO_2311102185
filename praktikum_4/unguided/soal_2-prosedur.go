package main

import (
	"fmt"
	"strings"
)

func hitungSkor(waktu [8]int) (int, int) {
	soal, skor := 0, 0
	for i := 0; i < 8; i++ {
		if waktu[i] < 301 {
			soal++
			skor += waktu[i]
		}
	}
	return soal, skor
}

func soal_2_prosedur() {
	var (
		nama, pemenang string
		waktu[8] int
		maxSoal, minSkor int
	)

	for {
		fmt.Print("Masukan nama peserta (atau 'selesai' untuk berhenti): ")
		
		fmt.Scanln()
		fmt.Scanln(&nama)
		if strings.ToLower(nama) == "selesai" {
			break
		}

		fmt.Println("Masukan waktu pengerjaan (dalam menit) untuk 8 soal:")
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		soal, skor := hitungSkor(waktu)

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			pemenang = nama
			maxSoal = soal
			minSkor = skor
		}
	}

	if pemenang != "" {
		fmt.Printf("\nNama pemenang: %s\n", pemenang)
		fmt.Printf("Jumlah soal yang diselesaikan: %d\n", maxSoal)
		fmt.Printf("Nilai (total waktu): %d\n", minSkor)
	} else {
		fmt.Println("Tidak ada peserta.")
	}
}