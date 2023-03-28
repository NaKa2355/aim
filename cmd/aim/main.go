package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/NaKa2355/aim/internal/app/aim/controllers/data_access"
	"github.com/NaKa2355/aim/internal/app/aim/controllers/web"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/web/server"
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

	i := interactor.New(d)
	h := web.NewHandler(i)
	s := server.New(h)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	s.Start(listener)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	s.Stop()
	fmt.Println("")
}
