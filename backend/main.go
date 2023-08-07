package main

import (
	"net/http"
)

func main() {
	logger := InitializeLogger()
	config := LoadConfig(logger)
	store := InitializeDataStore(config, logger)
	r := SetupRoutes(config, store, logger)

	// Run server
	logger.Fatal(http.ListenAndServe(":"+config.Server.Port, r))
}
