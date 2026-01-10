package logger

import (
	"fmt"
	"os"
	"time"
)

// Clears logs
func LoggerIni() {
	os.WriteFile("../log.txt", []byte{}, 0644)
}

// Adds a record to the last line of the logs
func Log(text string) {
	//Opening a file
	file, err := os.OpenFile("../log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка откртия файла для логирования . . .")
	}
	defer file.Close()
	//Text formatting
	layout := "02.01.2006 15:04"
	text = time.Now().Format(layout) + "  |  " + text
	//Log entry
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("ошибка логирования . . .")
	}
}
