package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/delos/aquafarm-management/src/models"
	"github.com/delos/aquafarm-management/src/request"
	"github.com/delos/aquafarm-management/src/response"
	"github.com/delos/aquafarm-management/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type Controller struct {
	FarmService      services.FarmService
	PondService      services.PondService
	StatisticService services.StatisticService
	UserService      services.UserService
}

func NewController(
	farmService services.FarmService,
	pondService services.PondService,
	statisticService services.StatisticService,
	userService services.UserService) *Controller {
	return &Controller{
		FarmService:      farmService,
		PondService:      pondService,
		StatisticService: statisticService,
		UserService:      userService,
	}
}

// FARM

func (c *Controller) CreateFarm(w http.ResponseWriter, r *http.Request) {
	payload := request.UpsertFarmRequest{}
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Request Validation
	validate := validator.New()
	err = validate.Struct(payload)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Record the statistic API call
	err = c.storeStatisticApiCall(r, payload)
	if err != nil {
		log.Print("Failed to store data statistic api call")
	}

	// Create the farm
	farm := &models.Farm{
		FarmName:  payload.FarmName,
		UserID:    payload.UserID,
		CreatedAt: time.Now(),
	}
	err = c.FarmService.CreateFarm(ctx, farm)
	if err != nil {
		response.ConstructResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	response.ConstructResponse(w, http.StatusCreated, farm, nil)
}

func (c *Controller) UpdateFarm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	farmID := mux.Vars(r)["farmID"]
	var farmReq request.UpsertFarmRequest
	err := json.NewDecoder(r.Body).Decode(&farmReq)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Request Validation
	validate := validator.New()
	err = validate.Struct(farmReq)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	farmData := &models.Farm{
		ID:        farmID,
		FarmName:  farmReq.FarmName,
		UserID:    farmReq.UserID,
		UpdatedAt: time.Now(),
	}

	// Record the statistic API call
	err = c.storeStatisticApiCall(r, farmReq)
	if err != nil {
		log.Print("Failed to store data statistic api call")
	}

	farm, err := c.FarmService.UpdateFarm(ctx, farmData)
	if err != nil {
		response.ConstructResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	response.ConstructResponse(w, http.StatusOK, farm, nil)
}

func (c *Controller) DeleteFarm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.DeleteFarmRequest
	farmID := mux.Vars(r)["farmID"]
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Validasi menggunakan validator
	req.FarmID = farmID
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Record the statistic API call
	err = c.storeStatisticApiCall(r, req)
	if err != nil {
		log.Print("Failed to store data statistic api call")
	}

	farmData := &models.Farm{
		ID:     req.FarmID,
		UserID: req.UserID,
	}

	err = c.FarmService.DeleteFarm(ctx, farmData)
	if err != nil {
		response.ConstructResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	response.ConstructResponse(w, http.StatusOK, nil, nil)
}

func (c *Controller) GetFarms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.URL.Query().Get("user_id")
	var req request.GetFarmsAndPondsRequest

	// Validasi menggunakan validator
	req.UserID = userID
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Record the statistic API call
	err = c.storeStatisticApiCall(r, req)
	if err != nil {
		log.Print("Failed to store data statistic api call")
	}

	farms, err := c.FarmService.GetAllFarmByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ConstructResponse(w, http.StatusNotFound, nil, err)
			return
		}
		response.ConstructResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	response.ConstructResponse(w, http.StatusOK, farms, nil)
}

func (c *Controller) GetFarmByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	farmID := mux.Vars(r)["farmID"]
	userID := r.URL.Query().Get("user_id")
	var req request.GetFarmsAndPondsRequest

	// Validasi menggunakan validator
	req.UserID = userID
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Record the statistic API call
	err = c.storeStatisticApiCall(r, req)
	if err != nil {
		log.Print("Failed to store data statistic api call")
	}

	farm, err := c.FarmService.GetFarmByID(ctx, farmID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ConstructResponse(w, http.StatusNotFound, nil, err)
			return
		}
		response.ConstructResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	response.ConstructResponse(w, http.StatusOK, farm, nil)
}

// POND

func (c *Controller) CreatePond(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload request.UpsertPondRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Request Validation
	validate := validator.New()
	err = validate.Struct(payload)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Record the statistic API call
	err = c.storeStatisticApiCall(r, payload)
	if err != nil {
		log.Print("Failed to store data statistic api call")
	}

	// Create the pond
	pond := &models.Pond{
		PondName:  payload.PondName,
		FarmID:    payload.FarmID,
		CreatedAt: time.Now(),
	}
	err = c.PondService.CreatePond(ctx, pond)
	if err != nil {
		response.ConstructResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	response.ConstructResponse(w, http.StatusCreated, pond, nil)
}

func (c *Controller) UpdatePond(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pondID := mux.Vars(r)["pondID"]

	var pondReq request.UpsertPondRequest
	err := json.NewDecoder(r.Body).Decode(&pondReq)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Request Validation
	validate := validator.New()
	err = validate.Struct(pondReq)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Record the statistic API call
	err = c.storeStatisticApiCall(r, pondReq)
	if err != nil {
		log.Print("Failed to store data statistic api call")
	}

	pondData := &models.Pond{
		ID:        pondID,
		PondName:  pondReq.PondName,
		FarmID:    pondReq.FarmID,
		UpdatedAt: time.Now(),
	}

	pond, err := c.PondService.UpdatePond(ctx, pondData)
	if err != nil {
		response.ConstructResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	response.ConstructResponse(w, http.StatusOK, pond, nil)
}

func (c *Controller) DeletePond(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.DeletePondRequest
	pondID := mux.Vars(r)["pondID"]
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Validasi menggunakan validator
	req.PondID = pondID
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Record the statistic API call
	err = c.storeStatisticApiCall(r, req)
	if err != nil {
		log.Print("Failed to store data statistic api call")
	}

	pondData := &models.Pond{
		ID: req.PondID,
	}

	err = c.PondService.DeletePond(ctx, pondData, req.UserID)
	if err != nil {
		response.ConstructResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	response.ConstructResponse(w, http.StatusOK, nil, nil)
}

func (c *Controller) GetPonds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := r.URL.Query().Get("user_id")
	var req request.GetFarmsAndPondsRequest

	// Validasi menggunakan validator
	req.UserID = userID
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Record the statistic API call
	err = c.storeStatisticApiCall(r, req)
	if err != nil {
		log.Print("Failed to store data statistic api call")
	}

	ponds, err := c.PondService.GetPondsByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ConstructResponse(w, http.StatusNotFound, nil, err)
			return
		}
		response.ConstructResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	response.ConstructResponse(w, http.StatusOK, ponds, nil)
}

func (c *Controller) GetPondByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pondID := mux.Vars(r)["pondID"]
	userID := r.URL.Query().Get("user_id")
	var req request.GetFarmsAndPondsRequest

	// Validasi menggunakan validator
	req.UserID = userID
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Record the statistic API call
	err = c.storeStatisticApiCall(r, req)
	if err != nil {
		log.Print("Failed to store data statistic api call")
	}

	farm, err := c.PondService.GetPondByID(ctx, pondID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ConstructResponse(w, http.StatusNotFound, nil, err)
			return
		}
		response.ConstructResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	response.ConstructResponse(w, http.StatusOK, farm, nil)
}

// USER

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userReq request.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	// Request Validation
	validate := validator.New()
	err = validate.Struct(userReq)
	if err != nil {
		response.ConstructResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	user := &models.User{
		UserName:  userReq.Username,
		CreatedAt: time.Now(),
	}

	err = c.UserService.CreateUser(ctx, user)
	if err != nil {
		response.ConstructResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	response.ConstructResponse(w, http.StatusOK, user, nil)
}

func (c *Controller) storeStatisticApiCall(r *http.Request, payload interface{}) error {
	ctx := r.Context()

	statistic := &models.Statistic{
		ID:        models.GenerateID(),
		Endpoint:  r.Method + r.RequestURI,
		Count:     1,
		CallAt:    time.Now(),
		CreatedAt: time.Now(),
	}

	switch payload.(type) {
	case request.UpsertFarmRequest:
		req := payload.(request.UpsertFarmRequest)
		statistic.UserID = req.UserID
	case request.UpsertPondRequest:
		req := payload.(request.UpsertPondRequest)
		statistic.UserID = req.UserID
	case request.DeleteFarmRequest:
		req := payload.(request.DeleteFarmRequest)
		statistic.UserID = req.UserID
	case request.DeletePondRequest:
		req := payload.(request.DeletePondRequest)
		statistic.UserID = req.UserID
	case request.GetFarmsAndPondsRequest:
		req := payload.(request.GetFarmsAndPondsRequest)
		statistic.UserID = req.UserID
	}

	err := c.createDataStatistic(ctx, statistic)
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) createDataStatistic(ctx context.Context, statistic *models.Statistic) error {
	err := c.StatisticService.CreateStatistic(ctx, statistic)
	if err != nil {
		return err
	}

	return nil
}

// STATISTICS
func (c *Controller) GetStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	statistics, err := c.StatisticService.GetStatistics(ctx)
	if err != nil {
		response.ConstructResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	responseData := response.StatisticsResponse{
		Endpoints: make(map[string]response.EndpointStatistics),
	}

	for _, stat := range statistics {
		endpointStats, exists := responseData.Endpoints[stat.Endpoint]
		if !exists {
			endpointStats = response.EndpointStatistics{}
		}
		endpointStats.Count += stat.Count
		endpointStats.UniqueUserAgent++
		responseData.Endpoints[stat.Endpoint] = endpointStats
	}

	response.ConstructResponse(w, http.StatusOK, responseData, nil)
}
