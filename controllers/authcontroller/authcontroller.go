package authcontroller

import (
	"encoding/json"
	"invite-wed/configs"
	"invite-wed/helpers"
	"invite-wed/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

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

	var existingUser models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := map[string]string{"message": "username or password invalid"}
			helpers.JsonResponse(w, http.StatusNotFound, response)
			return
		}

		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Username atau password salah"}
		helpers.JsonResponse(w, http.StatusNotFound, response)
		return
	}

	expTime := time.Now().Add(time.Minute * 1)
	claims := &configs.JWTClaims{
		UserId: existingUser.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-fashion-shop",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(configs.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})

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

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "Successful Logout"}
	helpers.JsonResponse(w, http.StatusOK, response)
}
