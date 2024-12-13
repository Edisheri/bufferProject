package synchronized

import (
	"fmt"
	"sync"
)

func ChannelSync(messages []string, numWriters, numReaders int) {
	buffer := make(chan string, 1) // Канал для буфера
	done := make(chan struct{})    // Канал завершения
	var wg sync.WaitGroup          // Для ожидания завершения всех горутин

	wg.Add(numWriters + numReaders)

	// Горутины для писателей
	for i := 0; i < numWriters; i++ {
		go func(id int) {
			defer wg.Done()
			for _, msg := range messages {
				buffer <- msg
				fmt.Printf("Писатель %d записал: %s\n", id, msg)
			}
		}(i)
	}

	// Горутины для читателей
	for i := 0; i < numReaders; i++ {
		go func(id int) {
			defer wg.Done()
			for range messages {
				msg, ok := <-buffer
				if ok {
					fmt.Printf("Читатель %d прочитал: %s\n", id, msg)
				}
			}
		}(i)
	}

	// Ждем завершения всех писателей и читателей
	go func() {
		wg.Wait()
		close(buffer) // Закрываем канал только после завершения всех операций
		close(done)
	}()

	<-done // Ожидание завершения
}
