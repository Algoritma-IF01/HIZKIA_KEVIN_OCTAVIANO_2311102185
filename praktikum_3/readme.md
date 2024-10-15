# <h1 align="center">Laporan Praktikum_3 - Modul 2 - Struktur Kontrol Perulangan & Percabangan</h1>

<p align="center">Hizkia Kevin Octaviano - 2311102185</p>

## Features

- [Guided](#guided)
1. [Program "greetings"](#1-program-greetings)
2. [Program "hipotenusa"](#2-program-hipotenusa)

- [Unguided](#unguided)
1. [Program "Input string"](#1-program-input-string)
2. [Program "Tahun kabisat"](#2-program-tahun-kabisat)
3. [Program Menghitung Volume Bola & Luas Bola](#3-program-menghitung-volume-bola--luas-bola)
4. [Program Konversi Suhu](#4-program-konversi-suhu)
5. [Program "ASCII"](#5-program-ascii)

## Guided

### 1. Soal_latihan1

**Source Code :**

```GO
package main

import "fmt"

func main (){ 
	urutanBenar := []string{"merah","kuning","hijau","ungu"}
	hasil := true

	for i := 1; i <= 5; i++ {
		var warna1, warna2, warna3, warna4 string
		fmt.Printf("Percobaan %d\n", i)
		fmt.Print("Masukan warna pertama : ")
		fmt.Scanln(&warna1)
		fmt.Print("Masukan warna kedua : ")
		fmt.Scanln(&warna2)
		fmt.Print("Masukan warna ketiga : ")
		fmt.Scanln(&warna3)
		fmt.Print("Masukan warna keempat : ")
		fmt.Scanln(&warna4)

		if warna1 != urutanBenar[0] || warna2 != urutanBenar[1] || warna3 != urutanBenar[2] || warna4 != urutanBenar[3]{
			hasil = false
		}
	}
	fmt.Println("Berhasil : ", hasil)
}

```

**Screenshot Output :**
![image](https://github.com/user-attachments/assets/40d8a01b-3831-4917-92e6-138c92c10a52)

### 2. Soal_latihan2

**Source Code :**

```GO
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pita string
	var bungaCount int

	for {
		fmt.Printf("Bunga %d: ", bungaCount+1)
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "selesai" {
			break
		}

		if pita == "" {
			pita = input
		} else {
			pita += " – " + input
		}
		bungaCount++
	}

	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", bungaCount)
}
```

**Screenshot Output :**
![screenshot_output_guided2](https://github.com/user-attachments/assets/e2226148-fd2d-48da-8297-265f84d356cd)

## Unguided

### 1. soal_3-perulangan

**Source Code :**

```GO
package main

import "fmt"

func main() {
    var beratKantong1, beratKantong2 float64

    for {
        fmt.Print("Masukan berat belanjaan di kedua kantong: ")
        fmt.Scan(&beratKantong1, &beratKantong2)

        if beratKantong1 < 0 || beratKantong2 < 0 {
            break
        }

        totalBerat := beratKantong1 + beratKantong2
        if totalBerat > 150 {
            break
        }

        selisih := beratKantong1 - beratKantong2
        if selisih < 0 {
            selisih = -selisih
        }

        fmt.Println("Sepeda motor pak Andi akan oleng:", selisih >= 9)
    }

    fmt.Println("Proses selesai.")
}
```

**Screenshot Output :**
![screenshot_output_unguided_soal-3-perulangan](https://github.com/user-attachments/assets/b0739622-8a03-4644-bc5f-36e15c1dd16a)

### 2. soal_4-perulangan

**Source Code :**

```GO
package main

import "fmt"

func main() {
	var K int
	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	fk := float64((4*K+2)*(4*K+2)) / float64((4*K+1)*(4*K+3))
	fmt.Printf("Nilai f(K) = %.10f\n", fk)

	hasil := 1.0
	for k := 0; k <= K; k++ {
		hasil *= float64((4*k+2)*(4*k+2)) / float64((4*k+1)*(4*k+3))
	}
	fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}
```

**Screenshot Output :**
![screenshot_output_unguided_soal-4-perulangan](https://github.com/user-attachments/assets/be2ec2e1-5e3c-42f3-9187-3978a83e5ba6)

### 3. soal_1-percabangan

**Source Code :**

```GO
package main

import "fmt"

func main() {
	var beratParsel int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratParsel)

	beratKg := beratParsel / 1000
	sisaGram := beratParsel % 1000

	biayaKg := beratKg * 10000
	var biayaSisaGram int

	if beratKg > 10 {
		biayaSisaGram = 0
	} else if sisaGram >= 500 {
		biayaSisaGram = sisaGram * 5
	} else {
		biayaSisaGram = sisaGram * 15
	}

	totalBiaya := biayaKg + biayaSisaGram

	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisaGram)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}

```

**Screenshot Output :**
![screenshot_output_unguided_soal-1-percabangan](https://github.com/user-attachments/assets/5b0683ad-e48a-40a1-bcfd-2a3ce1419572)

### 3. soal_2-percabangan

**Source Code :**

```GO
package main

import "fmt"

func main() {
	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)

	if nam > 80 {
		nmk = "A"
	}
	if nam > 72.5 {
		nmk = "AB"
	}
	if nam > 65 {
		nmk = "B"
	}
	if nam > 57.5 {
		nmk = "BC"
	}
	if nam > 50 {
		nmk = "C"
	}
	if nam > 40 {
		nmk = "D"
	} else if nam <= 40 {
		nmk = "E"
	}
	fmt.Println("Nilai mata kuliah: ", nmk)
}

```

Jawablah pertanyaan-pertanyaan berikut:

a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi
program tersebut sesuai spesifikasi soal?

Berikut hasil outputnya :

**Screenshot Output :**
![screenshot_output_unguided_soal-2-percabangan](https://github.com/user-attachments/assets/24d582de-647d-495a-be3b-92de3d42ab80)

jawab : Jika nilai `nam` yang diberikan adalah 80.1, keluaran dari program tersebut memang akan menjadi "Nilai mata kuliah: D". Ini terjadi karena program tersebut menjalankan semua pernyataan `if` secara berurutan, dan kondisi terakhir yang terpenuhi adalah `if nam > 40`, yang menghasilkan nilai "D". Meskipun benar bahwa 80.1 menghasilkan output "D", ini tetap tidak sesuai dengan spesifikasi yang diinginkan, yang di mana nilai 80.1 seharusnya mendapatkan nilai mata kuliah yang lebih tinggi.

b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program
seharusnya!

jawab : 

1. Kesalahan utama dalam program asli adalah penggunaan serangkaian pernyataan `if` terpisah, bukan struktur `if-else if` yang berurutan. Ini menyebabkan setiap kondisi diperiksa secara independen, mengakibatkan nilai `nmk` terus diperbarui hingga kondisi terakhir yang terpenuhi.

2. Urutan pemeriksaan dalam program asli sudah benar (dari nilai tertinggi ke terendah), tapi karena menggunakan `if` terpisah, nilai akhir yang dipilih adalah yang terakhir terpenuhi, biasanya merupakan nilai terendah yang masih memenuhi syarat.

3. Program asli tidak memiliki kondisi `else` untuk menangani nilai ≤ 40, melainkan menggunakan `else if nam <= 40`, yang tidak diperlukan karena sudah ada dalam logika sebelumnya.

### Mengapa demikian?

Karena menggunakan if terpisah, setiap kondisi akan diperiksa secara independen. Akibatnya, untuk nilai di atas 80, semua kondisi akan terpenuhi, dan nmk akan terus diubah hingga mencapai kondisi terakhir yang terpenuhi.

### Jelaskan alur program seharusnya!

Alur program yang benar seharusnya dimulai dengan meminta input dari user berupa nilai akhir mata kuliah yang disimpan dalam variabel `nam`. Program kemudian akan memeriksa nilai tersebut dengan serangkaian kondisi berbasis rentang nilai tertentu. Logika kondisi dimulai dengan memeriksa apakah nilai nam > 80 && nam <= 100. Jika kondisi ini benar, maka nilai yang diberikan adalah "A". Jika nilai tidak masuk dalam rentang tersebut, program melanjutkan ke kondisi berikutnya: apakah nilai nam > 72.5 && nam <= 80, dan mendapatkan nilai "AB" jika terpenuhi. Proses ini terus berlanjut untuk setiap rentang nilai, dengan nilai nam > 65 && nam <= 72.5 diberi nilai "B", nilai nam > 57.5 && nam <= 65 diberi nilai "BC", nilai nam > 50 && nam <= 57.5 diberi nilai "C", dan nilai nam > 40 && nam <= 50 diberi nilai "D". Jika semua kondisi di atas tidak terpenuhi (artinya nilai nam <= 40), maka program menetapkan nilai "E". Setelah kondisi terpenuhi, program mencetak nilai hasil akhir tersebut. Struktur kontrol yang digunakan adalah `if-else if` untuk memastikan bahwa hanya satu kondisi yang dieksekusi, sehingga rentang nilai diperiksa secara berurutan dan tepat sesuai aturan penilaian.

c. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya
keluaran yang diperoleh adalah ‘A’, ‘B’, dan ‘D’.

Berikut hasil source code yang sudah di perbaiki :

```GO
package main

import "fmt"

func main() {
	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)

	if nam > 80 && nam <= 100 {
		nmk = "A"
	} else if nam > 72.5 && nam <= 80 {
		nmk = "AB"
	} else if nam > 65 && nam <= 72.5 {
		nmk = "B"
	} else if nam > 57.5 && nam <= 65 {
		nmk = "BC"
	} else if nam > 50 && nam <= 57.5 {
		nmk = "C"
	} else if nam > 40 && nam <= 50 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	fmt.Println("Nilai mata kuliah:", nmk)
}
```

**Screenshot Output :**
![screenshot_output_unguided_soal-2-percabangan_pertanyaan_c](https://github.com/user-attachments/assets/dc2132a5-da15-4ca8-b519-cf67eb045da0)

### 4. soal_3-percabangan

**Source Code :**

```GO
package main

import "fmt"

func main() {
    var bilangan int
    fmt.Print("Bilangan: ")
    fmt.Scanln(&bilangan)

    if bilangan <= 1 {
        fmt.Println("Bilangan harus lebih besar atau sama dengan 1")
        return
    }

	print("Faktor: ")
    hitungfaktor := 0

    for i := 1; i <= bilangan; i++ {
        if bilangan%i == 0 {
            print(i, " ")
            hitungfaktor++
        }
    }

    if hitungfaktor == 2 {
        println("\nPrima: true")
    } else {
        println("\nPrima: false")
    }
}
```

**Screenshot Output :**
![screenshot_output_unguided_soal-3-percabangan](https://github.com/user-attachments/assets/ccad8690-288d-4339-8d75-1ebf73ddbc24)

### 5. soal_3-percabangan

**Source Code :**

```GO
package main

import "fmt"

func main() {
	var beratParsel int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratParsel)

	beratKg := beratParsel / 1000
	sisaGram := beratParsel % 1000

	biayaKg := beratKg * 10000
	var biayaSisaGram int

	if beratKg > 10 {
		biayaSisaGram = 0
	} else if sisaGram >= 500 {
		biayaSisaGram = sisaGram * 5
	} else {
		biayaSisaGram = sisaGram * 15
	}

	totalBiaya := biayaKg + biayaSisaGram

	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisaGram)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}

```

**Screenshot Output :**
![screenshot_output_unguided_soal-3-percabangan](https://github.com/user-attachments/assets/78a52876-d4b8-4a85-a742-02255f6a98ef)
