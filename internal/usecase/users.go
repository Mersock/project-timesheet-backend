package usecase

import (
	"database/sql"
	"errors"
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
func (uc *UsersUseCase) GetCount(req request.GetAllUsersReq) (int, error) {
	rows, err := uc.repo.Count(req)
	if err != nil {
		return rows, fmt.Errorf("UsersUseCase - GetCount - uc.repo.Count: %w", err)
	}
	return rows, nil
}

// GetAllUsers -.
func (uc *UsersUseCase) GetAllUsers(req request.GetAllUsersReq) ([]entity.Users, error) {
	users, err := uc.repo.SelectAll(req)
	if err != nil {
		return nil, fmt.Errorf("UsersUseCase - GetAllUsers - uc.repo.Select: %w", err)
	}
	return users, nil
}

// GetUser -.
func (uc *UsersUseCase) GetUser(req request.GetUserReq) (entity.Users, error) {
	var user entity.Users
	user, err := uc.repo.SelectUser(req)
	if err != nil {
		return user, fmt.Errorf("UsersUseCase - GetUserByID - uc.repo.SelectById: %w", err)
	}
	return user, nil
}

// DeleteUser -.
func (uc *UsersUseCase) DeleteUser(req request.DeleteUserReq) (int64, error) {
	var rowAffected int64

	reqUser := request.GetUserReq{
		ID: req.ID,
	}

	_, err := uc.repo.SelectUser(reqUser)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return rowAffected, sql.ErrNoRows
		}
		return rowAffected, fmt.Errorf("UsersUseCase - DeleteRole - uc.repo.SelectById: %w", err)
	}

	rowAffected, err = uc.repo.Delete(req)
	if err != nil {
		return rowAffected, fmt.Errorf("UsersUseCase - DeleteRole - uc.repo.Delete: %w", err)
	}

	return rowAffected, nil
}
