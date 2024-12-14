package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	var temp, hitung int
	temp = bilangan
	for temp > 0 {
		hitung++
		temp /= 10
	}

	tengah := (hitung + 1) / 2

	pembagi := 1
	for i := 0; i < tengah; i++ {
		pembagi *= 10
	}

	bagianKiri := bilangan / pembagi
	bagianKanan := bilangan % pembagi

	fmt.Println(bagianKiri, bagianKanan)
	fmt.Println(bagianKiri + bagianKanan)
}
