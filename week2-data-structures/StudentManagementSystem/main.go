package main

import "fmt"

type Mahasiswa struct {
	NIM      string
	Nama     string
	Jurusan  string
	Semester int
	IPK      float64
}

func (m Mahasiswa) Info() string {
	return fmt.Sprintf("%s - %s (%s) Semester %d, IPK: %.2f",
		m.NIM, m.Nama, m.Jurusan, m.Semester, m.IPK)
}

func (m *Mahasiswa) UpdateIPK(ipkBaru float64) {
	m.IPK = ipkBaru
	fmt.Println("IPK berhasil diupdate!")
}

type Kampus struct {
	Nama            string
	DaftarMahasiswa []Mahasiswa
}

func (k *Kampus) TambahMahasiswa(mhs Mahasiswa) {
	k.DaftarMahasiswa = append(k.DaftarMahasiswa, mhs)
	fmt.Printf("Mahasiswa %s berhasil ditambahkan!\n", mhs.Nama)
}

func (k Kampus) CariByNIM(nim string) *Mahasiswa {
	for i := range k.DaftarMahasiswa {
		if k.DaftarMahasiswa[i].NIM == nim {
			return &k.DaftarMahasiswa[i]
		}
	}
	return nil
}

func (k Kampus) MahasiswaBerprestasi(minIPK float64) []Mahasiswa {
	var hasil []Mahasiswa
	for _, mhs := range k.DaftarMahasiswa {
		if mhs.IPK >= minIPK {
			hasil = append(hasil, mhs)
		}
	}
	return hasil
}

func (k Kampus) TampilkanSemua() {
	fmt.Printf("\n=== Daftar Mahasiswa %s ===\n", k.Nama)
	for i, mhs := range k.DaftarMahasiswa {
		fmt.Printf("%d. %s\n", i+1, mhs.Info())
	}
}

func main() {
	kampus := Kampus{Nama: "Universitas Indonesia"}

	kampus.TambahMahasiswa(Mahasiswa{"001", "Budi", "Informatika", 3, 3.75})
	kampus.TambahMahasiswa(Mahasiswa{"002", "Ani", "SI", 2, 3.85})

	kampus.TampilkanSemua()
}
