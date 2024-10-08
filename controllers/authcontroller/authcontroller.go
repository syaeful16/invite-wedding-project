package authcontroller

import (
	"encoding/json"
	"invite-wed/helpers"
	"invite-wed/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// add validator message response
	if validate := helpers.Validation(userInput); validate != nil {
		response := map[string]interface{}{"error": validate}
		helpers.JsonResponse(w, http.StatusBadRequest, response)
		return
	}

	var existinguser models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&existinguser).Error; err != nil {
		response := map[string]string{"message": "Username atau password salah"}
		helpers.JsonResponse(w, http.StatusNotFound, response)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existinguser.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Username atau password salah"}
		helpers.JsonResponse(w, http.StatusNotFound, response)
		return
	}

	response := map[string]string{"message": "Berhasil login"}
	helpers.JsonResponse(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if validate := helpers.Validation(userInput); validate != nil {
		response := map[string]interface{}{"errors": validate}
		helpers.JsonResponse(w, http.StatusBadRequest, response)
		return
	}

	var existinguser models.User
	if err := models.DB.Where("username = ?", userInput.Username).Where("email = ?", userInput.Email).First(&existinguser).Error; err == nil {
		response := map[string]string{"message": "Username atau email sudah terdaftar"}
		helpers.JsonResponse(w, http.StatusBadRequest, response)
		return
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashedPassword)

	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": "Gagal register"}
		helpers.JsonResponse(w, http.StatusBadRequest, response)
		return
	}

	response := map[string]string{"message": "Berhasil register"}
	helpers.JsonResponse(w, http.StatusCreated, response)
}
