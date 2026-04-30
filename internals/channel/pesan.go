package channel

import (
	"fmt"
	"sync"
)

func Pesan(message string, msgChn chan string, wg *sync.WaitGroup) {
	wg.Done()
	msgChn <- message
	fmt.Printf("%s, dikirim\n", message)

}
