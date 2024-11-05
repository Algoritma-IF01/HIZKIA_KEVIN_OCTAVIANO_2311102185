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