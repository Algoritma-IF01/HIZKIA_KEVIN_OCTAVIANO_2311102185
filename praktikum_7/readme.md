# <h1 align="center">Laporan Praktikum_7 - Modul 12 & 13 - Pengurutan Data </h1>

<p align="center">Hizkia Kevin Octaviano - 2311102185</p>

## Features

- [Guided](#guided)
1. [contoh program latihan_1](#1-program-latihan_1)
2. [contoh program latihan_2](#2-program-latihan_2)
3. [contoh program latihan_3](#3-program-latihan_3)
4. [contoh program latihan_4](#4-program-latihan_4)

- [Unguided](#unguided)
1. [program soal_1](#1-program-soal1_selectionsort)
2. [program soal_2](#2-program-soal2_selectionsort)
3. [program soal_3](#3-program-soal1_insertionsort)
4. [program soal_4](#4-program-soal2_insertionsort)


## Guided

### 1. program latihan_1

**Source Code :**

```GO
package main

import "fmt"

type arrInt [4321]int

func selectionSort1(T *arrInt, n int) {
	/* I.S. terdefinisi array T yang berisi n bilangan bulat
	   F.S. array T terurut secara ascending atau membesar dengan SELECTION SORT */
	for i := 0; i < n-1; i++ {
		// Inisialisasi indeks minimum
		idx_min := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[idx_min] {
				idx_min = j
			}
		}
		// Tukar elemen T[i] dengan T[idx_min] jika perlu
		if idx_min != i {
			T[i], T[idx_min] = T[idx_min], T[i]
		}
	}
}

func main() {
	// Contoh penggunaan
	var T arrInt
	n := 5
	T[0], T[1], T[2], T[3], T[4] = 64, 34, 25, 12, 22

	fmt.Println("Array sebelum diurutkan:", T[:n])
	selectionSort1(&T, n)
	fmt.Println("Array setelah diurutkan:", T[:n])
}
```

**Screenshot Output :**
![screenshot_output_latihan1 go](https://github.com/user-attachments/assets/febcf2cf-eaca-481d-958a-df5e1d193f1c)

### 2. program latihan_2

**Source Code :**

```GO
package main

import "fmt"

type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs [2023]mahasiswa

func selectionSort2(T *arrMhs, n int) {
	/* I.S. terdefinisi array T yang berisi n data mahasiswa
	   F.S. array T terurut secara ascending berdasarkan ipk dengan
	   menggunakan algoritma SELECTION SORT */

	var idx_min int
	var temp mahasiswa

	for i := 0; i < n-1; i++ {
		// Inisialisasi indeks minimum
		idx_min = i

		// Cari elemen dengan IPK terkecil di subarray [i+1, n-1]
		for j := i + 1; j < n; j++ {
			if T[j].ipk < T[idx_min].ipk {
				idx_min = j
			}
		}

		// Tukar elemen di indeks i dengan elemen di idx_min jika perlu
		if idx_min != i {
			temp = T[i]
			T[i] = T[idx_min]
			T[idx_min] = temp
		}
	}
}

func main() {
	// Contoh data mahasiswa
	var T arrMhs
	T[0] = mahasiswa{"Alice", "123", "A", "Teknik Informatika", 3.8}
	T[1] = mahasiswa{"Bob", "124", "B", "Sistem Informasi", 3.2}
	T[2] = mahasiswa{"Charlie", "125", "A", "Teknik Informatika", 3.5}
	T[3] = mahasiswa{"Diana", "126", "B", "Sistem Informasi", 3.9}
	n := 4

	fmt.Println("Data mahasiswa sebelum diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}

	selectionSort2(&T, n)

	fmt.Println("\nData mahasiswa setelah diurutkan berdasarkan IPK:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}
}
```

**Screenshot Output :**
![screenshot_output_latihan2 go](https://github.com/user-attachments/assets/e927625e-51b1-4713-8bfd-1fb4360b2bff)

### 3. program latihan_3

**Source Code :**

```GO
package main

import "fmt"

type arrInt [4321]int

func insertionSort1(T *arrInt, n int) {
	/* I.S. terdefinisi array T yang berisi n bilangan bulat
	   F.S. array T terurut secara mengecil (descending) dengan INSERTION SORT */
	var temp, i, j int

	for i = 1; i < n; i++ {
		temp = T[i] // Simpan elemen ke-i
		j = i       // Inisialisasi indeks pembanding

		// Geser elemen-elemen sebelumnya yang lebih kecil dari temp
		for j > 0 && temp > T[j-1] {
			T[j] = T[j-1]
			j--
		}

		// Tempatkan temp pada posisi yang sesuai
		T[j] = temp
	}
}

func main() {
	// Contoh penggunaan
	var T arrInt
	n := 5
	T[0], T[1], T[2], T[3], T[4] = 22, 12, 34, 64, 25

	fmt.Println("Array sebelum diurutkan:", T[:n])
	insertionSort1(&T, n)
	fmt.Println("Array setelah diurutkan secara descending:", T[:n])
}
```

**Screenshot Output :**
![screenshot_output_latihan3 go](https://github.com/user-attachments/assets/c4ffed60-cc21-4f0a-aaad-4337a3f38159)

### 4. program latihan_4

**Source Code :**

```GO
package main

import "fmt"

type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs [2023]mahasiswa

func insertionSort2(T *arrMhs, n int) {
	/* I.S. terdefinisi array T yang berisi n data mahasiswa
	   F.S. array T terurut secara mengecil (descending) berdasarkan nama
	   dengan menggunakan algoritma INSERTION SORT */
	var temp mahasiswa
	var i, j int

	for i = 1; i < n; i++ {
		temp = T[i] // Simpan elemen ke-i
		j = i       // Inisialisasi indeks pembanding

		// Geser elemen-elemen sebelumnya
		for j > 0 && temp.nama > T[j-1].nama {
			T[j] = T[j-1]
			j--
		}

		// Tempatkan temp pada posisi yang sesuai
		T[j] = temp
	}
}

func main() {
	// Contoh data mahasiswa
	var T arrMhs
	T[0] = mahasiswa{"Charlie", "125", "A", "Teknik Informatika", 3.5}
	T[1] = mahasiswa{"Alice", "123", "A", "Teknik Informatika", 3.8}
	T[2] = mahasiswa{"Bob", "124", "B", "Sistem Informasi", 3.2}
	T[3] = mahasiswa{"Diana", "126", "B", "Sistem Informasi", 3.9}
	n := 4

	fmt.Println("Data mahasiswa sebelum diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}

	insertionSort2(&T, n)

	fmt.Println("\nData mahasiswa setelah diurutkan berdasarkan nama (descending):")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}
}
```

**Screenshot Output :**
![screenshot_output_latihan4 go](https://github.com/user-attachments/assets/5f136fbc-a91f-4603-89a4-8de016378f20)

## Unguided

### 1. program soal1_selectionsort

**Source Code :**

```GO
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
```

**Screenshot Output :**
![screenshot_output_soal1_selectionsort go](https://github.com/user-attachments/assets/66e7bd62-faed-4205-ae6c-3c5bd4ef542a)

### 2. program soal2_selectionsort

**Source Code :**

```GO
package main

import "fmt"

func selectionSrt(arr []int) {
    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            if arr[j] < arr[i] {
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }
}

func soal2_selectionsort() {
    var n int
    fmt.Print("Masukkan jumlah daerah kerabat: ")
    fmt.Scan(&n)

    if n <= 0 || n >= 1000 {
        fmt.Println("Jumlah daerah kerabat tidak boleh kurang dari 1 atau lebih dari 1000")
        return
    }

    hasilGanjil := make([][]int, n)
    hasilGenap := make([][]int, n)

    for i := 0; i < n; i++ {
        var m int
        fmt.Printf("\nMasukkan jumlah rumah kerabat di daerah %d: ", i+1)
        fmt.Scan(&m)

        if m <= 0 || m >= 1000000 {
            fmt.Println("Jumlah rumah kerabat tidak boleh kurang dari 1 atau lebih dari 1000000")
            return
        }

        fmt.Print("Masukkan nomor rumah kerabat: ")
        var ganjil, genap []int
        for j := 0; j < m; j++ {
            var num int
            fmt.Scan(&num)
            if num%2 == 0 {
                genap = append(genap, num)
            } else {
                ganjil = append(ganjil, num)
            }
        }
        
        selectionSrt(ganjil)
        selectionSrt(genap)
        hasilGanjil[i] = ganjil
        hasilGenap[i] = genap
    }

    fmt.Println("\nBerikut hasil data setelah diurutkan:")
    for i := 0; i < n; i++ {
        fmt.Printf("Daerah %d: ", i+1)
        for j, num := range hasilGanjil[i] {
            if j > 0 {
                fmt.Print(" ")
            }
            fmt.Print(num)
        }
        if len(hasilGanjil[i]) > 0 && len(hasilGenap[i]) > 0 {
            fmt.Print(" ")
        }
        for j, num := range hasilGenap[i] {
            if j > 0 {
                fmt.Print(" ")
            }
            fmt.Print(num)
        }
        fmt.Println()
    }
}
```

**Screenshot Output :**
![screenshot_output_soal2_selectionsort go](https://github.com/user-attachments/assets/bac506a6-2d3f-445b-af32-a677b3d5d71d)

### 3. program soal1_insertionsort

**Source Code :**

```GO
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
```

**Screenshot Output :**
![screenshot_output_soal1_insertionsort go](https://github.com/user-attachments/assets/459bf8d7-6e53-467f-9dd7-05e49facc69f)

### 4. program soal2_insertionsort

**Source Code :**

```GO
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
```

**Screenshot Output :**
![screenshot_output_soal2_insertionsort go](https://github.com/user-attachments/assets/be418859-0189-4748-a467-71d05272b44b)

![screenshot_output_soal2_insertionsort go_2](https://github.com/user-attachments/assets/1fca1860-0a38-4951-b158-8b26aafe0639)

