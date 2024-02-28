package middleware

import (
	"fmt"
	"ozanpay/config"
	"ozanpay/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JwtErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
}

func JwtSuccessHandler(c *fiber.Ctx) error {
	user := model.User{}
	tokenByte := c.Request().Header.Peek("Authorization")
	tokenStr := strings.ReplaceAll(string(c.Request().Header.Peek("Authorization")), "Bearer ", "")

	if len(tokenByte) == 0 {
		if len(c.Query("token")) == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Login Gerekli.1"})
		}
		tokenStr = c.Query("token")
	}

	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Login Gerekli.2"})
	}

	jwtToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get().Server.JwtSecret), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	if !jwtToken.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "JWT Geçersiz"})
	}
	claims := jwtToken.Claims.(jwt.MapClaims)

	userID := claims["user_id"]
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "JWT Eksikleri var. Login gerekli"})
	}
	if _, ok := userID.(float64); ok {
		user.ID = int64(userID.(float64))
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "JWT Eksikleri var. Login gerekli"})
	}
	nameSurname := claims["name_surname"]
	if nameSurname == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "JWT Eksikleri var. Login gerekli"})
	}
	user.Name = fmt.Sprintf("%v", nameSurname)

	role := claims["role"]
	if role == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "JWT Eksikleri var. Login gerekli"})
	}
	if _, ok := role.(float64); ok {
		user.Role = model.UserRole(int(role.(float64)))
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "JWT Eksikleri var. Login gerekli"})
	}

	c.Locals("user", user)
	return c.Next()
}
