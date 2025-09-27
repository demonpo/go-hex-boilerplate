package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"goHexBoilerplate/src/modules/user/application/rest/schemas"
	"goHexBoilerplate/src/modules/user/domain/services"
	"net/http"
	"time"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(UserService *services.UserService) *UserHandler {
	fmt.Println("NewUserHandler")
	fmt.Println(UserService)
	return &UserHandler{
		userService: UserService,
	}
}

func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var user schemas.CreateUserSchema
	if err := ctx.BodyParser(&user); err != nil {
		return HandleError(ctx, http.StatusBadRequest, err)
	}

	validate := validator.New()

	// Validate the User struct
	err := validate.Struct(user)
	if err != nil {
		// Validation failed, handle the error
		validationErrors := err.(validator.ValidationErrors)
		return HandleError(ctx, http.StatusBadRequest, validationErrors)
	}

	newUser, err := h.userService.Create(services.CreateInput{
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return HandleError(ctx, http.StatusBadRequest, err)
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "New user created successfully",
		"data":    newUser,
	})
}

func (h *UserHandler) ReadUser(ctx *fiber.Ctx) error {
	idString := ctx.Params("id")

	id, err := uuid.Parse(idString)
	if err != nil {
		return HandleError(ctx, http.StatusBadRequest, err)
	}

	user, err := h.userService.GetById(id)

	if err != nil {
		return HandleError(ctx, http.StatusBadRequest, err)
	}
	return ctx.Status(http.StatusOK).JSON(*user)
}

func (h *UserHandler) ReadUsers(ctx *fiber.Ctx) error {

	user, err := h.userService.GetById(uuid.UUID{})
	if err != nil {
		return HandleError(ctx, http.StatusBadRequest, err)
	}
	return ctx.Status(http.StatusOK).JSON(*user)
}

//func (h *UserHandler) UpdateUser(ctx *gin.Context) {
//	// Load API configuration
//	apiCfg, err := repository.LoadAPIConfig()
//	if err != nil {
//		HandleError(ctx, http.StatusBadRequest, err)
//		return
//	}
//
//	// Validate token
//	userID, err := ValidateToken(ctx.Request.Header.Get("Authorization"), apiCfg.JWTSecret)
//	if err != nil {
//		HandleError(ctx, http.StatusBadRequest, err)
//		return
//	}
//
//	// Update user
//	var user domain.User
//	if err := ctx.ShouldBindJSON(&user); err != nil {
//		HandleError(ctx, http.StatusBadRequest, err)
//		return
//	}
//
//	err = h.svc.UpdateUser(userID, user.Email, user.Password)
//	if err != nil {
//		HandleError(ctx, http.StatusBadRequest, err)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"message": "User updated successfully",
//	})
//}
//
//func (h *UserHandler) DeleteUser(ctx *gin.Context) {
//	apiCfg, err := repository.LoadAPIConfig()
//	if err != nil {
//		HandleError(ctx, http.StatusBadRequest, err)
//		return
//	}
//
//	userID, err := ValidateToken(ctx.Request.Header.Get("Authorization"), apiCfg.JWTSecret)
//	if err != nil {
//		HandleError(ctx, http.StatusBadRequest, err)
//		return
//	}
//
//	err = h.svc.DeleteUser(userID)
//	if err != nil {
//		HandleError(ctx, http.StatusBadRequest, err)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"message": "User deleted successfully",
//	})
//}

func ValidateToken(authHeader string, jwtSecret string) (string, error) {
	// Check if token exists in the header
	if authHeader == "" {
		return "", errors.New("token not found")
	}

	// Extract token from header
	tokenString := authHeader[7:]

	// Parse and validate token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("token not valid")
	}

	// Check if token has expired
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || claims.ExpiresAt == nil || claims.ExpiresAt.Before(time.Now().UTC()) {
		return "", errors.New("token has expired")
	}

	// Check if token is a refresh token
	if claims.Issuer == "LordMoMA-refresh" {
		return "", errors.New("token is a refresh token, please use access token")
	}

	// Extract user ID from token
	userID := claims.Subject

	return userID, nil
}
