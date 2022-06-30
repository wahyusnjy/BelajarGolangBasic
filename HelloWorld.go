package main

import "fmt"

func helloWorld() {
	fmt.Println("Hello World")
	/*
	Fungsi fmt.Println() berada dalam package fmt, 
	maka untuk menggunakannya perlu package tersebut untuk di-import terlebih dahulu.

	Fungsi fmt.Println() dapat menampung parameter yang tidak terbatas jumlahnya. 
	Semua data parameter akan dimunculkan dengan pemisah tanda spasi
	*/
}
// Uncomment func main dibawah ini dengan menghapus /* */ untuk menjalankan programnya go run variable.go
// func main (){
// 	helloWorld()
// }
