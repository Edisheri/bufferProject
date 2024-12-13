package synchronized

import (
	"fmt"
	"sync/atomic"
)

func AtomicSync(messages []string, numWriters, numReaders int) {
	var buffer atomic.Value
	done := make(chan struct{})

	// Запуск писателей
	for i := 0; i < numWriters; i++ {
		go func(id int) {
			for _, msg := range messages {
				buffer.Store(msg) // Атомарная запись
				fmt.Printf("Писатель %d записал: %s\n", id, msg)
			}
			done <- struct{}{}
		}(i)
	}

	// Запуск читателей
	for i := 0; i < numReaders; i++ {
		go func(id int) {
			for range messages {
				msg := buffer.Load() // Атомарное чтение
				if msg != nil {
					fmt.Printf("Читатель %d прочитал: %s\n", id, msg)
				}
			}
			done <- struct{}{}
		}(i)
	}

	// Ожидаем завершения
	for i := 0; i < numWriters+numReaders; i++ {
		<-done
	}
}
