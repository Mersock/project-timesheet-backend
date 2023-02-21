package usecase

import (
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/request"
	"github.com/Mersock/project-timesheet-backend/internal/response"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
)

const (
	_defaultRoleID = 2
)

// AuthUseCase -.
type AuthUseCase struct {
	repo UserRepo
}

// NewAuthUseCase -.
func NewAuthUseCase(u UserRepo) *AuthUseCase {
	return &AuthUseCase{repo: u}
}

func (au *AuthUseCase) Signup(req request.SignUpReq) (int64, error) {
	var userID int64

	reqCreateUser := request.CreateUserReq{
		Email:     req.Email,
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		RoleID:    _defaultRoleID,
	}

	count, err := au.repo.ChkUniqueInsert(reqCreateUser)
	if err != nil {
		return userID, fmt.Errorf("AuthUseCase - Signup - uc.repo.ChkUniqueInsert: %w", err)
	}

	if count > 0 {
		return userID, ErrDuplicateRow
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return userID, fmt.Errorf("AuthUseCase - Signup - utils.HashPassword: %w", err)
	}
	reqCreateUser.Password = hashPassword

	userID, err = au.repo.Insert(reqCreateUser)
	if err != nil {
		return userID, fmt.Errorf("AuthUseCase - Signup - uc.repo.Insert: %w", err)
	}
	return userID, nil
}

// SignIn -.
func (au *AuthUseCase) SignIn(req request.SignInReq) (response.SignInRes, error) {
	var res response.SignInRes

	return res, nil
}
