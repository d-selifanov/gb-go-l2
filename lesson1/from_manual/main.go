/*
1) Для закрепления навыков отложенного вызова функций, напишите программу, содержащую вызов функции,
которая будет создавать паническую ситуацию неявно. Затем создайте отложенный вызов,
который будет обрабатывать эту паническую ситуацию и, в частности, печатать предупреждение в консоль.
Критерием успешного выполнения задания является то, что программа не завершается аварийно ни при каких условиях.

2) Дополните функцию из п.1 возвратом собственной ошибки в случае возникновения панической ситуации.
Собственная ошибка должна хранить время обнаружения панической ситуации.
Критерием успешного выполнения задания является наличие обработки созданной ошибки в функции
main и вывод ее состояния в консоль

3) Для закрепления практических навыков программирования, напишите программу,
которая создаёт один миллион пустых файлов в известной, пустой директории файловой системы
используя вызов os.Create.
Ввиду наличия определенных ограничений операционной системы на число открытых файлов,
такая программа должна выполнять аварийную остановку. Запустите программу и дождитесь полученной ошибки.
Используя отложенный вызов функции закрытия файла, стабилизируйте работу приложения.
Критерием успешного выполнения программы является успешное создание миллиона пустых файлов в директории

4) Для самостоятельного изучения. Предложите реализацию примера так, чтобы аварийная остановка программы не выполнилась
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// MyError
type ErrorWithTimestamp struct {
	text string
	timestamp string
}
func MyError(text string) error {
	return &ErrorWithTimestamp {
		text: text,
		timestamp: time.Now().String(),
	}
}
func (e *ErrorWithTimestamp) Error() string {
	return fmt.Sprintf("Error: %s\nTimestamp: %s", e.text, e.timestamp)
}

// 1
func PanicAndRecover() {
	var a int
	var b int

	a = 1
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Recovered: ", v)
		}
	}()

	fmt.Println(a/b)
}

// 2
func PanicAndRecoverWithTimestamp() {
	var a int
	var b int

	a = 1

	defer func() {
		if v := recover(); v != nil {
			err := MyError("My Error from PANIC")
			fmt.Printf("%v\n", err)
		}
	}()

	fmt.Println(a/b)
}

// 3 
func createFiles() {
	_ = os.Mkdir("temp_files", 0700)
	for i := 0; i < 1000000; i++  {
		f, _ := os.Create("temp_files/file_" + strconv.Itoa(i+1))
		defer f.Close()
	}
}

// 4 
func panicInParallelStream() {
	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("Recover panic in Parallel Stream:", v)
			}
		}()
		panic("A-A-A!!!")
	}()
	time.Sleep(time.Second)
}

func main() {
	PanicAndRecover()
	PanicAndRecoverWithTimestamp()
	createFiles()
	panicInParallelStream()
}
