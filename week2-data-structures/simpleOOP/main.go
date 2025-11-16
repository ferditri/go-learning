package main

import "fmt"

type BankAccount struct {
	Nama  string
	Saldo float64
}

func (b *BankAccount) deposit(saldo float64) {
	if saldo > 0 {
		b.Saldo += saldo
		fmt.Printf("Sukses menambahkan saldo sebesar %.0f\n", saldo)
	} else {
		fmt.Println("Input jumlah saldo yang benar")
	}
}

func (b *BankAccount) withdraw(jumlah float64) bool {
	if b.Saldo > jumlah {
		b.Saldo -= jumlah
		fmt.Printf("Berhasil menarik saldo sebesar %.0f\n", jumlah)
		return true
	} else {
		fmt.Println("Saldo anda kurang")
		return false
	}
}

func (b *BankAccount) cekSaldo() {
	fmt.Println("Saldo anda adalah ", b.Saldo)
}
func (b *BankAccount) info() {
	fmt.Printf("Nama %s mempunyai saldo sebesar %.0f\n ", b.Nama, b.Saldo)
}

func main() {

	bank := BankAccount{
		Nama:  "Ferdi",
		Saldo: 200000,
	}

	bank.cekSaldo()

	bank.deposit(50000)
	bank.cekSaldo()
	bank.withdraw(100000)
	bank.cekSaldo()

	bank.info()

}
