package main

import "fmt"

type Buku struct {
	ID          int
	Judul       string
	Pengarang   string
	TahunTerbit int
	Tersedia    bool
}

func (b *Buku) Pinjam() {
	if b.Tersedia {
		b.Tersedia = false
		fmt.Printf("Buku '%s' berhasil dipinjam\n", b.Judul)
	} else {
		fmt.Printf("Buku '%s' sedang tidak tersedia\n", b.Judul)
	}
}

func (b *Buku) Kembalikan() {
	if !b.Tersedia {
		b.Tersedia = true
		fmt.Printf("Buku '%s' berhasil dikembalikan\n", b.Judul)
	} else {
		fmt.Printf("Buku %s sudah tersedia\n", b.Judul)
	}
}

func (b *Buku) Info() string {
	status := "Tersedia"
	if !b.Tersedia {
		status = "Dipinjam"
	}
	return fmt.Sprintf("ID: %d\nJudul: %s\nPengarang: %s\nTahun: %d\nStatus: %s",
		b.ID, b.Judul, b.Pengarang, b.TahunTerbit, status)
}

type Perpustakaan struct {
	Nama       string
	DaftarBuku []Buku
}

func (p *Perpustakaan) TambahBuku(buku Buku) {
	p.DaftarBuku = append(p.DaftarBuku, buku)
	fmt.Printf("Berhasil menambahkan buku %s ke perpustakaan\n", buku.Judul)
}

func (p *Perpustakaan) CariBuku(judul string) *Buku {
	for i := range p.DaftarBuku {
		if p.DaftarBuku[i].Judul == judul {
			return &p.DaftarBuku[i]
		}
	}
	return nil
}

func (p *Perpustakaan) TampilkanSemuaBuku() {
	if len(p.DaftarBuku) == 0 {
		fmt.Println("Tidak ada buku")
		return
	}

	for i, v := range p.DaftarBuku {
		status := "Tersedia"
		if !v.Tersedia {
			status = "Tidak Tersedia"
		}
		fmt.Printf("%d. %s - %s [%s]\n", i+1, v.Judul, v.Pengarang, status)
	}
}

func main() {

	perpus := Perpustakaan{
		Nama: "Raja Perpustakaan",
	}

	buku1 := Buku{ID: 1, Judul: "Laskar Pelangi", Pengarang: "Andrea Hirata", TahunTerbit: 2005, Tersedia: true}
	buku2 := Buku{ID: 2, Judul: "Bumi Manusia", Pengarang: "Pramoedya Ananta Toer", TahunTerbit: 1980, Tersedia: true}
	buku3 := Buku{ID: 3, Judul: "Dilan 1990", Pengarang: "Pidi Baiq", TahunTerbit: 2014, Tersedia: true}

	perpus.TambahBuku(buku1)
	perpus.TambahBuku(buku2)
	perpus.TambahBuku(buku3)

	perpus.TampilkanSemuaBuku()

	perpus.CariBuku("Dilan 1990")
}
