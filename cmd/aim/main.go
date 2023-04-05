package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/NaKa2355/aim/internal/app/aim/controllers/data_access"
	"github.com/NaKa2355/aim/internal/app/aim/controllers/web"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/web/server"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/interactor"
)

func main() {
	var mem runtime.MemStats
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
	t := time.NewTicker(1 * time.Second)
L:
	for {
		select {
		case <-t.C:
			runtime.ReadMemStats(&mem)
			fmt.Println(mem.StackSys, mem.HeapSys, mem.Sys, float32(mem.Alloc)/float32(mem.HeapSys))
		case <-quit:
			break L
		}
	}

	s.Stop()
	fmt.Println("")
}
