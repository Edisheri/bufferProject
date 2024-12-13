package performance

import (
	"fmt"
	"time"
)

// MeasurePerformance замеряет время выполнения функции синхронизации, выводит анализ и обоснование
func MeasurePerformance(syncFunc func([]string, int, int), messages []string, numWriters, numReaders int) {
	start := time.Now()
	syncFunc(messages, numWriters, numReaders)
	duration := time.Since(start)

	// Вывод результатов производительности
	fmt.Printf("Время выполнения: %v\n", duration)
	fmt.Printf("Число сообщений: %d, Число писателей: %d, Число читателей: %d, Размер сообщений: %d байт\n",
		len(messages), numWriters, numReaders, len(messages[0]))

	// Анализ производительности
	fmt.Println("Анализ эффективности:")
	if numWriters > numReaders {
		fmt.Println("- Высокая конкуренция на запись.")
		fmt.Println("  Обоснование: Число писателей превышает число читателей, что увеличивает задержки при записи.")
	} else if numReaders > numWriters {
		fmt.Println("- Высокая конкуренция на чтение.")
		fmt.Println("  Обоснование: Число читателей превышает число писателей, что приводит к простою буфера при ожидании новых данных.")
	} else {
		fmt.Println("- Баланс между чтением и записью.")
		fmt.Println("  Обоснование: Число писателей и читателей одинаково, что минимизирует задержки и оптимизирует использование буфера.")
	}

	fmt.Println("Подходит для:")
	if duration < time.Millisecond*10 {
		fmt.Println("- Малых объемов сообщений.")
		fmt.Println("  Обоснование: Время выполнения меньше 10 мс, что позволяет эффективно обрабатывать небольшие данные.")
	} else if duration < time.Millisecond*100 {
		fmt.Println("- Средних объемов сообщений.")
		fmt.Println("  Обоснование: Время выполнения находится в пределах 10–100 мс, что оптимально для обработки данных средней величины.")
	} else {
		fmt.Println("- Больших объемов сообщений.")
		fmt.Println("  Обоснование: Время выполнения превышает 100 мс, что свидетельствует о необходимости обработки больших объемов данных.")
	}
	fmt.Println("---------------------------------------------------------")
}
