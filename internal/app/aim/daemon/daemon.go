package daemon

import (
	"context"
	"encoding/json"
	"net"
	"os"
	"syscall"

	"github.com/NaKa2355/aim/internal/app/aim/controllers/dataAccess"
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

func NewWithoutConfig(dbFilePath string, enableReflection bool, logger *slog.Logger) (*Daemon, error) {
	var err error = nil
	var d = &Daemon{}
	d.logger = logger

	repo, err := dataAccess.New(dbFilePath)
	if err != nil {
		d.logger.Error(
			"faild to access to database",
			"error", err.Error(),
		)
		return d, err
	}

	streamingContext, closeAllStreaming := context.WithCancel(context.Background())

	h := handler.New(streamingContext)
	i := interactor.New(repo, h)
	h.SetInteractor(i)
	d.srv = server.New(h, enableReflection, closeAllStreaming)
	return d, nil
}

func New(configPath string, dbFilePath string, logger *slog.Logger) (*Daemon, error) {
	var err error = nil
	var d = &Daemon{}

	config, err := d.readConf(configPath)
	if err != nil {
		d.logger.Error(
			"faild to read config file",
			"error", err.Error(),
		)
		return d, err
	}

	return NewWithoutConfig(dbFilePath, config.EnableReflection, logger)
}

func (d *Daemon) start(listener net.Listener) error {
	d.srv.Start(listener)
	d.srv.WaitSigAndStop(syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	d.logger.Info("shutting down daemon...")
	d.logger.Info("stopped daemon")
	return nil
}

func (d *Daemon) StartWithDomainSocket(domainSocketPath string) error {

	listener, err := net.Listen("unix", domainSocketPath)
	if err != nil {
		d.logger.Error("faild to make a socket", "error", err)
		return err
	}

	err = os.Chmod(domainSocketPath, 0770)
	if err != nil {
		d.logger.Error("faild to change permisson", "error", err)
		return err
	}

	d.logger.Info(
		"daemon starting",
		"unix domain socket path", domainSocketPath,
	)

	return d.start(listener)
}

func (d *Daemon) StartWithNet(network string, address string) error {
	listener, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	return d.start(listener)
}
