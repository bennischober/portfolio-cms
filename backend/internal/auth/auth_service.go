package auth

import (
	"context"
	"fmt"
	"time"

	"backend/main/internal/auth/helper"
	"backend/main/internal/config"
	"backend/main/internal/models"
)

type IAuthService interface {
	RegisterUser(ctx context.Context, user *models.User) error
	LoginUser(ctx context.Context, user *models.User) (string, error)
}

type AuthService struct {
	Config       *config.Config
	Store        models.DataStore
	TokenService ITokenService
	RedisService IRedisService
}

func NewAuthService(config *config.Config,
	store models.DataStore,
	tokenService ITokenService,
	redisService IRedisService) *AuthService {
	return &AuthService{Config: config, Store: store, TokenService: tokenService, RedisService: redisService}
}

func (a *AuthService) RegisterUser(ctx context.Context, user *models.User) error {
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := a.Store.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (a *AuthService) LoginUser(ctx context.Context, user *models.User) (string, error) {
	if authenticated := helper.CheckCredentials(ctx, a.Store, user); !authenticated {
		return "", fmt.Errorf("invalid credentials")
	}

	tokenDuration := time.Duration(a.Config.Authentication.TokenExpire) * time.Minute

	// generate token
	token, err := a.TokenService.GenerateToken(user.Username, tokenDuration)
	if err != nil {
		return "", err
	}

	err = a.RedisService.SetToken(user.Username, token, tokenDuration)
	if err != nil {
		return "", err
	}

	return token, nil
}
