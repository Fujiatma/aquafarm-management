package service_test

import (
	"context"
	"github.com/stretchr/testify/assert"

	"errors"
	"github.com/delos/aquafarm-management/src/models"
	"github.com/delos/aquafarm-management/src/services"
	"github.com/delos/aquafarm-management/test/mocks"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	userRepo := &mocks.UserRepositoryMock{}
	userService := services.NewUserService(userRepo)

	ctx := context.Background()

	// Define test case
	testCases := []struct {
		name          string
		user          *models.User
		mockResponse  error
		expectedError error
	}{
		{
			name: "Create user successfully",
			user: &models.User{
				ID:       "userID",
				UserName: "John Doe",
			},
			mockResponse:  nil,
			expectedError: nil,
		},
		{
			name: "Failed to create user",
			user: &models.User{
				ID:       "userID",
				UserName: "John Doe",
			},
			mockResponse:  errors.New("failed to create user"),
			expectedError: errors.New("failed to create user"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Set up mock behavior
			userRepo.On("CreateUser", ctx, testCase.user).Return(testCase.mockResponse).Once()

			err := userService.CreateUser(ctx, testCase.user)
			assert.Equal(t, testCase.expectedError, err)

			// Verify that the expected method is called
			userRepo.AssertCalled(t, "CreateUser", ctx, testCase.user)
		})
	}
}
