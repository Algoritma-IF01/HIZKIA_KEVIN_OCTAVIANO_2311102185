package main

import "fmt"

func menu() {
	var pilih int

	fmt.Println("Menu: ")
	fmt.Println("1. Soal no.2")
	fmt.Println("2. Soal no.3")
	fmt.Println("3. Soal no.4")
	fmt.Print("Pilih[1-3]: ")
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		soal_2()
	case 2:
		soal_3()
	case 3:
		soal_4()
	default:
		fmt.Println("Pilihan tidak tersedia")
	}
}

func main() {
	menu()
}