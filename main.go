package main

import (
	"net/http"
	"oms/oms/core"
	"oms/oms/repo"
	"oms/oms/server"
	"oms/providers"
	"oms/routines"
	"time"
)

func main() {
	dbInstance := providers.GetDbConnection()
	repoInstance := repo.Repo{DB: dbInstance}
	coreInstance := core.Core{Repo: repoInstance}
	serverInstance := server.Server{Core: coreInstance}
	RegisterRoutes(serverInstance)

	orderConfig := providers.GetOrderConfig()
	go routines.StartOrderCreation(repoInstance, orderConfig)
	time.Sleep(5 * time.Second) // Wait for orders to be created before starting order updater
	go routines.StartPendingOrderUpdater(repoInstance, orderConfig)
	go routines.StartProcessingOrderUpdater(repoInstance, orderConfig)

	http.ListenAndServe(":8080", nil)
}

func RegisterRoutes(serverInstance server.Server) {
	// http.HandleFunc("/order", serverInstance.HandleOrder())
	http.HandleFunc("/order_metrics", serverInstance.HandleOrderMetrics())
}
