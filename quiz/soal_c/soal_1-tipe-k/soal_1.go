package main

import "fmt"

func main() {
	var jumlahPeserta int
	fmt.Scan(&jumlahPeserta)

	var jumlahEsTebak, jumlahEsCendol, jumlahLamang int

	for i := 0; i < jumlahPeserta; i++ {
		var nomorKartu int
		fmt.Scan(&nomorKartu)

		var digitPertama, digitTerakhir int
		digitPertama = nomorKartu % 10
		for nomorKartu > 0 {
			digitTerakhir = nomorKartu % 10
			if digitTerakhir != digitPertama {
				break
			}
			nomorKartu /= 10
		}

		if nomorKartu == 0 && digitPertama%2 == 1 {
			fmt.Println("Es Tebak")
			jumlahEsTebak++
		} else if digitTerakhir%2 == 0 {
			fmt.Println("Es Cendol")
			jumlahEsCendol++
		} else {
			fmt.Println("Lamang")
			jumlahLamang++
		}
	}

	fmt.Println("\nJumlah Es Tebak:", jumlahEsTebak)
	fmt.Println("Jumlah Es Cendol:", jumlahEsCendol)
	fmt.Println("Jumlah Lamang:", jumlahLamang)
}
