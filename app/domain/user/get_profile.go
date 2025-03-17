package user

import (
	"auth-service/types"
	"context"
	"fmt"
)

func (svc *service) GetProfile(ctx context.Context, id int) (*types.ProfileResponse, error) {
	user, err := svc.userRepo.GetUserById(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %v", err)
	}

	userResponse := user.ConvertToProfileResponse()

	return &userResponse, nil
}
