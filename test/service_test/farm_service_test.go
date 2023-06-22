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

func TestFarmService_CreateFarm(t *testing.T) {
	farmRepo := &mocks.FarmRepositoryMock{}
	farmService := services.NewFarmService(farmRepo)

	ctx := context.Background()

	// Define test case
	testCases := []struct {
		name          string
		farm          *models.Farm
		expectedError error
	}{
		{
			name: "Create farm successfully",
			farm: &models.Farm{
				ID:       "farmID",
				FarmName: "Farm 1",
				UserID:   "userID",
			},
			expectedError: nil,
		},
		{
			name: "Farm with the same name already exists",
			farm: &models.Farm{
				ID:       "farmID",
				FarmName: "Farm 1",
				UserID:   "userID",
			},
			expectedError: errors.New("farm with the same name already exists"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Set up mock behavior
			if testCase.expectedError == nil {
				farmRepo.On("GetFarmByName", ctx, testCase.farm.FarmName).Return(nil, gorm.ErrRecordNotFound).Once()
			} else {
				farmRepo.On("GetFarmByName", ctx, testCase.farm.FarmName).Return(&models.Farm{}, testCase.expectedError).Once()
			}
			farmRepo.On("CreateFarm", ctx, testCase.farm).Return(nil).Once().Run(func(args mock.Arguments) {
				farmArg := args.Get(1).(*models.Farm)
				farmArg.ID = "generatedID" // Assign a generated ID to the farm argument
			})

			err := farmService.CreateFarm(ctx, testCase.farm)
			assert.Equal(t, testCase.expectedError, err)

			// Verify that the expected methods are called
			farmRepo.AssertCalled(t, "GetFarmByName", ctx, testCase.farm.FarmName)
			farmRepo.AssertCalled(t, "CreateFarm", ctx, mock.AnythingOfType("*models.Farm"))

		})
	}
}

func TestFarmService_UpdateFarm(t *testing.T) {
	farmRepo := &mocks.FarmRepositoryMock{}
	farmService := services.NewFarmService(farmRepo)

	ctx := context.Background()

	// Define test case
	testCases := []struct {
		name          string
		farmID        string
		farmData      *models.Farm
		expectedFarm  *models.Farm
		expectedError error
	}{
		{
			name:   "Farm not found, create new farm",
			farmID: "nonExistentFarmID",
			farmData: &models.Farm{
				ID:        "farmID",
				FarmName:  "Farm 1",
				UserID:    "userID",
				UpdatedAt: time.Now(),
				CreatedAt: time.Now(),
			},
			expectedFarm: &models.Farm{
				ID:        "generatedID",
				FarmName:  "Farm 1",
				UserID:    "userID",
				UpdatedAt: time.Now(),
				CreatedAt: time.Now(),
			},
			expectedError: nil,
		},
		{
			name:   "Farm found, update existing farm",
			farmID: "existingFarmID",
			farmData: &models.Farm{
				ID:        "farmID",
				FarmName:  "Updated Farm",
				UserID:    "userID",
				UpdatedAt: time.Now(),
				CreatedAt: time.Now(),
			},
			expectedFarm: &models.Farm{
				ID:        "existingFarmID",
				FarmName:  "Updated Farm",
				UserID:    "userID",
				UpdatedAt: time.Now(),
				CreatedAt: time.Now(),
			},
			expectedError: nil,
		},
		// Add more test cases...
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Set up mock behavior
			farmRepo.On("GetFarmByID", ctx, testCase.farmData.ID).Return(testCase.expectedFarm, nil).Once()

			// If the farm does not exist, create a new farm
			if testCase.expectedFarm == nil {
				farmRepo.On("CreateFarm", ctx, testCase.farmData).Return(nil).Once().Run(func(args mock.Arguments) {
					farmArg := args.Get(1).(*models.Farm)
					farmArg.ID = "generatedID" // Assign a generated ID to the farm argument
				})
			} else {
				// If the farm exists, update the farm
				farmRepo.On("UpdateFarm", ctx, mock.AnythingOfType("*models.Farm")).Return(nil).Once()
			}

			_, err := farmService.UpdateFarm(ctx, testCase.farmData)
			assert.Equal(t, testCase.expectedError, err)

			// Verify that the expected methods are called
			farmRepo.AssertCalled(t, "GetFarmByID", ctx, testCase.farmData.ID)
			if testCase.expectedFarm == nil {
				farmRepo.AssertCalled(t, "CreateFarm", ctx, mock.AnythingOfType("*models.Farm"))
			} else {
				farmRepo.AssertCalled(t, "UpdateFarm", ctx, mock.AnythingOfType("*models.Farm"))
			}
		})
	}
}

func TestFarmService_DeleteFarm(t *testing.T) {
	farmRepo := &mocks.FarmRepositoryMock{}
	farmService := services.NewFarmService(farmRepo)

	ctx := context.Background()

	// Define test case
	testCases := []struct {
		name          string
		payload       *models.Farm
		expectedError error
	}{
		{
			name: "Delete farm successfully",
			payload: &models.Farm{
				ID:     "farmID",
				UserID: "userID",
			},
			expectedError: nil,
		},
		{
			name: "User does not have permission to delete the farm",
			payload: &models.Farm{
				ID:     "farmID",
				UserID: "otherUserID",
			},
			expectedError: errors.New("user does not have permission to delete the farm"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Set up mock behavior
			farmRepo.On("GetFarmByID", ctx, testCase.payload.ID).Return(&models.Farm{
				ID:     testCase.payload.ID,
				UserID: testCase.payload.UserID,
			}, nil).Once()

			if testCase.expectedError == nil {
				farmRepo.On("SoftDeleteFarm", ctx, mock.AnythingOfType("*models.Farm")).Return(nil).Once()
			}

			err := farmService.DeleteFarm(ctx, testCase.payload)
			assert.Equal(t, testCase.expectedError, err)

			// Verify that the expected methods are called
			farmRepo.AssertCalled(t, "GetFarmByID", ctx, testCase.payload.ID)

			if testCase.expectedError == nil {
				farmRepo.AssertCalled(t, "SoftDeleteFarm", ctx, mock.AnythingOfType("*models.Farm"))
			}
		})
	}
}

func TestFarmService_GetFarmByID(t *testing.T) {
	farmRepo := &mocks.FarmRepositoryMock{}
	farmService := services.NewFarmService(farmRepo)

	ctx := context.Background()

	// Define test case
	testCases := []struct {
		name          string
		farmID        string
		expectedFarm  *models.Farm
		expectedError error
	}{
		{
			name:   "Get farm by ID successfully",
			farmID: "farmID",
			expectedFarm: &models.Farm{
				ID:       "farmID",
				FarmName: "Farm 1",
				UserID:   "userID",
			},
			expectedError: nil,
		},
		{
			name:          "Farm not found",
			farmID:        "nonExistentFarmID",
			expectedFarm:  nil,
			expectedError: gorm.ErrRecordNotFound,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Set up mock behavior
			farmRepo.On("GetFarmByID", ctx, testCase.farmID).Return(testCase.expectedFarm, testCase.expectedError).Once()

			farm, err := farmService.GetFarmByID(ctx, testCase.farmID)
			assert.Equal(t, testCase.expectedFarm, farm)
			assert.Equal(t, testCase.expectedError, err)

			// Verify that the expected method is called
			farmRepo.AssertCalled(t, "GetFarmByID", ctx, testCase.farmID)
		})
	}
}

func TestFarmService_GetAllFarmByUserID(t *testing.T) {
	farmRepo := &mocks.FarmRepositoryMock{}
	farmService := services.NewFarmService(farmRepo)

	ctx := context.Background()

	// Define test case
	testCases := []struct {
		name          string
		userID        string
		expectedFarms []*models.Farm
		expectedError error
	}{
		{
			name:   "Get all farms by user ID successfully",
			userID: "userID",
			expectedFarms: []*models.Farm{
				{
					ID:       "farmID1",
					FarmName: "Farm 1",
					UserID:   "userID",
				},
				{
					ID:       "farmID2",
					FarmName: "Farm 2",
					UserID:   "userID",
				},
			},
			expectedError: nil,
		},
		{
			name:          "No farms found for the user",
			userID:        "nonExistentUserID",
			expectedFarms: nil,
			expectedError: gorm.ErrRecordNotFound,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Set up mock behavior
			farmRepo.On("GetFarmsByUserID", ctx, testCase.userID).Return(testCase.expectedFarms, testCase.expectedError).Once()

			farms, err := farmService.GetAllFarmByUserID(ctx, testCase.userID)
			assert.Equal(t, testCase.expectedFarms, farms)
			assert.Equal(t, testCase.expectedError, err)

			// Verify that the expected method is called
			farmRepo.AssertCalled(t, "GetFarmsByUserID", ctx, testCase.userID)
		})
	}
}

func TestFarmService_GetFarmByName(t *testing.T) {
	farmRepo := &mocks.FarmRepositoryMock{}
	farmService := services.NewFarmService(farmRepo)

	ctx := context.Background()

	// Define test case
	testCases := []struct {
		name          string
		farmName      string
		expectedFarm  *models.Farm
		expectedError error
	}{
		{
			name:     "Get farm by name successfully",
			farmName: "Farm 1",
			expectedFarm: &models.Farm{
				ID:       "farmID",
				FarmName: "Farm 1",
				UserID:   "userID",
			},
			expectedError: nil,
		},
		{
			name:          "Farm not found",
			farmName:      "Nonexistent Farm",
			expectedFarm:  nil,
			expectedError: gorm.ErrRecordNotFound,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Set up mock behavior
			farmRepo.On("GetFarmByName", ctx, testCase.farmName).Return(testCase.expectedFarm, testCase.expectedError).Once()

			farm, err := farmService.GetFarmByName(ctx, testCase.farmName)
			assert.Equal(t, testCase.expectedFarm, farm)
			assert.Equal(t, testCase.expectedError, err)

			// Verify that the expected method is called
			farmRepo.AssertCalled(t, "GetFarmByName", ctx, testCase.farmName)
		})
	}
}
