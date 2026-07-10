// Middleware to check for jwt auth
package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Aadi022/Backend_Golang/internal/response"
	"github.com/Aadi022/Backend_Golang/internal/service"
)

type contextKey string // Custom type to avoid context key collisions

const UserIDKey contextKey = "user_id" // Key used to store the logged-in user's ID

func Auth(jwt *service.JWTService) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization") // Read the Authorization header

			if authHeader == "" {
				response.Error(w, http.StatusUnauthorized, "missing authorization header")
				return
			}

			if !strings.HasPrefix(authHeader, "Bearer ") { // Ensure the token uses the Bearer scheme
				response.Error(w, http.StatusUnauthorized, "invalid authorization header")
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ") // Remove "Bearer " from the header

			userID, err := jwt.Validate(token) // Validate the JWT and extract the user ID

			if err != nil {
				response.Error(w, http.StatusUnauthorized, "invalid or expired token")
				return
			}

			ctx := context.WithValue(
				r.Context(), // Existing request context
				UserIDKey,   // Key
				userID,      // Value
			)

			next.ServeHTTP(
				w,
				r.WithContext(ctx), // Pass the updated request with the user ID
			)
		})
	}
}

/*
//Entire flow of middleware and why it returns a middleware and http handler

Auth( Requires and loads jwtService)
      ↓
Stores the JWT service
      ↓
Returns the Auth middleware
      ↓
Middleware receives the handler (e.g. user.Profile)
      ↓
Returns a new handler with JWT authentication added
*/
