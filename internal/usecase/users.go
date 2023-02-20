package usecase

import (
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// UsersUseCase -.
type UsersUseCase struct {
	repo UserRepo
}

// NewUsersUseCase -.
func NewUsersUseCase(r UserRepo) *UsersUseCase {
	return &UsersUseCase{repo: r}
}

// GetCount -.
func (uc *UsersUseCase) GetCount(req request.GetUsersReq) (int, error) {
	rows, err := uc.repo.Count(req)
	if err != nil {
		return rows, fmt.Errorf("UsersUseCase - GetCount - uc.repo.Count: %w", err)
	}
	return rows, nil
}

// GetAllUsers -.
func (uc *UsersUseCase) GetAllUsers(req request.GetUsersReq) ([]entity.Users, error) {
	users, err := uc.repo.Select(req)
	if err != nil {
		return nil, fmt.Errorf("UsersUseCase - GetAllUsers - uc.repo.Select: %w", err)
	}
	return users, nil
}

// GetUserByID -.
func (uc *UsersUseCase) GetUserByID(userID int) (entity.Users, error) {
	var user entity.Users
	user, err := uc.repo.SelectById(userID)
	if err != nil {
		return user, fmt.Errorf("UsersUseCase - GetUserByID - uc.repo.SelectById: %w", err)
	}
	return user, nil
}
