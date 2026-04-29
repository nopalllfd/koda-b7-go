package acts

import (
	"fmt"
	"sync"
	"time"
)

func Mandi(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Memulai Mandi")
	time.Sleep(2 * time.Second)
	fmt.Println("Mandi Selesai")
}
func BuatKopi(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Memulai Membuat Kopi")
	time.Sleep(1 * time.Second)
	fmt.Println("Membuat Kopi Selesai")
}
func MenyiapkanSarapan(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Memulai Menyiapkan Sarapan")
	time.Sleep(2 * time.Second)
	fmt.Println("Menyiapkan Sarapan Selesai")
}
func Merapikan(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Memulai Merapikan")
	time.Sleep(1 * time.Second)
	fmt.Println("Merapikan Selesai")
}
