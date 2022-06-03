#### Сохранить трассировку
go run main.go 2>trace.out

#### Проверить на наличие гонки
go run -race main.go