package tests

// GenerateMessages создает массив сообщений для тестов
func GenerateMessages(numMessages, messageSize int) []string {
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
