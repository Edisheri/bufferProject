package unsynced

import (
	"fmt"
	"sync"
)

var buffer string // Глобальная переменная для буфера

func DemonstrateUnsynced(messages []string, numWriters, numReaders int) {
	var wg sync.WaitGroup
	wg.Add(numWriters + numReaders)

	for i := 0; i < numWriters; i++ {
		go func(id int) {
			defer wg.Done()
			for _, msg := range messages {
				buffer = msg
				fmt.Printf("Писатель %d записал: %s\n", id, msg)
			}
		}(i)
	}

	for i := 0; i < numReaders; i++ {
		go func(id int) {
			defer wg.Done()
			for range messages {
				if buffer != "" {
					fmt.Printf("Читатель %d прочитал: %s\n", id, buffer)
					buffer = "" // Очистка буфера
				}
			}
		}(i)
	}

	wg.Wait()
}
