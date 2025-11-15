package main

import "fmt"

func main() {

	// latihan 1

	namaLengkap := "Ferdiyanto Tri Setiawan"
	NPM := 25082010135
	jurusan := "Sistem Informasi"
	semester := 1
	sudahBekerja := false
	alamat := "Sidoarjo"

	fmt.Println("Nama Lengkap saya:\t", namaLengkap)
	fmt.Println("NPM saya:\t", NPM)
	fmt.Println("Jurusan saya:\t", jurusan)
	fmt.Println("Saya semester:\t", semester)
	fmt.Println("Sudah Bekerja:\t", sudahBekerja)
	fmt.Println("Alamat:\t", alamat)

	// latihan 2

	angka1 := 5
	angka2 := 5

	perkalian := angka1 * angka2
	pembagian := angka1 / angka2
	penambahan := angka1 + angka2
	pengurangan := angka1 - angka2

	fmt.Println("Perkalian: ", perkalian)
	fmt.Println("Pembagian: ", pembagian)
	fmt.Println("Penambahan: ", penambahan)
	fmt.Println("Pengurangan: ", pengurangan)

	// latihan 3

	var celcius, fahrenheit, kelvin float64

	celcius = 18

	var C = celcius
	fahrenheit = (C * 9 / 5) + 32
	kelvin = C + 273.15

	fmt.Println("Fahrenheit: ", fahrenheit)
	fmt.Println("Kelvin: ", kelvin)

}
