package config

import "os"

type AuthConfig struct {
    JWTSecret string
}

func NewAuthConfig() *AuthConfig {
    return &AuthConfig{
        JWTSecret: os.Getenv("JWT_SECRET"),
    }
}