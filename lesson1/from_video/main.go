/*
Написать програмку

Функция которая создает файл на файловой системе
используя отложенный вызов для безопасного закрытия файла

сгенерировать когда функция не может создать файл - это будет неявная ошибка,
эту ошибку нужно обработать и написать информацию не только от том что произошла ошибка, а добавить дату выполнения (или доп информацию).

*/

package main

import (
	"fmt"
	"os"
	"time"
)

// My Custom Error
type ErrorWithTimestamp struct {
	text      string
	timestamp string
}

func MyError(text string) error {
	return &ErrorWithTimestamp{
		text:      text,
		timestamp: time.Now().String(),
	}
}

func (e *ErrorWithTimestamp) Error() string {
	return fmt.Sprintf("Error: %s\nTimestamp: %s", e.text, e.timestamp)
}

var tempDir = "temp_files"

// Создаем файл
func createFile() {
	f, err := os.Create(tempDir + "/" + "test_file")
	if err != nil {
		err := MyError("Can not create file")
		fmt.Println(err)
	}

	defer f.Close()
}

func main() {
	createFile()

}
