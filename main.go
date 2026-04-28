package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	filereader "github.com/nopalllfd/koda-b7-go/internals/FileReader"
	"github.com/nopalllfd/koda-b7-go/internals/calc"
	"github.com/nopalllfd/koda-b7-go/internals/insert"
	"github.com/nopalllfd/koda-b7-go/internals/model"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=== MENU ===")
		fmt.Println("1. Luas Lingkaran")
		fmt.Println("2. Keliling Lingkaran")
		fmt.Println("3. Luas dan Keliling Lingkaran")
		fmt.Println("4. Segitiga Siku-Siku")
		fmt.Println("5. Sisip angka")
		fmt.Println("6. User")
		fmt.Println("7. Read file")
		fmt.Println("8. Person method")
		fmt.Println("0. Exit")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "0":
			fmt.Println("Keluar...")
			return
		case "1":
			fmt.Print("Masukkan jari-jari lingkaran : ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			radius64, _ := strconv.ParseFloat(input, 32)
			radius := float32(radius64)
			result := calc.Luas(radius)
			fmt.Printf("\n Luas lingkaran adalah = %.2f\n\n\n", result)
		case "2":
			fmt.Print("Masukkan jari-jari lingkaran : ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			radius64, _ := strconv.ParseFloat(input, 32)
			radius := float32(radius64)
			result := calc.Keliling(radius)
			fmt.Printf("\n Keliling lingkaran adalah = %.2f\n\n\n", result)
		case "3":
			fmt.Print("Masukkan jari-jari lingkaran : ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			radius64, _ := strconv.ParseFloat(input, 32)
			radius := float32(radius64)
			l, k := calc.LuasDanKeliling(radius)
			fmt.Printf("\n Luas lingkaran adalah = %.2f\n Keliling lingkaran adalah = %.2f\n\n", l, k)
		case "4":
			fmt.Print("Masukkan tinggi segitiga : ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			tinggi, _ := strconv.Atoi(input)
			calc.SegitigaSikuSiku(tinggi)
		case "5":

			fmt.Print("Masukkan angka : ")
			secIn, _ := reader.ReadString('\n')
			secIn = strings.TrimSpace(secIn)

			fmt.Print("Masukkan urutan ke- : ")
			firstIn, _ := reader.ReadString('\n')
			firstIn = strings.TrimSpace(firstIn)

			index, _ := strconv.Atoi(firstIn)
			angka, _ := strconv.Atoi(secIn)

			index = index - 1

			result := insert.SisipAngka(angka, index)
			fmt.Println(result)
		case "6":
			data := model.TampilinData()
			fmt.Println(data)
		case "7":
			fmt.Print("Masukkan path file : ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			filereader.ReadFile(input)
		case "8":
			fmt.Print("Masukkan nama : ")
			nameInput, _ := reader.ReadString('\n')
			nameInput = strings.TrimSpace(nameInput)
			fmt.Print("Masukkan alamat : ")
			addressInput, _ := reader.ReadString('\n')
			addressInput = strings.TrimSpace(addressInput)

			fmt.Print("Masukkan telepon : ")
			phoneInput, _ := reader.ReadString('\n')
			phoneInput = strings.TrimSpace(phoneInput)

			user := model.NewPerson(nameInput, addressInput, phoneInput)
			user.GreetPerson()
			user.PrintPerson()
		}
	}

	// a := 10
	// var b *int = &a
	// fmt.Println("Nilai a", a)
	// *b = 20
	// fmt.Println("Alamat a", &a)
	// fmt.Println("nilai b", *b)
	// filereader.ReadFile()
	// nopal := model.NewPerson("Nopal", "Bogor", "8128321")
	// fmt.Println(*nopal)
	// nopal.GreetPerson()
	// nopal.SetName("Padil")
	// nopal.GreetPerson()
	// nopal.PrintPerson()
}
