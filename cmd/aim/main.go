package main

import (
	"context"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/controllers/data_access"
	"github.com/NaKa2355/aim/internal/app/aim/controllers/presenter"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
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
	i.AddButton(context.Background(), boundary.AddButtonInput{Name: "hello", DeviceID: "wawawa"})
	i.AddCommand(context.Background(), boundary.AddCommandInput{AppID: "01GW9ZF4CXPJ89S95KZW3MHDZF", Name: "Hello"})
	i.GetAppliances(context.Background())
}
