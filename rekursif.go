package main

import (
	"fmt"
	"unicode"
)

var jumlah = 0.0
var jumlahI = 0.0

func deretReal(i float64, n float64) float64 {

	if i <= 1 {
		fmt.Printf("%.2f + ", n)
		jumlah = jumlah + n
		return n
	}
	y := deretReal(i-1, n)
	jumlah = jumlah + n/i
	fmt.Printf("%.2f + ", n/i)

	return y
}

func deretPecahan(i float64) float64 {
	if i <= 1 {
		jumlahI = jumlahI + 1
		jumlah = jumlah + 2
		fmt.Printf("1/2x + ")
		return i
	}
	y := deretPecahan(i - 1)
	jumlahI = jumlahI + i
	x := i * i
	jumlah = jumlah + x
	fmt.Printf("%.0f/%.0fx + ", i, x)
	return y
}

func deretE(i float64) float64 {
	if i <= 1 {

		jumlahI = jumlahI + 1
		jumlah = jumlah + 2
		fmt.Printf("2x/1 + ")
		return i
	}
	y := deretE(i - 1)
	jumlahI = jumlahI + i
	x := i * i
	jumlah = x + jumlah
	fmt.Printf("%.0fx/%.0f + ", x, i)
	return y
}

func nonKapital(i int, x string) int {
	if i < 0 {
		return i
	}
	y := rune(x[i])
	if unicode.IsLower(y) {
		jumlah = jumlah + 1
	}

	return nonKapital(i-1, x)

}

func cariAngka(i int, x string) int {
	if i < 0 {
		return i
	}
	y := rune(x[i])
	if unicode.IsNumber(y) {
		jumlah = jumlah + 1
	}

	return cariAngka(i-1, x)

}

func soal1() {
	var a, input float64
	fmt.Print("Masukkan panjang deret : ")
	fmt.Scan(&input)
	fmt.Print("Masukkan angka awal : ")
	fmt.Scan(&a)

	fmt.Print("SUM = ")
	deretReal(input, a)
	fmt.Printf("= %.2f", jumlah)

}

var pilihan int
var kata string
var batas float64

func menu() {

	fmt.Println("================== PILIHAN ===================")
	fmt.Println("1. Deret Pembagian Bilangan Real (a, b, c)")
	fmt.Println("2. Deret Pembagian Bilangan Pecahan")
	fmt.Println("3. Deret Pembagian Angka")
	fmt.Println("4. Hitung Jumlah Huruf Non Kapital")
	fmt.Println("5. Hitung Jumlah Angka Dalam Suatu Kata")
	fmt.Println("==============================================")
	fmt.Print("Masukkan Pilihan : ")
	fmt.Scan(&pilihan)
	fmt.Println("==============================================")

	if pilihan == 1 {
		soal1()
	} else if pilihan == 2 {
		fmt.Print("Masukkan Panjang Deret : ")
		fmt.Scan(&batas)

		fmt.Print("SUM = ")
		deretPecahan(batas)
		fmt.Printf("= %.0f/%.0fx", jumlahI, jumlah)
	} else if pilihan == 3 {
		fmt.Print("Masukkan Panjang Deret : ")
		fmt.Scan(&batas)

		fmt.Print("SUM = ")
		deretE(batas)
		fmt.Printf("= %.0fx/%.0f", jumlah, jumlahI)
	} else if pilihan == 4 {
		fmt.Print("Masukkan sebuah kata : ")
		fmt.Scan(&kata)

		nonKapital(len(kata)-1, kata)
		fmt.Print("Jumlah huruf non kapital dalam kata : ", jumlah)
	} else if pilihan == 5 {
		fmt.Print("Masukkan sebuah kata : ")
		fmt.Scan(&kata)

		cariAngka(len(kata)-1, kata)
		fmt.Print("Jumlah angka dalam kata : ", jumlah)
	}
}

func main() {
	menu()
}
