# <h1 align="center">Laporan Praktikum_6 - Modul 10 - Struck dan Array </h1>

<p align="center">Hizkia Kevin Octaviano - 2311102185</p>

## Features

- [Guided](#guided)
1. [contoh program latihan_1](#1-program-latihan_1)
2. [contoh program latihan_2](#2-program-latihan_2)

- [Unguided](#unguided)
1. [program soal_1](#1-program-soal_1)
2. [program soal_2](#2-program-soal_2)
3. [program soal_3](#3-program-soal_3)


## Guided

### 1. program latihan_1

**Source Code :**

```GO
package main

import "fmt"

type arrInt [2023]int

// Fungsi untuk mencari indeks dari nilai terkecil
func terkecil_2(tabInt arrInt, n int) int {
	var idx int = 0 // indeks data pertama
	var j int = 1   // pencarian dimulai dari data kedua
	for j < n {
		if tabInt[idx] > tabInt[j] { // cek apakah tabInt[j] lebih kecil dari tabInt[idx]
			idx = j // update idx ke indeks baru yang nilainya lebih kecil
		}
		j = j + 1
	}
	return idx // mengembalikan indeks dari nilai terkecil
}

func main() {
	var n int
	var data arrInt

	// Input jumlah elemen N
	fmt.Print("Masukkan jumlah elemen (N <= 2023): ")
	fmt.Scan(&n)

	// Validasi N agar tidak melebihi kapasitas array
	if n <= 0 || n > 2023 {
		fmt.Println("Jumlah elemen harus antara 1 dan 2023")
		return
	}

	// Input elemen-elemen array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	// Panggil fungsi untuk mencari indeks nilai terkecil
	idxTerkecil := terkecil_2(data, n)
	fmt.Printf("Indeks nilai terkecil: %d\n", idxTerkecil)
	fmt.Printf("Nilai terkecil: %d\n", data[idxTerkecil])
}
```

**Screenshot Output :**
![screenshot_output_latihan1 go](https://github.com/user-attachments/assets/f9ef0d6a-8e61-4a7b-ba80-b771b3a6a90f)

### 2. program latihan_2

**Source Code :**

```GO
package main

import "fmt"

// Mendefinisikan tipe data mahasiswa
type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

// Mendefinisikan array mahasiswa dengan kapasitas 2023
type arrMhs [2023]mahasiswa

// Fungsi untuk mencari indeks mahasiswa dengan IPK tertinggi
func IPK_2(T arrMhs, n int) int {
	// idx menyimpan indeks mahasiswa dengan IPK tertinggi sementara
	var idx int = 0
	var j int = 1
	for j < n {
		if T[idx].ipk < T[j].ipk {
			idx = j
		}
		j = j + 1
	}
	return idx
}

func main() {
	var n int
	var data arrMhs

	// Input jumlah mahasiswa
	fmt.Print("Masukkan jumlah mahasiswa (N <= 2023): ")
	fmt.Scan(&n)

	// Validasi jumlah mahasiswa
	if n <= 0 || n > 2023 {
		fmt.Println("Jumlah mahasiswa harus antara 1 dan 2023")
		return
	}

	// Input data mahasiswa
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data mahasiswa ke-%d\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&data[i].nama)
		fmt.Print("NIM: ")
		fmt.Scan(&data[i].nim)
		fmt.Print("Kelas: ")
		fmt.Scan(&data[i].kelas)
		fmt.Print("Jurusan: ")
		fmt.Scan(&data[i].jurusan)
		fmt.Print("IPK: ")
		fmt.Scan(&data[i].ipk)
	}

	// Panggil fungsi untuk mencari indeks mahasiswa dengan IPK tertinggi
	idxTertinggi := IPK_2(data, n)

	// Tampilkan data mahasiswa dengan IPK tertinggi
	fmt.Println("\nMahasiswa dengan IPK tertinggi:")
	fmt.Printf("Nama    : %s\n", data[idxTertinggi].nama)
	fmt.Printf("NIM     : %s\n", data[idxTertinggi].nim)
	fmt.Printf("Kelas   : %s\n", data[idxTertinggi].kelas)
	fmt.Printf("Jurusan : %s\n", data[idxTertinggi].jurusan)
	fmt.Printf("IPK     : %.2f\n", data[idxTertinggi].ipk)
}
```

**Screenshot Output :**
![screenshot_output_latihan2 go](https://github.com/user-attachments/assets/8f15c6ce-c16b-4ae3-9fbe-f7dd0657d18c)

## Unguided

### 1. program soal_1

**Source Code :**

```GO
package main

import "fmt"

func soal_1() {
	var N int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000")
		return
	}

	weights := make([]float64, N)
	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < N; i++ {
		fmt.Scan(&weights[i])
	}

	minWeight, maxWeight := weights[0], weights[0]
	for _, weight := range weights {
		if weight < minWeight {
			minWeight = weight
		}
		if weight > maxWeight {
			maxWeight = weight
		}
	}

	fmt.Printf("Berat anak kelinci terkecil: %.2f\n", minWeight)
	fmt.Printf("Berat anak kelinci tertinggi: %.2f\n", maxWeight)
}
```

**Screenshot Output :**
![screenshot_output_soal1 go](https://github.com/user-attachments/assets/aa591df3-b3d9-4afc-a753-4263b09958ac)

### 2. program soal_2

**Source Code :**

```GO
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
```

**Screenshot Output :**
![screenshot_output_soal2 go](https://github.com/user-attachments/assets/68bf93ea-1f4c-43d9-a0ee-a743eed115cb)

### 3. program soal_3

**Source Code :**

```GO
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
```

**Screenshot Output :**
![screenshot_output_soal_3 go](https://github.com/user-attachments/assets/03200bcc-72c0-4dec-8938-25f00ae3b32c)
