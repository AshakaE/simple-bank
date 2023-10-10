package main

import (
	"github.com/AshakaE/banking/app"
	"github.com/AshakaE/banking/logger"
)

func main() {
	logger.Info("Starting app...")
	app.Start()
}
