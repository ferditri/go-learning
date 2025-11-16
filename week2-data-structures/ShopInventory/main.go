package main

import (
	"fmt"
)

func main() {
	produk := make(map[string]int)

	var pilihan, stok int

	var namaProduk string

	for {
		fmt.Println("Shop Inventory")
		fmt.Println("1. Tambah Produk")
		fmt.Println("2. Update Stok")
		fmt.Println("3. Cek Stok")
		fmt.Println("4. Tampilkan Semua Produk")
		fmt.Println("5. Hapus Produk")
		fmt.Println("6. Cek (low stok)")
		fmt.Println("7. Keluar")
		fmt.Print("Masukkan Pilihan Menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("Tambah Produk")
			fmt.Print("Masukkan nama produk: ")
			fmt.Scan(&namaProduk)
			fmt.Print("Masukkan jumlah stok: ")
			fmt.Scan(&stok)

			if _, exists := produk[namaProduk]; exists {
				fmt.Println("Produk sudah ada")
			} else {
				produk[namaProduk] = stok
				fmt.Println("Sukses menambahkan barang:", namaProduk, "dengan stok", stok)
			}
		case 2:
			fmt.Println("Tambah Stok")
			fmt.Print("Masukkan nama produk: ")
			fmt.Scan(&namaProduk)

			if _, exists := produk[namaProduk]; exists {
				fmt.Print("Masukkan stok baru: ")
				fmt.Scan(&stok)
				produk[namaProduk] = stok
				fmt.Println("Stok berhasil diupdate:", namaProduk, "=", stok)
			} else {
				fmt.Println("Produk tidak ada")
			}
		case 3:
			fmt.Println("Cek stok")
			fmt.Print("Masukkan nama produk: ")
			fmt.Scan(&namaProduk)

			if stok, exists := produk[namaProduk]; exists {
				fmt.Println("Stok", namaProduk, ":", stok)
			} else {
				fmt.Println("Produk tidak ada")
			}
		case 4:
			fmt.Println("Menampilkan semua produk")
			if len(produk) == 0 {
				fmt.Println("Tidak ada produk dalam inventory.")
			} else {
				num := 1
				for key, value := range produk {
					fmt.Println(num, ". ", key, " = ", value)
					num++
				}
			}
		case 5:
			fmt.Println("Hapus produk")
			fmt.Print("Masukkan nama produk yang akan dihapus: ")
			fmt.Scan(&namaProduk)

			if _, exists := produk[namaProduk]; exists {
				delete(produk, namaProduk)
				fmt.Println("Produk", namaProduk, "berhasil dihapus.")
			} else {
				fmt.Println("Produk tidak ada")
			}
		case 6:
			fmt.Println("Cek stok")
			threshold := 5
			found := false

			fmt.Printf("Produk dengan stok kurang dari %d:\n", threshold)
			for key, value := range produk {
				if value < threshold {
					fmt.Println("-", key, ":", value)
					found = true
				}
			}

			if !found {
				fmt.Println("Tidak ada produk dengan low stok.")
			}
		case 7:
			fmt.Println("Terimakasih")
			return
		default:
			fmt.Println("Pilihan Tidak Valid")
		}
	}

}
