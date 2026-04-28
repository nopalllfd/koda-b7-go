package checkout

import "fmt"

// type DataPembayaran struct {

// }

type Payment interface {
	TransferSuccess(prices []int)
	TransferFiktif(prices []int) string
}

type MethodBank struct {
	Amount []int
}

type MethodOnline struct {
	Amount []int
}

func (b *MethodBank) TransferSuccess(prices []int) {
	total := 0
	for _, v := range prices {
		total += v
	}
	fmt.Printf("----Pembayaran sukses----\n")
	fmt.Printf("Total Pembayaran : %d\n", total)
	fmt.Println("Metode yang digunakan adalah Bank")
	fmt.Println("-------------------------------")

}
func (o *MethodOnline) TransferSuccess(prices []int) {
	total := 0
	for _, v := range prices {
		total += v
	}
	fmt.Printf("----Pembayaran sukses----\n")
	fmt.Printf("Total Pembayaran : %d\n", total)
	fmt.Println("Metode yang digunakan adalah Online")
	fmt.Println("-------------------------------")

}

func (b *MethodBank) TransferFiktif(prices []int) string {
	total := 0
	for _, v := range prices {
		if v <= 0 {
			return "Error"
		}
		total += v
	}
	b.Amount = append(b.Amount, total)
	return "success"
}
func (o *MethodOnline) TransferFiktif(prices []int) string {
	total := 0
	for _, v := range prices {
		if v <= 0 {
			return "Error"
		}
		total += v
	}
	o.Amount = append(o.Amount, total)
	return "success"
}
