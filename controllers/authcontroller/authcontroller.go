package authcontroller

import (
	"encoding/json"
	"invite-wed/helpers"
	"invite-wed/models"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

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

	response := map[string]interface{}{"data": userInput}
	helpers.JsonResponse(w, http.StatusOK, response)
}
