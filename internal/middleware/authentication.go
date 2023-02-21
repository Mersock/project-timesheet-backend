package middleware

import (
	"errors"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey    = "authorization"
	authorizationHeaderBearer = "bearer"
	authorizationPayloadKey   = "authorization_payload"
)

func authMiddleware(tokenMaker jwt.JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is null")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errRes(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header is null")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errRes(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationHeaderBearer {
			err := fmt.Errorf("authorization type not support")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errRes(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errRes(err))
			return
		}
		ctx.Set(authorizationPayloadKey, payload)

		ctx.Next()
	}
}

func errRes(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
