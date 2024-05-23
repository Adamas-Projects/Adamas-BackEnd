package webserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity/reqs"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/go-chi/jwtauth"
)

type WebInstitutionHandler struct {
	institutionService *service.InstitutionService
}

func NewWebInstiHandler(institutionService *service.InstitutionService) *WebInstitutionHandler {
	return &WebInstitutionHandler{
		institutionService: institutionService,
	}
}

func (wih *WebInstitutionHandler) CreateInstitution(w http.ResponseWriter, r *http.Request, tokenAuth *jwtauth.JWTAuth) {
	var institution *reqs.InstitutionCreateRequest
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&institution)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	result, err := wih.institutionService.CreateInstitution(institution.Name, institution.Email, institution.Password, institution.CNPJ)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	} else {
		claims := map[string]interface{}{"id": result.USER.ID, "name": result.USER.Name, "email": result.USER.Email, "exp": jwtauth.ExpireIn(time.Minute * 10)}
		_, tokenString, _ = tokenAuth.Encode(claims)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": tokenString,
		})
	}

}

func (wih *WebInstitutionHandler) LoginInstitution(w http.ResponseWriter, r *http.Request, tokenAuth *jwtauth.JWTAuth) {
	var login *reqs.LoginRequest
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	result, err := wih.institutionService.LoginInstitution(login.Email, login.Password)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	} else {
		claims := map[string]interface{}{"id": result.USER.ID, "name": result.USER.Name, "email": result.USER.Email, "user_type": result.USER.UserType, "exp": jwtauth.ExpireIn(time.Minute * 10)}
		_, tokenString, _ = tokenAuth.Encode(claims)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": tokenString,
		})
	}
}
