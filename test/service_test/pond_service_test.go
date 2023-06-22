package service_test

import (
	"context"
	"errors"
	"github.com/delos/aquafarm-management/src/models"
	"github.com/delos/aquafarm-management/src/services"
	"github.com/delos/aquafarm-management/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestPondService_CreatePond(t *testing.T) {
	pondRepo := &mocks.PondRepositoryMock{}
	farmRepo := &mocks.FarmRepositoryMock{}
	pondService := services.NewPondService(pondRepo, farmRepo)

	ctx := context.Background()

	// Define test case
	testCases := []struct {
		name          string
		pond          *models.Pond
		expectedError error
	}{
		{
			name: "Create pond successfully",
			pond: &models.Pond{
				ID:       "pondID",
				PondName: "Pond 1",
				FarmID:   "farmID",
			},
			expectedError: nil,
		},
		{
			name: "Pond with the same name already exists",
			pond: &models.Pond{
				ID:       "pondID",
				PondName: "Pond 1",
				FarmID:   "farmID",
			},
			expectedError: errors.New("pond with the same name already exists"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Set up mock behavior
			if testCase.expectedError == nil {
				pondRepo.On("GetPondByName", ctx, testCase.pond.PondName).Return(nil, gorm.ErrRecordNotFound).Once()
			} else {
				pondRepo.On("GetPondByName", ctx, testCase.pond.PondName).Return(&models.Pond{}, testCase.expectedError).Once()
			}
			pondRepo.On("CreatePond", ctx, testCase.pond).Return(nil).Once().Run(func(args mock.Arguments) {
				pondArg := args.Get(1).(*models.Pond)
				pondArg.ID = "generatedID" // Assign a generated ID to the pond argument
			})

			err := pondService.CreatePond(ctx, testCase.pond)
			assert.Equal(t, testCase.expectedError, err)

			// Verify that the expected methods are called
			pondRepo.AssertCalled(t, "GetPondByName", ctx, testCase.pond.PondName)
			pondRepo.AssertCalled(t, "CreatePond", ctx, mock.AnythingOfType("*models.Pond"))
		})
	}
}

func TestPondService_UpdatePond(t *testing.T) {
	pondRepo := &mocks.PondRepositoryMock{}
	farmRepo := &mocks.FarmRepositoryMock{}
	pondService := services.NewPondService(pondRepo, farmRepo)

	ctx := context.Background()

	// Define test case
	testCases := []struct {
		name          string
		pondID        string
		pondData      *models.Pond
		expectedError error
	}{
		{
			name:   "Update pond successfully",
			pondID: "pondID",
			pondData: &models.Pond{
				ID:        "pondID",
				PondName:  "Pond 1",
				FarmID:    "farmID",
				UpdatedAt: time.Now(),
			},
			expectedError: nil,
		},
		{
			name:   "Pond not found",
			pondID: "pondID",
			pondData: &models.Pond{
				ID:        "pondID",
				PondName:  "Pond 1",
				FarmID:    "farmID",
				UpdatedAt: time.Now(),
			},
			expectedError: errors.New("pond not found"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Set up mock behavior
			pondRepo.On("GetPondByID", ctx, testCase.pondID).Return(testCase.pondData, testCase.expectedError).Once()
			if testCase.expectedError == nil {
				pondRepo.On("UpdatePond", ctx, testCase.pondData).Return(nil).Once()
			}

			_, err := pondService.UpdatePond(ctx, testCase.pondData)
			assert.Equal(t, testCase.expectedError, err)

			// Verify that the expected methods are called
			pondRepo.AssertCalled(t, "GetPondByID", ctx, testCase.pondID)
			if testCase.expectedError == nil {
				pondRepo.AssertCalled(t, "UpdatePond", ctx, testCase.pondData)
			}
		})
	}
}

func TestPondService_GetPondByID(t *testing.T) {
	pondRepo := &mocks.PondRepositoryMock{}
	farmRepo := &mocks.FarmRepositoryMock{}
	pondService := services.NewPondService(pondRepo, farmRepo)

	ctx := context.Background()

	// Define test case
	testCases := []struct {
		name          string
		pondID        string
		mockResponse  *models.Pond
		expectedError error
	}{
		{
			name:   "Get pond by ID successfully",
			pondID: "pondID",
			mockResponse: &models.Pond{
				ID:        "pondID",
				PondName:  "Pond 1",
				FarmID:    "farmID",
				UpdatedAt: time.Now(),
			},
			expectedError: nil,
		},
		{
			name:          "Pond not found",
			pondID:        "pondID",
			mockResponse:  nil,
			expectedError: errors.New("pond not found"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Set up mock behavior
			pondRepo.On("GetPondByID", ctx, testCase.pondID).Return(testCase.mockResponse, testCase.expectedError).Once()

			pond, err := pondService.GetPondByID(ctx, testCase.pondID)
			assert.Equal(t, testCase.mockResponse, pond)
			assert.Equal(t, testCase.expectedError, err)

			// Verify that the expected method is called
			pondRepo.AssertCalled(t, "GetPondByID", ctx, testCase.pondID)
		})
	}
}

func TestPondService_GetPondsByUserID(t *testing.T) {
	pondRepo := &mocks.PondRepositoryMock{}
	farmRepo := &mocks.FarmRepositoryMock{}
	pondService := services.NewPondService(pondRepo, farmRepo)

	ctx := context.Background()

	// Define test case
	testCases := []struct {
		name          string
		userID        string
		mockResponse  []*models.Pond
		expectedError error
	}{
		{
			name:   "Get ponds by user ID successfully",
			userID: "userID",
			mockResponse: []*models.Pond{
				{
					ID:        "pondID1",
					PondName:  "Pond 1",
					FarmID:    "farmID",
					UpdatedAt: time.Now(),
				},
				{
					ID:        "pondID2",
					PondName:  "Pond 2",
					FarmID:    "farmID",
					UpdatedAt: time.Now(),
				},
			},
			expectedError: nil,
		},
		{
			name:          "No ponds found for the user",
			userID:        "userID",
			mockResponse:  nil,
			expectedError: errors.New("no ponds found for the user"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Set up mock behavior
			pondRepo.On("GetPondsByUserID", ctx, testCase.userID).Return(testCase.mockResponse, testCase.expectedError).Once()

			ponds, err := pondService.GetPondsByUserID(ctx, testCase.userID)
			assert.Equal(t, testCase.mockResponse, ponds)
			assert.Equal(t, testCase.expectedError, err)

			// Verify that the expected method is called
			pondRepo.AssertCalled(t, "GetPondsByUserID", ctx, testCase.userID)
		})
	}
}
