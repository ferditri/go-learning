package main

import "fmt"

type Contact struct {
	Name        string
	PhoneNumber int
}

func (c Contact) Display() string {
	return fmt.Sprintf("Name: %s\nPhone: %d\n", c.Name, c.PhoneNumber)
}

func (c *Contact) UpdatePhone(value int) {
	c.PhoneNumber = value
	fmt.Printf("Success change %s phone number to %d\n", c.Name, value)
}

type ContactList struct {
	Contacts []Contact
}

func (cl *ContactList) Add(contact Contact) {
	cl.Contacts = append(cl.Contacts, contact)
	fmt.Println("Contact added successfully")
}

func (cl *ContactList) FindByName(name string) *Contact {
	for i := range cl.Contacts {
		if cl.Contacts[i].Name == name {
			return &cl.Contacts[i]
		}
	}
	return nil
}

func (cl *ContactList) Count() int {
	return len(cl.Contacts)
}

func main() {

	contactList := ContactList{}

	c1 := Contact{Name: "Ferdi", PhoneNumber: 62874646474}
	c2 := Contact{Name: "King", PhoneNumber: 62863122535}

	contactList.Add(c1)
	contactList.Add(c2)

	fmt.Println(c1.Display())
	fmt.Println(c2.Display())
}
