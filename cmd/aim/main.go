package main

import (
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/controllers/data_access"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/interactor"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Hello, World!")
	d, err := data_access.New("./test.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer d.Close()
	interactor.New(d)
}
