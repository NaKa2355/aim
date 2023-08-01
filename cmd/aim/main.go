package main

import (
	"os"
	"os/user"

	"github.com/NaKa2355/aim/internal/app/aim/daemon"
	"golang.org/x/exp/slog"
)

const ConfigFilePath = "/etc/aimd.json"
const DomainSocketPath = "/tmp/aimd/aimd.sock"

func main() {
	logger := slog.New(slog.Default().Handler())
	user, err := user.Current()
	if err != nil {
		logger.Error(
			"faild to get current user",
			"error", err.Error(),
		)
		os.Exit(-1)
	}

	dbFilePath := user.HomeDir + "/.aim.db"
	d, err := daemon.New(ConfigFilePath, dbFilePath, logger)
	if err != nil {
		os.Exit(-1)
	}

	err = d.Start(DomainSocketPath)
	if err != nil {
		os.Exit(-1)
	}
}
