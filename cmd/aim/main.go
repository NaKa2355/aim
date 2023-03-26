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
	i.DeleteAppliance(context.Background(), boundary.DeleteAppInput{AppID: "01GW8REDZFS5WNR99FB85MDHHB"})
	i.GetButton(context.Background(), boundary.GetAppInput{AppID: "01GW8REDZFS5WNR99FB85MDHHB"})
	i.GetAppliances(context.Background())
}
