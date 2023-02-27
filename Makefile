MODULE := $(shell go list -m)

TARGET := $(MODULE)/cmd/aim
GO_OUT := bin/

build:
	go build -o $(GO_OUT) $(TARGET)