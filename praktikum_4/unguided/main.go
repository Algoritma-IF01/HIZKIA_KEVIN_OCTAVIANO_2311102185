package main

import "fmt"

func menu() {
	var pilih int

	fmt.Println("Menu: ")
	fmt.Println("1. Soal no.2 modul fungsi")
	fmt.Println("2. Soal no.3 modul fungsi")
	fmt.Println("3. Soal no.2 modul prosedur")
	fmt.Println("4. Soal no.3 modul prosedur")
	fmt.Print("Pilih[1-4]: ")
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		soal_2_fungsi()
	case 2:
		soal_3_fungsi()
	case 3:
		soal_2_prosedur()
	case 4:
		soal_3_prosedur()
	default:
		fmt.Println("Pilihan tidak tersedia")
	}
}

func main() {
	menu()
}