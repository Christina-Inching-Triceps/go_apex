package main

import (
	"gopex/adapter/controller"
	"gopex/domain/application"
	apiApplication "gopex/domain/application/api"
	"gopex/infrastructure"
	"gopex/infrastructure/config"
	"gopex/infrastructure/external"
	"gopex/infrastructure/persistence"

	"github.com/labstack/echo/v4"
)

func main() {
	// ===== Init Echo Server
	e := echo.New()

	// ===== Load Config File
	config, err := config.Load()
	if err != nil {
		e.Logger.Fatal(err)
	}

	// ===== Init DB Instance
	// See: https://please-sleep.cou929.nu/go-sql-db-connection-pool.html
	db, err := infrastructure.NewDB(config, infrastructure.OpenMySQLDatabase(config))
	if err != nil {
		e.Logger.Fatal(err)
	}

	apexTrackerRepository := persistence.NewApexTrackerRepository(db)
	apexTrackerClient := external.NewApexTrackerApi(config)

	// FIXME: apiのinteractorをcontrollerに仕込んで逐次実行して

	// TODO: create usecase
	userReadApexProfileSummaryInteractor := application.NewUserReadApexProfileSummaryInteractor(&apexTrackerRepository)
	getApexTrackerStatsInteractor := apiApplication.NewGetApexTrackerStatsInteractor(&apexTrackerClient)

	// ===== Setup Router
	apexTrackerController := controller.NewApexController(&userReadApexProfileSummaryInteractor, &getApexTrackerStatsInteractor, config)
	infrastructure.InitApexTrackerRouting(e, config, apexTrackerController)

	e.Logger.Info(e.Start(":1323"))
}
