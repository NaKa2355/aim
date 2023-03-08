package main

import (
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/controllers/data_access"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
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

	err = d.CreateTable()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = d.AddAppTypeQuery()
	if err != nil {
		fmt.Println(err)
		//return
	}

	t, err := appliance.NewThermostat("エアコ", "10", 0.5, 10, 20, 10, 20)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = d.SaveApp(t)
	if err != nil {
		fmt.Println(err)
	}

	apps, err := d.GetAppList()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, app := range apps {
		fmt.Println(app.GetName())
	}
	fmt.Println(apps[0].GetOpt())
}
