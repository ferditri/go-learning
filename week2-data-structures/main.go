package main

import "fmt"

func main() {

	// day 1-2

	// array

	var arr [5]int
	arr[0] = 12
	arr[1] = 23
	arr[2] = 55
	arr[3] = 99
	arr[4] = 52

	fmt.Println("Array: ", arr)
	fmt.Println("Array pertama: ", arr[0])
	fmt.Println("Panjang array: ", len(arr))

	buah := [3]string{"Apple", "Banana", "Grapes"}

	for x := 0; x < len(buah); x++ {
		fmt.Printf("Index %d: %s\n", x, buah[x])
	}

	for index, value := range buah {
		fmt.Printf("%d. %s\n", index+1, value)
	}

	// slices

	angka := []int{12, 43, 77, 40, 100}
	fmt.Println("Slice: ", angka)

	angka = append(angka, 15, 66, 88)
	fmt.Println("Setelah Append: ", angka)

	subset := angka[2:5]
	fmt.Println("[2:5]: ", subset)

	var nama []string
	nama = append(nama, "Ferdi")
	nama = append(nama, "Tri")
	nama = append(nama, "Setiawan")
	fmt.Println("Nama: ", nama)

	// latihan 1

	var nilai []float64
	var inputNilai, rataRata, hasilRataRata float64
	var nilaiTerkecil, nilaiTerbesar float64

	for x := 0; x < 5; x++ {
		fmt.Print("Masukkan nilai ke- ", x+1, ": ")
		fmt.Scan(&inputNilai)
		nilai = append(nilai, inputNilai)
		rataRata += inputNilai
	}
	hasilRataRata = rataRata / float64(len(nilai))

	nilaiTerbesar = nilai[0]
	nilaiTerkecil = nilai[0]

	for x := 0; x < len(nilai); x++ {
		if nilai[x] > nilaiTerbesar {
			nilaiTerbesar = nilai[x]
		}
		if nilai[x] < nilaiTerkecil {
			nilaiTerkecil = nilai[x]
		}
	}

	fmt.Println("Nilai rata-rata: ", hasilRataRata)
	fmt.Println("Nilai Terbesar: ", nilaiTerbesar)
	fmt.Println("Nilai Terkecil: ", nilaiTerkecil)

	for x := 0; x < len(nilai); x++ {
		if nilai[x] > hasilRataRata {
			fmt.Println("Nilai yang diatas Rata-rata: ", nilai[x])
		}
	}

}
