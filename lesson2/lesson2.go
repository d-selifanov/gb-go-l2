package lesson2

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Customizing error
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

// Task 1. Get panic and recovery later
func PanicAndRecover() {
	var a int
	var b int

	a = 1
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Recovered: ", v)
		}
	}()

	fmt.Println(a / b)
}

// Task 2. Get panic and recovery later.
// Usage customized error with timestamp
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

	fmt.Println(a / b)
}

// Task 3. Func for generate one millions files.
func CreateFiles() {
	_ = os.Mkdir("temp_files", 0700)
	for i := 0; i < 1000000; i++ {
		f, _ := os.Create("temp_files/file_" + strconv.Itoa(i+1))
		defer f.Close()
	}
}

// Task 4. Handling panic in goroutine
func PanicInParallelStream() {
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
