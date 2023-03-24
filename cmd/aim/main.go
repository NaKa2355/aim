package main

import (
	"context"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/controllers/data_access"
	"github.com/NaKa2355/aim/internal/app/aim/controllers/presenter"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/interactor"
)

func main() {
	fmt.Println("Hello, World!")
	d, err := data_access.New("./test.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer d.Close()
	i := interactor.New(d, presenter.StdOut{})
	i.GetAppliances(context.Background())
}
