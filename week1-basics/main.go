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

	// latihan 4

	var (
		UTS float64 = 40
		UAS float64 = 60
	)

	var nilaiUTS, nilaiUAS float64 = 80, 76

	hasil := (nilaiUTS * (UTS / 100)) + (nilaiUAS * (UAS / 100))

	var grade rune
	if hasil >= 85 {
		grade = 'A'
	} else if hasil >= 70 {
		grade = 'B'
	} else if hasil >= 60 {
		grade = 'C'
	} else if hasil >= 50 {
		grade = 'D'
	} else {
		grade = 'E'
	}

	if grade == 'A' || grade == 'B' {
		fmt.Println("LULUS")
	} else {
		fmt.Println("Tidak Lulus")
	}

	// latihan 5

	var barang, jumlah, total, diskon, hasilAkhir float64

	barang = 5000
	jumlah = 40
	total = barang * jumlah

	if total >= 500000 {
		diskon = total * 20 / 100
	} else if total >= 200000 {
		diskon = total * 10 / 100
	} else if total >= 100000 {
		diskon = total * 5 / 100
	}

	hasilAkhir = total - diskon

	fmt.Println("Diskon: ", diskon)
	fmt.Println("Total: ", total)
	fmt.Println("Total Akhir: ", hasilAkhir)

	// latihan 6

	for x := 1; x <= 10; x++ {
		for y := 1; y <= 10; y++ {
			fmt.Println(x, " x ", y, " = ", x*y)
		}
	}

}
