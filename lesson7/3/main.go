/*
3. *не обязательное*. Написать кодогенератор под какую-нибудь задачу.
*/
package main

import (
	"github.com/google/uuid"
)

//go:generate go run ./generate/main.go
//go:generate goimports -w ./generated_file.go

type Product struct {
	Code  uuid.UUID
	Name  string
	Price float64
	Count int64
}

func main() {

}
