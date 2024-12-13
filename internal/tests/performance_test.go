package tests

import (
	"bufferProject/internal/synchronized"
	"testing"
	"time"
)

func TestPerformanceScenarios(t *testing.T) {
	scenarios := []struct {
		name        string
		numMessages int
		messageSize int
		numWriters  int
		numReaders  int
	}{
		{"Маленькая нагрузка", 10, 10, 1, 1},
		{"Средняя нагрузка", 100, 100, 2, 2},
		{"Большая нагрузка", 1000, 500, 5, 5},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			messages := GenerateMessages(scenario.numMessages, scenario.messageSize)

			t.Logf("Сценарий: %d сообщений, размер %d байт, %d писателей, %d читателей",
				scenario.numMessages, scenario.messageSize, scenario.numWriters, scenario.numReaders)

			t.Log("\n--- Синхронизация с использованием WaitGroup ---")
			measureAndLogPerformance(t, synchronized.WaitGroupSync, messages, scenario.numWriters, scenario.numReaders)

			t.Log("\n--- Синхронизация с использованием Channel ---")
			measureAndLogPerformance(t, synchronized.ChannelSync, messages, scenario.numWriters, scenario.numReaders)

			t.Log("\n--- Синхронизация с использованием Mutex ---")
			measureAndLogPerformance(t, synchronized.MutexSync, messages, scenario.numWriters, scenario.numReaders)
		})
	}
}

func measureAndLogPerformance(t *testing.T, syncFunc func([]string, int, int), messages []string, numWriters, numReaders int) {
	start := time.Now()
	syncFunc(messages, numWriters, numReaders)
	duration := time.Since(start)

	t.Logf("Время выполнения: %v", duration)
}
