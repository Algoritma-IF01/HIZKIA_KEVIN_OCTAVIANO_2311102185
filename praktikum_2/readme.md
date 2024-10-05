# <h1 align="center">Laporan Praktikum_2 - Modul 2 - Review Struktur Kontrol</h1>

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

### 1. Program "greetings"

**Source Code :**

```GO
package main

import "fmt"

func main() {
	var greetings = "Selamat datang di dunia DAP"
	var a, b int

	fmt.Println(greetings)
	fmt.Scanln(&a, &b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}

```

**Screenshot Output :**
![image](https://github.com/user-attachments/assets/c5cd7a8c-6a54-4902-8138-3ab0d2ec03a3)

### 2. Program "hipotenusa"

**Source Code :**

```GO
package main

import "fmt"

func main(){
	var a, b, c float64
	var hipotenusa bool

	fmt.Print("Masukkan nilai A: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai B: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan nilai C: ")
	fmt.Scanln(&c)
	hipotenusa = (c*c) == 	(a*a + b*b)
	fmt.Println( "Sisi c adalah hipotenusa segitiga a, b dan c: ", hipotenusa)
}
```

**Screenshot Output :**
![image](https://github.com/user-attachments/assets/8b1f5e27-df3b-4306-b914-52c2f069cbf0)

## Unguided

### 1. Program "Input string"

**Source Code :**

```GO
package main

import "fmt"

func main(){
	var a, b, c float64
	var hipotenusa bool

	fmt.Print("Masukkan nilai A: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai B: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan nilai C: ")
	fmt.Scanln(&c)
	hipotenusa = (c*c) == 	(a*a + b*b)
	fmt.Println( "Sisi c adalah hipotenusa segitiga a, b dan c: ", hipotenusa)
}
```

**Screenshot Output :**
![image](https://github.com/user-attachments/assets/f7fc4ff4-ed21-4e5c-98ed-6af8a8420964)

### 2. Program "Tahun kabisat"

**Source Code :**

```GO
package main

import "fmt"

func main() {
    var tahun int
    fmt.Print("Tahun: ")
    fmt.Scanln(&tahun)

    kabisat := (tahun % 400 == 0) || (tahun % 4 == 0 && tahun % 100 != 0)

    fmt.Printf("Kabisat: %t", kabisat)
}
```

**Screenshot Output :**
![image](https://github.com/user-attachments/assets/9b35105f-dfa4-464e-b00b-f839d8304fad)

### 3. Program Menghitung Volume Bola & Luas Bola

**Source Code :**

```GO
package main

import "fmt"

func main() {
    var r int
    fmt.Print("Jejari: ")
    fmt.Scanln(&r)

    pi := 3.1415926535
    volume := (4.0 / 3.0) * pi * float64(r*r*r)
    luas := 4 * pi * float64(r*r)

    fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f", r, volume, luas)
}
```

**Screenshot Output :**
![image](https://github.com/user-attachments/assets/378d16a9-8d8e-4690-8536-e6e091789aa0)

### 4. Program Konversi Suhu

**Source Code :**

```GO
package main

import "fmt"

func main() {
	var celsius int
	fmt.Print("Temperatur Celsius: ")
	fmt.Scanln(&celsius)

	reamur := celsius * 4 / 5
	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273
	
	fmt.Printf("Derajat Reamur: %d\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %d\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %d\n", kelvin)
}
```

**Screenshot Output :**
![image](https://github.com/user-attachments/assets/2e0d8f48-2e7f-407c-a6e5-e365e1a80504)

### 5. Program "ASCII"

**Source Code :**

```GO
package main

import "fmt"

func main() {
	var num int

	fmt.Print("Masukkan 5 angka antara (32-127) dipisahkan spasi: ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&num)
		fmt.Printf("%c", num)
	}
	fmt.Println()

	var input string
	fmt.Print("Masukkan 3 karakter: ")
	fmt.Scan(&input)

	if len(input) <= 3 {
		fmt.Print("Hasil konversi: ")
		for i := 0; i < len(input); i++ {
			fmt.Printf("%c", input[i]+1)
		}
		fmt.Println()
	} else {
		fmt.Println("Tidak boleh lebih dari 3 karakter.")
	}
}
```

**Screenshot Output :**
![image](https://github.com/user-attachments/assets/53a0f2a4-80af-4a6f-ac5a-2c0360418c9d)
