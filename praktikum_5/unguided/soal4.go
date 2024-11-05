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