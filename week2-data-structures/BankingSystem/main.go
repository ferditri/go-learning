package main

import "fmt"

type Customer struct {
	ID    int
	Name  string
	Email string
	Phone string
}

type Account struct {
	Number      string
	CustomerID  int
	Balance     float64
	AccountType string // "savings", "checking", "business"
	IsActive    bool
}

// TODO: Buat methods untuk Account:
// 1. func (a *Account) Deposit(amount float64) error
// 2. func (a *Account) Withdraw(amount float64) error
// 3. func (a *Account) Transfer(to *Account, amount float64) error
// 4. func (a Account) GetAccountInfo() string
// 5. func (a Account) CalculateInterest(rate float64) float64

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	} else {
		a.Balance += amount
		return nil
	}
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdraw amount must be positive")
	} else if amount > a.Balance {
		return fmt.Errorf("insufficient balance")
	} else {
		a.Balance -= amount
		return nil
	}
}

func (a *Account) Transfer(to *Account, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("transfer amount must be positive")
	} else if amount > a.Balance {
		return fmt.Errorf("insufficient balance")
	} else {
		a.Balance -= amount
		to.Balance += amount
		return nil
	}
}

func (a Account) GetAccountInfo() string {
	Status := "Active"
	if !a.IsActive {
		Status = "Inactive"
	}
	return fmt.Sprintf("Account Number: %s\nCustomer ID: %d\nBalance: %.2f\nAccount Type: %s\nStatus: %s",
		a.Number, a.CustomerID, a.Balance, a.AccountType, Status)
}

func (a Account) CalculateInterest(rate float64) float64 {
	return a.Balance * rate / 100
}

type Bank struct {
	Name      string
	Customers []Customer
	Accounts  []Account
}

// TODO: Buat methods untuk Bank:
// 1. func (b *Bank) AddCustomer(customer Customer) error
// 2. func (b *Bank) OpenAccount(customerID int, accType string) *Account
// 3. func (b *Bank) CloseAccount(accountNumber string) error
// 4. func (b Bank) FindCustomerByID(id int) *Customer
// 5. func (b Bank) FindAccountsByCustomerID(customerID int) []Account
// 6. func (b Bank) GetTotalBankBalance() float64
// 7. func (b Bank) GetWealthyCustomers(threshold float64) []Customer

func (b *Bank) AddCustomer(customer Customer) error {
	for _, c := range b.Customers {
		if c.ID == customer.ID {
			return fmt.Errorf("customer with ID %d already exists", customer.ID)
		}
	}
	b.Customers = append(b.Customers, customer)
	return nil
}

func (b *Bank) OpenAccount(customerID int, accType string) *Account {
	account := Account{
		Number:      fmt.Sprintf("ACC%03d", len(b.Accounts)+1),
		CustomerID:  customerID,
		Balance:     0,
		AccountType: accType,
		IsActive:    true,
	}
	b.Accounts = append(b.Accounts, account)
	return &b.Accounts[len(b.Accounts)-1]
}

func (b *Bank) CloseAccount(accountNumber string) error {
	for i, a := range b.Accounts {
		if a.Number == accountNumber {
			b.Accounts = append(b.Accounts[:i], b.Accounts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("account with number %s not found", accountNumber)
}

func (b Bank) FindCustomerByID(id int) *Customer {
	for i := range b.Customers {
		if b.Customers[i].ID == id {
			return &b.Customers[i]
		}
	}
	return nil
}

func (b Bank) FindAccountsByCustomerID(customerID int) []Account {
	var result []Account
	for _, a := range b.Accounts {
		if a.CustomerID == customerID {
			result = append(result, a)
		}
	}
	return result
}

func (b Bank) GetTotalBankBalance() float64 {
	var total float64
	for _, a := range b.Accounts {
		total += a.Balance
	}
	return total
}

func (b Bank) GetWealthyCustomers(threshold float64) []Customer {
	var result []Customer
	for _, c := range b.Customers {
		accounts := b.FindAccountsByCustomerID(c.ID)
		var totalBalance float64
		for _, a := range accounts {
			totalBalance += a.Balance
		}
		if totalBalance >= threshold {
			result = append(result, c)
		}
	}
	return result
}

func main() {
	// TODO: Test banking system dengan transfer, deposit, withdraw
	bank := Bank{
		Name: "Bank GoLang",
	}
	fmt.Println("Welcome to", bank.Name)

	customer1 := Customer{
		ID:    1,
		Name:  "Alice",
		Email: "",
		Phone: "123-456-7890",
	}
	bank.AddCustomer(customer1)
	account1 := bank.OpenAccount(customer1.ID, "savings")
	account1.Deposit(1000)
	fmt.Println(account1.GetAccountInfo())

	customer2 := Customer{
		ID:    2,
		Name:  "Bob",
		Email: "",
		Phone: "987-654-3210",
	}
	bank.AddCustomer(customer2)
	account2 := bank.OpenAccount(customer2.ID, "checking")
	account2.Deposit(500)
	fmt.Println(account2.GetAccountInfo())
	account1.Transfer(account2, 200)
	fmt.Println("After transfer:")
	fmt.Println(account1.GetAccountInfo())
	fmt.Println(account2.GetAccountInfo())

}
