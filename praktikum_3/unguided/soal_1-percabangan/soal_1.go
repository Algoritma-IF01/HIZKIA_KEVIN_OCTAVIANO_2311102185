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
