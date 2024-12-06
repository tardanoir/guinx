package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt"
    "golang.org/x/crypto/bcrypt"
    "guinx/internal/config"
    "guinx/internal/models"
    "gorm.io/gorm"
)

type AuthHandler struct {
    db     *gorm.DB
    config *config.AuthConfig
}

func NewAuthHandler(db *gorm.DB, config *config.AuthConfig) *AuthHandler {
    return &AuthHandler{
        db:     db,
        config: config,
    }
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
    Name     string `json:"name" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := h.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := h.generateJWT(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Register(c *gin.Context) {
    var req RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    user := models.User{
        Email:    req.Email,
        Password: string(hashedPassword),
        Name:     req.Name,
    }

    if err := h.db.Create(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
        return
    }

    token, err := h.generateJWT(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) generateJWT(user models.User) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": user.Email,
        "exp":   time.Now().Add(time.Hour * 24).Unix(),
    })

    return token.SignedString([]byte(h.config.JWTSecret))
}