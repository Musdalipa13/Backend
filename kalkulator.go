package main

import "fmt"

func main() {
	var a, b float32
	var operator int

	fmt.Print("Program Kalkulator Sederhana di Golang:\n 1. Operasi Perkalian\n 2. Operasi Pembagian\n 3. Operasi Penambahan\n 4. Operasi Pengurangan\n 5. Operasi Modulus\n 6. Exit\nPilih salah satu angka di atas: ")
	fmt.Scanln(&operator)

	if operator >= 6 {
		fmt.Println("Keluar dari program.")
		return
	}

	fmt.Print("Masukkan angka Pertama: ")
	fmt.Scanln(&a)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&b)

	if operator == 1 {
		fmt.Printf("Hasil: %.2f\n", a*b)
	} else if operator == 2 {
		if b != 0 {
			fmt.Printf("Hasil: %.2f\n", a/b)
		} else {
			fmt.Println("Error.")
		}
	} else if operator == 3 {
		fmt.Printf("Hasil: %.2f\n", a+b)
	} else if operator == 4 {
		fmt.Printf("Hasil: %.2f\n", a-b)
	} else if operator == 5 {
		intA := int(a)
		intB := int(b)
		if intB != 0 {
			fmt.Printf("Hasil: %d\n", intA%intB)
		} else {
			fmt.Println("Error.")
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}
