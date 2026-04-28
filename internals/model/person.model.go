package model

import "fmt"

type Person struct {
	Name    string
	Address string
	Phone   string
}

func NewPerson(name, address, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) PrintPerson() {
	fmt.Printf("Nama : %s, Alamat : %s, Telepon : %s\n", p.Name, p.Address, p.Phone)
}
func (p *Person) GreetPerson() {
	fmt.Printf("Hallo, %s\n", p.Name)
}
