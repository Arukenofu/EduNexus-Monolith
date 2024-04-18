package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/utils"
)

func HandleLogin(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		login := r.PostFormValue("login")
		if is_login_valid := utils.ValidateLogin(login); is_login_valid != nil {
			types.NewJsonResponse(struct {
				Message any `json:"error_message"`
			}{is_login_valid.Error()}, http.StatusUnauthorized).Respond(w)
			return
		}
		password := r.PostFormValue("password")
		if is_password_valid := utils.ValidatePassword(password); is_password_valid != nil {
			types.NewJsonResponse(struct {
				Message any `json:"error_message"`
			}{is_password_valid.Error()}, http.StatusUnauthorized).Respond(w)
			return
		}
		hash_password, err := DB.GetPasswordByLogin(context.Background(), login)
		if err != nil {
			types.NewJsonResponse(struct {
				Message any `json:"error_message"`
			}{"Problem with database"}, http.StatusInternalServerError).Respond(w)
			return
		}
		if is_correct := utils.CheckPassword(password, hash_password); !is_correct {
			types.NewJsonResponse(struct {
				Message any `json:"error_message"`
			}{"Password or Login is wrong"}, http.StatusUnauthorized).Respond(w)
			return
		}
		// TODO: RETURN TOKEN
		claims, err := DB.GetClaimsByLogin(context.Background(), login)
		if err != nil {
			types.NewJsonResponse(struct {
				Message any `json:"error_message"`
			}{"Problem with database"}, http.StatusInternalServerError).Respond(w)
			return
		}
		token, err := json.Marshal(struct {
			Id        int64  `json:"id"`
			User_role string `json:"user_role"`
		}{Id: claims.ID, User_role: claims.Title.String})
		if err != nil {
			log.Println(err)
			return
		}
		utils.Encrypt(&token)
		types.NewJsonResponse(struct {
			Token any `json:"token"`
		}{token}, http.StatusOK).Respond(w)
	}
}