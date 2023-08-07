package helper

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"backend/main/internal/models"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordAgainstHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func CheckCredentials(ctx context.Context, store models.DataStore, user *models.User) bool {
	fetchedUser, err := store.GetUser(ctx, user.Username)
	if err != nil {
		return false
	}

	return CheckPasswordAgainstHash(user.Password, fetchedUser.Password)
}
