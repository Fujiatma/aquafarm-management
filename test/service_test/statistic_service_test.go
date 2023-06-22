package service_test

import (
	"context"
	"github.com/stretchr/testify/assert"

	"errors"
	"github.com/delos/aquafarm-management/src/models"
	"github.com/delos/aquafarm-management/src/services"
	"github.com/delos/aquafarm-management/test/mocks"
	"testing"
	"time"
)

func TestStatisticService_CreateStatistic(t *testing.T) {
	statRepo := &mocks.StatisticRepositoryMock{}
	statService := services.NewStatisticService(statRepo)

	ctx := context.Background()

	// Define test case
	testCases := []struct {
		name          string
		statistic     *models.Statistic
		mockError     error
		expectedError error
	}{
		{
			name: "Create statistic successfully",
			statistic: &models.Statistic{
				Endpoint: "/api/v1/users",
				UserID:   "userID",
				Count:    1,
			},
			mockError:     nil,
			expectedError: nil,
		},
		{
			name: "Failed to create statistic",
			statistic: &models.Statistic{
				Endpoint: "/api/v1/users",
				UserID:   "userID",
				Count:    1,
			},
			mockError:     errors.New("failed to create statistic"),
			expectedError: errors.New("failed to create statistic"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Set up mock behavior
			statRepo.On("CreateStatistic", ctx, testCase.statistic).Return(testCase.mockError).Once()

			err := statService.CreateStatistic(ctx, testCase.statistic)
			assert.Equal(t, testCase.expectedError, err)

			// Verify that the expected method is called
			statRepo.AssertCalled(t, "CreateStatistic", ctx, testCase.statistic)
		})
	}
}

func TestStatisticService_GetAllStatistics(t *testing.T) {
	statRepo := &mocks.StatisticRepositoryMock{}
	statService := services.NewStatisticService(statRepo)

	ctx := context.Background()

	// Define test case
	testCases := []struct {
		name          string
		mockResponse  []*models.Statistic
		expectedError error
	}{
		{
			name: "Get all statistics successfully",
			mockResponse: []*models.Statistic{
				{
					ID:        "statID1",
					Endpoint:  "/api/v1/users",
					UserID:    "userID",
					Count:     10,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					ID:        "statID2",
					Endpoint:  "/api/v1/farms",
					UserID:    "userID",
					Count:     5,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			expectedError: nil,
		},
		{
			name:          "No statistics found",
			mockResponse:  nil,
			expectedError: errors.New("no statistics found"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Set up mock behavior
			statRepo.On("GetAllStatistics", ctx).Return(testCase.mockResponse, testCase.expectedError).Once()

			statistics, err := statService.GetStatistics(ctx)
			assert.Equal(t, testCase.mockResponse, statistics)
			assert.Equal(t, testCase.expectedError, err)

			// Verify that the expected method is called
			statRepo.AssertCalled(t, "GetAllStatistics", ctx)
		})
	}
}
