package routes

import (
	"github.com/delos/aquafarm-management/src/controllers"
	"github.com/delos/aquafarm-management/src/repositories"
	"github.com/delos/aquafarm-management/src/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func SetupRoutes(router *mux.Router, db *gorm.DB) {
	farmRepository := repositories.NewFarmRepository(db)
	pondRepository := repositories.NewPondRepository(db)
	statisticRepository := repositories.NewStatisticRepository(db)
	userRepository := repositories.NewUserRepository(db)

	farmService := services.NewFarmService(farmRepository)
	pondService := services.NewPondService(pondRepository, farmRepository)
	statisticService := services.NewStatisticService(statisticRepository)
	userService := services.NewUserService(userRepository)

	controller := controllers.NewController(farmService, pondService, statisticService, userService)

	aquafarm := router.PathPrefix("/v1/aquafarm").Subrouter()

	// User Endpoints
	aquafarm.HandleFunc("/user", controller.CreateUser).Methods(http.MethodPost)

	// Farms Endpoints
	aquafarm.HandleFunc("/farm", controller.CreateFarm).Methods(http.MethodPost)
	aquafarm.HandleFunc("/farm/{farmID}", controller.UpdateFarm).Methods(http.MethodPut)
	aquafarm.HandleFunc("/farm/{farmID}", controller.DeleteFarm).Methods(http.MethodDelete)
	aquafarm.HandleFunc("/farms", controller.GetFarms).Methods(http.MethodGet)
	aquafarm.HandleFunc("/farm/{farmID}", controller.GetFarmByID).Methods(http.MethodGet)

	// Ponds Endpoints
	aquafarm.HandleFunc("/pond", controller.CreatePond).Methods(http.MethodPost)
	aquafarm.HandleFunc("/pond/{pondID}", controller.UpdatePond).Methods(http.MethodPut)
	aquafarm.HandleFunc("/pond/{pondID}", controller.DeletePond).Methods(http.MethodDelete)
	aquafarm.HandleFunc("/ponds", controller.GetPonds).Methods(http.MethodGet)
	aquafarm.HandleFunc("/pond/{pondID}", controller.GetPondByID).Methods(http.MethodGet)

	// Statistics Endpoints
	aquafarm.HandleFunc("/statistics", controller.GetStatistics).Methods(http.MethodGet)
}
