package auth

import (
	"github.com/gin-gonic/gin"
	"parma/pkg/models"
)

type UseCase interface {
	SignIn(ctx gin.Context, username, password string) (string, error)
	SignUp(ctx gin.Context, username, password, email string) error
	ParseToken(ctx gin.Context, accessToken string) (*models.User, error)
}
