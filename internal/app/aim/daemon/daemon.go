package daemon

import (
	"encoding/json"
	"net"
	"os"
	"syscall"

	"github.com/NaKa2355/aim/internal/app/aim/controllers/data_access"
	"github.com/NaKa2355/aim/internal/app/aim/controllers/web/handler"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/web/server"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/interactor"
	"golang.org/x/exp/slog"
)

type Config struct {
	EnableReflection bool `json:"enable_reflection"`
}

type Daemon struct {
	logger *slog.Logger
	srv    *server.Server
}

func (d *Daemon) readConf(filePath string) (*Config, error) {
	config := &Config{}
	config_data, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(config_data, config)
	return config, err
}

func New(configPath string, dbFilePath string, logger *slog.Logger) (*Daemon, error) {
	var err error = nil
	var d = &Daemon{}
	d.logger = logger

	config, err := d.readConf(configPath)
	if err != nil {
		d.logger.Error(
			"faild to read config file",
			"error", err.Error(),
		)
		return d, err
	}

	repo, err := data_access.New(dbFilePath)
	if err != nil {
		d.logger.Error(
			"faild to access to database",
			"error", err.Error(),
		)
		return d, err
	}

	h := handler.New()
	i := interactor.New(repo, h)
	h.SetInteractor(i)
	d.srv = server.New(h, config.EnableReflection)
	return d, nil
}

func (d *Daemon) Start(domainSocket string) error {
	listener, err := net.Listen("unix", domainSocket)
	if err != nil {
		d.logger.Error("faild to make a socket", "error", err)
		return err
	}

	err = os.Chmod(domainSocket, 0770)
	if err != nil {
		d.logger.Error("faild to change permisson", "error", err)
		return err
	}

	d.srv.Start(listener)

	d.logger.Info(
		"daemon started",
		"unix domain socket path", domainSocket,
	)

	d.srv.WaitSigAndStop(syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	d.logger.Info("shutting down daemon...")
	d.logger.Info("stopped daemon")
	return nil
}
