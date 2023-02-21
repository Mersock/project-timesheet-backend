package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/config"
	"github.com/Mersock/project-timesheet-backend/internal/request"
	"github.com/Mersock/project-timesheet-backend/internal/response"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
	"github.com/Mersock/project-timesheet-backend/pkg/token"
)

const (
	_defaultRoleID = 2
)

// AuthUseCase -.
type AuthUseCase struct {
	repo       UserRepo
	tokenMaker token.Maker
	cfg        *config.Config
}

// NewAuthUseCase -.
func NewAuthUseCase(u UserRepo, tokenMaker token.Maker, cfg *config.Config) *AuthUseCase {
	return &AuthUseCase{
		repo:       u,
		tokenMaker: tokenMaker,
		cfg:        cfg,
	}
}

func (au *AuthUseCase) Signup(req request.SignUpReq) (int64, error) {
	var userID int64

	reqCreateUser := request.CreateUserReq{
		Email:     req.Email,
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
	var session response.SignInRes

	user, err := au.repo.SelectByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return session, sql.ErrNoRows
		}
		return session, fmt.Errorf("AuthUseCase - SignIn - uc.repo.SelectByEmail: %w", err)
	}

	err = utils.CheckPassword(req.Password, *user.Password)
	if err != nil {
		return session, fmt.Errorf("AuthUseCase - SignIn - uc.repo.CheckPassword: %w", err)
	}
	accessToken, accessTokenPayload, err := au.tokenMaker.CreateToken(*user.Email, au.cfg.AccessTokenDuration)
	if err != nil {
		return session, fmt.Errorf("AuthUseCase - SignIn - au.tokenMaker.CreateToken - accessToken: %w", err)
	}

	refreshToken, refreshTokenPayload, err := au.tokenMaker.CreateToken(*user.Email, au.cfg.RefreshTokenDuration)
	if err != nil {
		return session, fmt.Errorf("AuthUseCase - SignIn - au.tokenMaker.CreateToken- refreshToken: %w", err)
	}

	session = response.SignInRes{
		AccessToken:          accessToken,
		AccessTokenExpireAt:  accessTokenPayload.ExpireAt,
		RefreshToken:         refreshToken,
		RefreshTokenExpireAt: refreshTokenPayload.ExpireAt,
	}

	return session, nil
}

// RenewAccess -.
func (au *AuthUseCase) RenewAccess(req request.RenewTokenReq) (response.RenewTokenRes, error) {
	var session response.RenewTokenRes

	refreshPayload, err := au.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return session, fmt.Errorf("AuthUseCase - RenewAccess - au.tokenMaker.VerifyToken: %w", err)
	}

	accessToken, accessTokenPayload, err := au.tokenMaker.CreateToken(refreshPayload.Username, au.cfg.AccessTokenDuration)
	if err != nil {
		return session, fmt.Errorf("AuthUseCase - RenewAccess - au.tokenMaker.CreateToken: %w", err)
	}

	session = response.RenewTokenRes{
		AccessToken:         accessToken,
		AccessTokenExpireAt: accessTokenPayload.ExpireAt,
	}

	return session, nil
}
