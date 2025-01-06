package middleware

import (
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/controllers"
)

func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next(w, r)
	}
}
