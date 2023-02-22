package middleware

import (
	"errors"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/response"
	"github.com/Mersock/project-timesheet-backend/pkg/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey    = "authorization"
	authorizationHeaderBearer = "bearer"
	authorizationUserID       = "user_id"
	authorizationRole         = "role"
)

// AuthMiddleware -.
func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is null")
			response.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header is null")
			response.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationHeaderBearer {
			err := fmt.Errorf("authorization type not support")
			response.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			response.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}
		ctx.Set(authorizationUserID, payload.UserID)
		ctx.Set(authorizationRole, payload.Role)

		ctx.Next()
	}
}
