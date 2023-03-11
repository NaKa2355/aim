package main

import (
	"context"
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
	i := interactor.New(d)
	a, err := i.AddCustom(context.Background(), "おまんこ", "10")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.Name, a.ID)
	apps, err := i.GetAppliances(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, app := range apps {
		fmt.Println(app.Name)
	}

	//i.RenameAppliance(context.Background(), "01GV7XC8GFMQR5D9YJMBDZ9EQF", "ちんぽ")
	//err = i.RenameAppliance(context.Background(), "01GV2B3W95SAWV1E5HH1RS6ZX9", "うんこ")
	err = i.AddCommand(context.Background(), "01GV803PDAZ6M7ZTFDVW98CEMY", "aaaa")
	fmt.Println(err)
	/*coms, err := i.GetCommands(context.Background(), "01GV7XC8GFMQR5D9YJMBDZ9EQF")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range coms {
		fmt.Println(c.Name)
	}*/
}
