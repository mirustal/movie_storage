package middleware

import (
	"fmt"
	"movie_storage/internal/models"
	"movie_storage/pkg/utils"
	"net/http"
)


type RequestAuthDTO struct {
	Guid     string `json:"guid"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("accesst")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var body models.LoginRequest
		err = utils.DecodeJSONRequest(w, r, &body);
		if err != nil {
			fmt.Printf("error bad json body: %v\n", err)
			http.Error(w, fmt.Sprintf("error bad json body: %v", err), http.StatusBadRequest)
			return
		}

		claims, errParse := utils.ParseAccessToken(cookie.Value)
		if errParse != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !utils.ValidatePayloadToken(body, claims) { 
			http.Error(w, "Invalid token structure", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}


func AuthOnlyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("accesst")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, errParse := utils.ParseAccessToken(cookie.Value)
		if errParse != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

