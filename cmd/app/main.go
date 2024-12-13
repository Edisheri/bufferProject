package main

import (
	"bufferProject/internal/synchronized"
	"bufferProject/internal/unsynced"
	"bufferProject/performance"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Получаем параметры от пользователя
	messages, numWriters, numReaders := getUserInput()

	fmt.Println("\n--- Без синхронизации ---")
	unsynced.DemonstrateUnsynced(messages, numWriters, numReaders)

	fmt.Println("\n--- Синхронизация с использованием WaitGroup ---")
	performance.MeasurePerformance(synchronized.WaitGroupSync, messages, numWriters, numReaders)

	fmt.Println("\n--- Синхронизация с использованием Channel ---")
	performance.MeasurePerformance(synchronized.ChannelSync, messages, numWriters, numReaders)

	fmt.Println("\n--- Синхронизация с использованием Mutex ---")
	performance.MeasurePerformance(synchronized.MutexSync, messages, numWriters, numReaders)

}

// getUserInput собирает данные от пользователя
func getUserInput() ([]string, int, int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите количество сообщений:")
	numMessages := readInt(reader)

	fmt.Println("Введите размер каждого сообщения (в байтах):")
	messageSize := readInt(reader)

	fmt.Println("Введите количество потоков писателей:")
	numWriters := readInt(reader)

	fmt.Println("Введите количество потоков читателей:")
	numReaders := readInt(reader)

	// Генерируем сообщения заданного размера
	messages := generateMessages(numMessages, messageSize)
	return messages, numWriters, numReaders
}

// readInt читает и преобразует ввод в int
func readInt(reader *bufio.Reader) int {
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Удаляем лишние пробелы и символы новой строки
		value, err := strconv.Atoi(input)
		if err == nil && value > 0 {
			return value
		}
		fmt.Println("Пожалуйста, введите положительное число:")
	}
}

// generateMessages генерирует список сообщений заданного размера
func generateMessages(numMessages, messageSize int) []string {
	messages := make([]string, numMessages)
	message := make([]byte, messageSize)
	for i := range message {
		message[i] = 'a'
	}
	for i := range messages {
		messages[i] = string(message)
	}
	return messages
}
