# <h1 align="center">Laporan Praktikum_5 - Modul 7 - Struck dan Array </h1>

<p align="center">Hizkia Kevin Octaviano - 2311102185</p>

## Features

- [Guided](#guided)
1. [contoh program latihan_1](#1-program-latihan_1)
2. [contoh program latihan_2](#2-program-latihan_2)
3. [contoh program latihan_3](#3-program-latihan_3)
4. [contoh program latihan_4](#4-program-latihan_4)
5. [contoh program latihan_5](#5-program-latihan_5)

- [Unguided](#unguided)
1. [program soal_2](#1-program-soal_2)
2. [program soal_3](#2-program-soal_3)
3. [program soal_4](#3-program-soal_4)


## Guided

### 1. program latihan_1

**Source Code :**

```GO
package main

import "fmt"

type bilangan int
type pecahan float64

func main() {
	var a, b bilangan
	var hasil pecahan
	a = 9
	b = 5
	hasil = pecahan(a) / pecahan(b)
	fmt.Println(hasil)

}
```

**Screenshot Output :**
![screenshot_output_latihan1 go](https://github.com/user-attachments/assets/1af892ac-4b11-481b-99d1-b223e6528a36)

### 2. program latihan_2

**Source Code :**

```GO
package main

import "fmt"

type waktu struct {
 jam, menit, detik int
}

func main(){
 	var wParkir, wPulang, durasi waktu
	var dParkir, dPulang, lParkir int
	fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik)
 	fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik)
 	dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600
 	dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600
 	lParkir = dPulang - dParkir
 	durasi.jam = lParkir / 3600
 	durasi.menit = lParkir % 3600 / 60
 	durasi.detik = lParkir % 3600 % 60
 	fmt.Printf("Lama parkir: %d jam %d menit %d detik",
 	durasi.jam, durasi.menit, durasi.detik)
}
```

**Screenshot Output :**
![screenshot_output_latihan2 go](https://github.com/user-attachments/assets/45579705-fb48-4292-abad-46c084abed7d)

### 3. program latihan_3

**Source Code :**

```GO
package main

import "fmt"

// Definisi tipe CircType
type CircType struct {
	Radius float64
}

// Definisi tipe NewType
type NewType struct {
	Name string
}

func main() {
	var (
		// Array arr mempunyai 73 elemen, masing-masing bertipe CircType
		arr [73]CircType

		// Array buf dengan 5 elemen, dengan nilai awal 7,3,5,2,11
		buf = [5]byte{7, 3, 5, 2, 11}

		// Array mhs berisi 2000 elemen NewType
		mhs [2000]NewType

		// Array dua dimensi rec berisi nilai float64
		rec [20][40]float64
	)

	// Mengisi beberapa elemen contoh
	arr[0] = CircType{Radius: 5.5}
	mhs[0] = NewType{Name: "Alice"}
	rec[0][0] = 3.14

	// Contoh penggunaan variabel
	fmt.Println("arr[0]:", arr[0])
	fmt.Println("buf:", buf)
	fmt.Println("mhs[0]:", mhs[0])
	fmt.Println("rec[0][0]:", rec[0][0])
}
```

**Screenshot Output :**
![screenshot_output_latihan3 go](https://github.com/user-attachments/assets/c913bc37-c4eb-44ff-a148-2301c6536dfc)

### 4. program latihan_4

**Source Code :**

```GO
package main

import "fmt"

func main() {
	// Membuat map dengan tipe string sebagai kunci dan integer sebagai nilai
	scores := map[string]int{
		"John": 90,
		"Anne": 85,
	}

	// Mengambil nilai dari kunci
	johnScore := scores["John"]
	fmt.Println("Nilai John:", johnScore)

	// Mengganti nilai dari kunci yang ada
	scores["John"] = 95
	fmt.Println("Nilai John setelah diganti:", scores["John"])

	// Menambah kunci baru dengan nilai-
	scores["David"] = 88
	fmt.Println("Nilai David ditambahkan:", scores["David"])

	// Menghapus kunci dari map
	delete(scores, "Anne")
	fmt.Println("Map setelah kunci 'Anne' dihapus:", scores)

	// Mengecek apakah kunci ada dalam map
	if score, ada := scores["David"]; ada {
		fmt.Println("Nilai David ditemukan:", score)
	} else {
		fmt.Println("Nilai David tidak ditemukan")
	}
}
```

**Screenshot Output :**
![screenshot_output_latihan4 go](https://github.com/user-attachments/assets/7f177754-e83d-4c08-a7dd-d2cd57f80cae)

### 5. program latihan_5

**Source Code :**

```GO
package main

import (
	"fmt"
	"math"
)

// Definisi tipe bentukan untuk titik
type Titik struct {
	x int
	y int
}

// Definisi tipe bentukan untuk lingkaran
type Lingkaran struct {
	center Titik
	radius int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(p Titik, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func didalam(c Lingkaran, p Titik) bool {
	return jarak(p, c.center) < float64(c.radius)
}

func main() {
	var (
		// Mengambil input untuk lingkaran 1
		lingkaran1 Lingkaran
		// Mengambil input untuk lingkaran 2
		lingkaran2 Lingkaran
		// Mengambil input untuk titik sembarang
		point Titik
	)

	// Input untuk lingkaran 1 (cx, cy, r)
	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 1 (cx cy r):")
	fmt.Scan(&lingkaran1.center.x, &lingkaran1.center.y, &lingkaran1.radius)

	// Input untuk lingkaran 2 (cx, cy, r)
	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 2 (cx cy r):")
	fmt.Scan(&lingkaran2.center.x, &lingkaran2.center.y, &lingkaran2.radius)

	// Input untuk titik sembarang (x, y)
	fmt.Println("Masukkan koordinat titik sembarang (x y):")
	fmt.Scan(&point.x, &point.y)

	// Mengecek posisi titik terhadap kedua lingkaran
	inLingkaran1 := didalam(lingkaran1, point)
	inLingkaran2 := didalam(lingkaran2, point)

	if inLingkaran1 && inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```

**Screenshot Output :**
![screenshot_output_latihan5 go](https://github.com/user-attachments/assets/b57e56b0-e6e5-44cd-b6d8-46a3f9e5a302)

## Unguided

### 1. program soal_2

**Source Code :**

```GO
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
```

**Screenshot Output :**
![screenshot_output_soal2 go](https://github.com/user-attachments/assets/c4c160af-c103-48b6-bf55-5832830f7325)

### 2. program soal_3

**Source Code :**

```GO
package main

import "fmt"

func soal_3() {
    var klubA, klubB string
    var skorA, skorB int
    var pemenang []string

    pertandingan := 1

    fmt.Print("Klub A: ")
    fmt.Scanln(&klubA)
    fmt.Print("Klub B: ")
    fmt.Scanln(&klubB)

    for {
        fmt.Printf("Pertandingan %d : ", pertandingan)
        fmt.Scanln(&skorA, &skorB)

            if skorA < 0 || skorB < 0 {
            break
            }

            if skorA > skorB {
                pemenang = append(pemenang, klubA)
			} else if skorA < skorB {
				pemenang = append(pemenang, klubB)
            } else {
                pemenang = append(pemenang, "Draw")
            }
            pertandingan++
        }

	for i := 0; i < len(pemenang); i++ {
			fmt.Printf("Hasil: %d. %s\n", i+1, pemenang[i])
	}
	fmt.Println("Pertandingan selesai")
}
```

**Screenshot Output :**
![screenshot_output_soal3 go](https://github.com/user-attachments/assets/28d462e7-3051-4aed-8041-b80100164fca)

### 3. program soal_4

**Source Code :**

```GO
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	*n = 0
	fmt.Print("Teks : ")
	for *n < NMAX {
		var ch rune
		fmt.Scanf("%c", &ch)
		if ch == '.' {
			break
		}
		if ch != ' ' {
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func soal_4() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	balikanArray(&tab, m)
	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)

	isPalindrom := palindrom(tab, m)
	fmt.Printf("Palindrom? %v\n", isPalindrom)
}
```

**Screenshot Output :**
![screenshot_output_soal4 go](https://github.com/user-attachments/assets/12f0578e-b808-4f31-a3f0-4a195d62518d)
