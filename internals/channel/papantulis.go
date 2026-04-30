package channel

import (
	"fmt"
	"sync"
	"time"
)

// msgChn chan string
func PapanTulis(msgChn chan string, wg *sync.WaitGroup) {
	wg.Done()
	for v := range msgChn {
		fmt.Printf("Message receive: %s\n", v)
		time.Sleep(1 * time.Second)
	}

}
