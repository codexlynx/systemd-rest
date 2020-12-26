package core

import (
	"fmt"
	"os"
)

type Mode int

const (
	DebugMode   = iota
	ReleaseMode = iota
)

func GetConfiguration() (Mode, string) {
	var mode Mode
	if os.Getenv("MODE") == "debug" {
		mode = DebugMode
	} else {
		mode = ReleaseMode
	}

	address := os.Getenv("ADDRESS")
	if len(address) == 0 {
		address = "127.0.0.1"
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "6789"
	}

	listenAddress := fmt.Sprintf("%s:%s", address, port)
	return mode, listenAddress
}

