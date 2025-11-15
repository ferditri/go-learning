package main

import "fmt"

func main() {
	itemBelanja := []string{}
	var item string
	var pilihan int
	var index int

	for {
		fmt.Println("MENU SHOOPING LIST")
		fmt.Println("1. Tambah Item")
		fmt.Println("2. Lihat Semua Item")
		fmt.Println("3. Hapus Item")
		fmt.Println("4. Keluar")
		fmt.Print("Masukkan pilihan anda: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Print("Masukkan nama item yang ingin ditambahkan: ")
			fmt.Scan(&item)
			itemBelanja = append(itemBelanja, item)
			fmt.Println("Item telah ditambahkan")
		} else if pilihan == 2 {
			fmt.Println("List ITEM")
			for x := 0; x < len(itemBelanja); x++ {
				fmt.Printf("%d. %s\n", x+1, itemBelanja[x])
			}
		} else if pilihan == 3 {
			fmt.Println("List ITEM")
			for x := 0; x < len(itemBelanja); x++ {
				fmt.Printf("%d. %s\n", x+1, itemBelanja[x])
			}

			fmt.Print("Masukkan nomor yang ingin dihapus: ")
			fmt.Scan(&index)
			// misal index (2)
			// itemBelanja := []string{"Baju", "celana"}
			itemBelanja = append(itemBelanja[:index-1], itemBelanja[index:]...)
			fmt.Println("Hapus sukses")
			fmt.Println("List ITEM")
			for x := 0; x < len(itemBelanja); x++ {
				fmt.Printf("%d. %s\n", x+1, itemBelanja[x])
			}

		} else if pilihan == 4 {
			break
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}

}
