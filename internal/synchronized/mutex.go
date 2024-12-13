package synchronized

import (
	"fmt"
	"sync"
)

func MutexSync(messages []string, numWriters, numReaders int) {
	var buffer string
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(numWriters + numReaders)

	for i := 0; i < numWriters; i++ {
		go func(id int) {
			defer wg.Done()
			for _, msg := range messages {
				lock.Lock()
				buffer = msg
				fmt.Printf("Писатель %d записал: %s\n", id, msg)
				lock.Unlock()
			}
		}(i)
	}

	for i := 0; i < numReaders; i++ {
		go func(id int) {
			defer wg.Done()
			for range messages {
				lock.Lock()
				if buffer != "" {
					fmt.Printf("Читатель %d прочитал: %s\n", id, buffer)
					buffer = "" // Очистка буфера
				}
				lock.Unlock()
			}
		}(i)
	}

	wg.Wait()
}
