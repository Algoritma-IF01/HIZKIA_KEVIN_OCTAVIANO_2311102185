# <h1 align="center">Laporan Praktikum_3 - Modul 3 & 4 - Fungsi dan Prosedur </h1>

<p align="center">Hizkia Kevin Octaviano - 2311102185</p>

## Features

- [Guided](#guided)
1. [contoh program fungsi](#1-contoh-program-fungsi)
2. [contoh program prosedur](#2-contoh-program-prosedur)
3. [contoh program latihan1](#3-contoh-program-latihan1)

- [Unguided](#unguided)
1. [soal_2-fungsi](#1-soal_2-fungsi)
2. [soal_3-fungsi](#2-soal_3-fungsi)
3. [soal_2-prosedur](#3-soal_2-prosedur)
4. [soal_3-prosedur](#4-soal_3-prosedur)

## Guided

### 1. contoh program fungsi

**Source Code :**

```GO
package main

import "fmt"

func main() {
	var a,b int
    fmt.Scan(&a, &b)
    if a >= b {
        fmt.Println(permutasi(a,b))
    }else{
        fmt.Println(permutasi(b,a))
    }
}

func faktorial(n int) int{
    var hasil int = 1
    var i int
    for i = 1; i <= n; i++ {
            hasil = hasil * i
    }
    return hasil
}
func permutasi(n,r int) int {
    return faktorial(n) / faktorial(n-r)
   }
```

**Screenshot Output :**
![Screenshot_281](https://github.com/user-attachments/assets/f74b0a86-d421-4606-8888-452d876d6558)

### 2. contoh program prosedur

**Source Code :**

```GO
package main 

import "fmt"


func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main () {
	fmt.Println("Deret Fibonacci hingga suku ke-10: ")
	for i := 0; i<=10; i++ {
		fmt.Printf("Fibonacci(%d)= %d\n", i, fibonacci(i))
	}
}
```

**Screenshot Output :**
![Screenshot_282](https://github.com/user-attachments/assets/b2f4d493-17fe-456a-9ace-559f7beef51d)

### 3. contoh program latihan1

**Source Code :**

```GO
package main

import "fmt"

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n int, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n int, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int
	
	fmt.Println("Masukkan nilai a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	p1 := permutasi(a, c)
	c1 := kombinasi(a, c)
	
	p2 := permutasi(b, d)
	c2 := kombinasi(b, d)

	fmt.Printf("P(%d, %d) = %d\n", a, c, p1)
	fmt.Printf("C(%d, %d) = %d\n", a, c, c1)
	fmt.Printf("P(%d, %d) = %d\n", b, d, p2)
	fmt.Printf("C(%d, %d) = %d\n", b, d, c2)
}
```

**Screenshot Output :**
![Screenshot_283](https://github.com/user-attachments/assets/ea1704c1-2ca2-46b1-b68f-fd61b9b534bf)

## Unguided

### 1. soal_2-fungsi

**Source Code :**

```GO
package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func fogoh(a int) int {
	return f(g(h(a)))
}

func gohof(b int) int {
	return g(h(f(b)))
}

func hofog(c int) int {
	return h(f(g(c)))
}

func soal_2_fungsi() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Printf("\n(fogog)(%d) = %d\n", a, fogoh(a))
	fmt.Printf("(gohof)(%d) = %d\n", b, gohof(b))
	fmt.Printf("(hofog)(%d) = %d\n", c, hofog(c))
}
```

**Screenshot Output :**
![Screenshot_284](https://github.com/user-attachments/assets/6060cae1-1890-43a3-887d-9dbe61e2d25f)

### 2. soal_3-fungsi

**Source Code :**

```GO
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d int) float64 {
	return math.Sqrt(float64((a-c)*(a-c) + (b-d)*(b-d)))
}

func didalam(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func soal_3_fungsi() {
	var cx1, cy1, r1 int
	fmt.Scan(&cx1, &cy1, &r1)

	var cx2, cy2, r2 int
	fmt.Scan(&cx2, &cy2, &r2)

	var x, y int
	fmt.Scan(&x, &y)

	diLingkaran1 := didalam(cx1, cy1, r1, x, y)
	diLingkaran2 := didalam(cx2, cy2, r2, x, y)

	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```

**Screenshot Output :**
![Screenshot_285](https://github.com/user-attachments/assets/5a2f4552-e3c4-4a3d-8dad-b63069c06405)

### 3. soal_2-prosedur

**Source Code :**

```GO
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
```

**Screenshot Output :**
![Screenshot_286](https://github.com/user-attachments/assets/d64d3f01-2567-4a84-92bf-5c9f444e74bf)

### 4. soal_3-prosedur

**Source Code :**

```GO
package main

import "fmt"

func cetakDeret(n int) {
    for n != 1 {
        fmt.Printf("%d ", n)
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 3*n + 1
        }
    }
    fmt.Printf("1\n")
}

func soal_3_prosedur() {
    var n int
    fmt.Scan(&n)
    if n > 0 && n < 1000000 {
        cetakDeret(n)
    }
}
```

**Screenshot Output :**
![Screenshot_287](https://github.com/user-attachments/assets/70097e13-10a4-4de7-a68e-731ad5c0011b)
