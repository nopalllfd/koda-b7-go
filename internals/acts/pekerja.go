package acts

import (
	"fmt"
	"time"
)

func Mandi() {
	fmt.Println("Memulai Mandi")
	time.Sleep(2 * time.Second)
	fmt.Println("Mandi Selesai")
}
func BuatKopi() {
	fmt.Println("Memulai Membuat Kopi")
	time.Sleep(1 * time.Second)
	fmt.Println("Membuat Kopi Selesai")
}
func MenyiapkanSarapan() {
	fmt.Println("Memulai Menyiapkan Sarapan")
	time.Sleep(2 * time.Second)
	fmt.Println("Menyiapkan Sarapan Selesai")
}
func Merapikan() {
	fmt.Println("Memulai Merapikan")
	time.Sleep(1 * time.Second)
	fmt.Println("Merapikan Selesai")
}
