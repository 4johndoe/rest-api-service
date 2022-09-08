package main

import (
	"banking/app"
	"banking/logger"
)

func main() {

	logger.Log.Info("Starting our app...")
	app.Start()
}
