package AuthMiddleware

import (
	"fmt"
	"os"
	"portalapp/Database"
	usermodel "portalapp/Model/UserModel"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func verifyToken(tokenString string) (*jwt.Token, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}

func RequireAuth(ctx *fiber.Ctx) error {
	var tokenString string
	tokenString = ctx.Get("Authorization")

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	} else {
		return ctx.Status(400).JSON(fiber.Map{"sucess": false, "message": "Token format error."})
	}

	token, err := verifyToken(tokenString)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"sucess": false, "message": "Token error."})
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"sucess": false, "message": "Token error."})
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"sucess": false, "message": "Token expired."})
	}

	var user usermodel.User
	user_id := claims["sub"]
	Database.DB.First(&user, user_id)
	if user.ID == 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"sucess": false, "message": "Unauthorized."})
	}
	return ctx.Next()
}
