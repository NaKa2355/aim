package main

import (
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/daemon"
	"golang.org/x/exp/slog"
)

func main() {
	fmt.Println("Hello, World!")
	logger := slog.New(slog.Default().Handler())
	daemon, err := daemon.NewWithoutConfig("./test.db", true, logger)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = daemon.StartWithNet("tcp", ":3000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
