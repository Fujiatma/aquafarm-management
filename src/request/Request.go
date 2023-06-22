package request

type UpsertFarmRequest struct {
	FarmName string `json:"farm_name" validate:"required"`
	UserID   string `json:"user_id" validate:"required"`
}

type UpsertPondRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	PondName string `json:"pond_name" validate:"required"`
	FarmID   string `json:"farm_id" validate:"required"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
}

type DeleteFarmRequest struct {
	FarmID string `json:"farm_id" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}

type DeletePondRequest struct {
	PondID string `json:"pond_id" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}

type GetFarmsAndPondsRequest struct {
	UserID string `json:"user_id" validate:"required"`
}
