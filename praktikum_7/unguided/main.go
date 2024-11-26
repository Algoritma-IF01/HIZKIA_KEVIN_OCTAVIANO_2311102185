package main

import "fmt"

func menu() {
	var pilih int

	fmt.Println("Menu: ")
	fmt.Println("1. Soal_selectionsort no.1")
	fmt.Println("2. Soal_selectionsort no.2")
	fmt.Println("3. Soal_insertionsort no.1")
	fmt.Println("4. Soal_insertionsort no.2")
	fmt.Print("Pilih[1-3]: ")
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		soal1_selectionsort()
	case 2:
		soal2_selectionsort()
	case 3:
		soal1_insertionsort()
	case 4:
		soal2_insertionsort()
	default:
		fmt.Println("Pilihan tidak tersedia")
	}
}

func main() {
	menu()
}