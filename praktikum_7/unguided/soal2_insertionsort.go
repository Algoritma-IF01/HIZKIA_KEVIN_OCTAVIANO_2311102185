package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scanln(n)

	for i := 0; i < *n; i++ {
		fmt.Printf("\nMasukkan data untuk buku ke-%d\n", i+1)

		for {
			fmt.Print("Masukkan ID buku: ")
			fmt.Scanln(&(*pustaka)[i].id)

			duplikat := false
			for j := 0; j < i; j++ {
				if (*pustaka)[j].id == (*pustaka)[i].id {
					duplikat = true
					break
				}
			}

			if duplikat {
				fmt.Println("ID buku sudah digunakan, masukkan ID yang berbeda.")
			} else {
				break
			}
		}

		fmt.Print("Masukkan judul buku: ")
		fmt.Scanln(&(*pustaka)[i].judul)

		fmt.Print("Masukkan penulis buku: ")
		fmt.Scanln(&(*pustaka)[i].penulis)

		fmt.Print("Masukkan penerbit buku: ")
		fmt.Scanln(&(*pustaka)[i].penerbit)

		fmt.Print("Masukkan eksemplar buku: ")
		fmt.Scanln(&(*pustaka)[i].eksemplar)

		fmt.Print("Masukkan tahun buku: ")
		fmt.Scanln(&(*pustaka)[i].tahun)

		for {
			fmt.Print("Masukkan rating buku (1-10): ")
			fmt.Scanln(&(*pustaka)[i].rating)
			if (*pustaka)[i].rating < 1 || (*pustaka)[i].rating > 10 {
				fmt.Println("Rating harus antara 1 sampai 10.")
			} else {
				break
			}
		}
	}
}


func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku.")
		return
	}

	maxRating := pustaka[0].rating
	idx := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
			idx = i
		}
	}

	fmt.Println("\nBuku Terfavorit:")
	fmt.Println("Judul     :", pustaka[idx].judul)
	fmt.Println("Penulis   :", pustaka[idx].penulis)
	fmt.Println("Penerbit  :", pustaka[idx].penerbit)
	fmt.Println("Tahun     :", pustaka[idx].tahun)
	fmt.Println("Eksemplar :", pustaka[idx].eksemplar)
	fmt.Println("Rating    :", pustaka[idx].rating)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		temp := (*pustaka)[i]
		j := i - 1
		for j >= 0 && (*pustaka)[j].rating < temp.rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}
		(*pustaka)[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("\nBuku dengan Rating Tertinggi:")
	maks := 5
	if n < 5 {
		maks = n
	}
	for i := 0; i < maks; i++ {
		fmt.Printf("%d. %s\n", i+1, pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			fmt.Println("\nBuku Ditemukan:")
			fmt.Println("Judul     :", pustaka[mid].judul)
			fmt.Println("Penulis   :", pustaka[mid].penulis)
			fmt.Println("Penerbit  :", pustaka[mid].penerbit)
			fmt.Println("Tahun     :", pustaka[mid].tahun)
			fmt.Println("Eksemplar :", pustaka[mid].eksemplar)
			fmt.Println("Rating    :", pustaka[mid].rating)
			return
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating yang dimaksud.")
}

func soal2_insertionsort() {
	var pustaka DaftarBuku
	var n, rating int

	DaftarkanBuku(&pustaka, &n)
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	fmt.Print("\nMasukkan rating yang ingin dicari: ")
	fmt.Scanln(&rating)
	CariBuku(pustaka, n, rating)
}