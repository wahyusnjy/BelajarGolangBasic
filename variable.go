package main

import "fmt"

func variable() {

	var firstname string = "Wahyu"

	var lastname string
	lastname = "Sanjaya"

	//fmt.Printf("Halo!\n", firstname, lastname)
	/*Hasil print diatas ketika tidak memakai %s adalah
	= Halo!%!(EXTRA string=Wahyu, string=Sanjaya)
	*/

	//fmt.Printf("Halo %s %s !\n", firstname, lastname)
	// karakter %s disitu akan diganti dengan data string
	/*Hasil print diatas ketika memakai karakter %s adalah
	= Halo Wahyu Sanjaya !
	*/

	fmt.Println("halo", firstname, lastname+"!")
	// tanda + ("halo", firstname, lastname '+' "!") jika ditempatkan diantara string, fungsinya adalah untuk penggabungan string atau stringconcatenation
	/*
		Fungsi fmt.Printf() tidak menghasilkan baris baru di akhir text,
		oleh karena itu digunakanlah literal newline yaitu \n, untuk memunculkan baris baru di akhir.
		Hal ini sangat berbeda jika dibandingkan dengan fungsi fmt.Println() yang secara otomatis menghasilkan new line (baris baru) di akhir.
	*/
}

// Uncomment func main dibawah ini dengan menghapus /* */ untuk menjalankan programnya go run variable.go

func main() {
	variable()
}
