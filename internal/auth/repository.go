package auth

import (
	"github.com/gin-gonic/gin"
	"parma/pkg/models"
)

type UserRepository interface {
	CreateUser(ctx gin.Context, user *models.User) error
	GetUser(ctx gin.Context, username, password string) (*models.User, error)
}
