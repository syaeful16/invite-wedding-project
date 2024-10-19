package invitationcontroller

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"invite-wed/helpers"
	"invite-wed/middlewares"
	"invite-wed/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.IdKey).(uint)

	var invitation []models.Invitation
	if err := models.DB.Where("user_id = ?", userID).Find(&invitation).Error; err != nil {
		response := map[string]string{"message": "Error fetching data"}
		helpers.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	if len(invitation) == 0 {
		response := map[string]string{"message": "Belum ada Undangan yang dibuat"}
		helpers.JsonResponse(w, http.StatusNotFound, response)
		return
	}

	helpers.JsonResponse(w, http.StatusOK, invitation)
}

func Store(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.IdKey).(uint)

	var invitation models.Invitation
	decode := json.NewDecoder(r.Body)

	bytes := make([]byte, 4)
	_, err := rand.Read(bytes)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	uniqueCodeInvitation := base64.StdEncoding.EncodeToString(bytes)

	invitation.InvitationCode = uniqueCodeInvitation
	invitation.UserID = userID

	if err := decode.Decode(&invitation); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if validate := helpers.Validation(invitation); validate != nil {
		response := map[string]interface{}{"error": validate}
		helpers.JsonResponse(w, http.StatusBadRequest, response)
		return
	}

	if err := models.DB.Create(&invitation).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{"message": "Undangan berhasil dibuat", "data": invitation}
	helpers.JsonResponse(w, http.StatusCreated, response)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusBadRequest, response)
		return
	}

	var invitation models.Invitation
	if err := models.DB.Where("id = ?", id).First(&invitation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := map[string]string{"message": err.Error()}
			helpers.JsonResponse(w, http.StatusNotFound, response)
			return
		}
		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	helpers.JsonResponse(w, http.StatusOK, invitation)
}

func Update(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.IdKey).(uint)

	var invitationInput models.Invitation
	decode := json.NewDecoder(r.Body)

	if err := decode.Decode(&invitationInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if validate := helpers.Validation(invitationInput); validate != nil {
		response := map[string]interface{}{"error": validate}
		helpers.JsonResponse(w, http.StatusBadRequest, response)
		return
	}

	var existingInvitation models.Invitation
	if err := models.DB.Where("id = ?", invitationInput.ID).Where("user_id = ?", userID).First(&existingInvitation).Error; err != nil {
		response := map[string]string{"message": "Data tidak ditemukan"}
		helpers.JsonResponse(w, http.StatusNotFound, response)
		return
	}

	if err := models.DB.Model(&existingInvitation).Updates(&invitationInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{"message": "Data berhasil diupdate", "data": existingInvitation}
	helpers.JsonResponse(w, http.StatusOK, response)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.IdKey).(uint)

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusBadRequest, response)
		return
	}

	var invitation models.Invitation
	if err := models.DB.Where("id = ?", id).Where("user_id = ?", userID).First(&invitation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := map[string]string{"message": "Data tidak ditemukan"}
			helpers.JsonResponse(w, http.StatusNotFound, response)
			return
		}
		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	if err := models.DB.Delete(&invitation).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Data berhasil dihapus"}
	helpers.JsonResponse(w, http.StatusOK, response)
}
