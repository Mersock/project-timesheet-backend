package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
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

// UpdateUser -.
func (uc *UsersUseCase) UpdateUser(req request.UpdateUserReq) (int64, error) {
	var rowAffected int64

	_, err := uc.repo.SelectById(req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return rowAffected, sql.ErrNoRows
		}
		return rowAffected, fmt.Errorf("UsersUseCase - UpdateUser - uc.repo.SelectById: %w", err)
	}

	count, err := uc.repo.ChkUniqueUpdate(req)
	if err != nil {
		return rowAffected, fmt.Errorf("UsersUseCase - UpdateUser - uc.repo.ChkUniqueUpdate: %w", err)
	}

	if count > 0 {
		return rowAffected, ErrDuplicateRow
	}

	rowAffected, err = uc.repo.Update(req)
	if err != nil {
		return rowAffected, fmt.Errorf("UsersUseCase - UpdateUser - uc.repo.Update: %w", err)
	}
	return rowAffected, nil
}

// UpdateUserPassword -.
func (uc *UsersUseCase) UpdateUserPassword(req request.UpdateUserPasswordReq) (int64, error) {
	var rowAffected int64

	_, err := uc.repo.SelectById(req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return rowAffected, sql.ErrNoRows
		}
		return rowAffected, fmt.Errorf("UsersUseCase - UpdateUserPassword - uc.repo.SelectById: %w", err)
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return rowAffected, fmt.Errorf("AuthUseCase - Signup - utils.HashPassword: %w", err)
	}

	req.Password = hashPassword

	rowAffected, err = uc.repo.UpdatePassword(req)
	if err != nil {
		return rowAffected, fmt.Errorf("UsersUseCase - UpdateUserPassword - uc.repo.Update: %w", err)
	}
	return rowAffected, nil
}

// DeleteUser -.
func (uc *UsersUseCase) DeleteUser(req request.DeleteUserReq) (int64, error) {
	var rowAffected int64

	_, err := uc.repo.SelectById(req.ID)
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
