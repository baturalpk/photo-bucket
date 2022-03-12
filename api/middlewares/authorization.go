package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/baturalpk/photo-bucket/utils/crypto"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
	AuthHeaderKey       = "Authorization"
	AuthTypeIdentifier  = "Bearer"
	AuthHeaderDelimiter = AuthTypeIdentifier + " "
	AuthContextKey      = "id"

	// Errors

	ErrorMissingAuthHeaderCode    = http.StatusUnauthorized
	ErrorMissingAuthHeaderMessage = "missing authorization header"

	ErrorBadAuthHeaderCode    = http.StatusUnauthorized
	ErrorBadAuthHeaderMessage = "bad authorization header, expecting: " + AuthTypeIdentifier

	ErrorBadAuthTokenCode    = http.StatusForbidden
	ErrorBadAuthTokenMessage = "invalid, malformed and/or expired token"
)

func AuthorizationGuard(c *fiber.Ctx) error {
	authHeader := c.Get(AuthHeaderKey)
	if authHeader == "" {
		return c.Status(ErrorMissingAuthHeaderCode).SendString(ErrorMissingAuthHeaderMessage)
	}

	splitHeader := strings.Split(authHeader, AuthHeaderDelimiter)
	if len(splitHeader) != 2 {
		return c.Status(ErrorBadAuthHeaderCode).SendString(ErrorBadAuthHeaderMessage)
	}

	tokenString := strings.TrimSpace(splitHeader[1])
	claims, err := crypto.ValidateAndExtractJWTClaims(tokenString)
	if err != nil {
		return c.Status(ErrorBadAuthTokenCode).SendString(ErrorBadAuthTokenMessage)
	}

	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		return c.Status(ErrorBadAuthTokenCode).SendString(ErrorBadAuthTokenMessage)
	}

	userID, err := uuid.Parse(sub)
	if err != nil {
		return c.Status(ErrorBadAuthTokenCode).SendString(ErrorBadAuthTokenMessage)
	}

	withUserID := context.WithValue(c.UserContext(), AuthContextKey, userID)
	c.SetUserContext(withUserID)
	return c.Next()
}
