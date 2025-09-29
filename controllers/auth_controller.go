package controllers

import (
	"encoding/json"
	"net/http"
	"jwt-auth-system/models"
	"jwt-auth-system/utils"
	"gorm.io/gorm"
)

type AuthInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var input AuthInput
	json.NewDecoder(r.Body).Decode(&input)

	hashedPwd, _ := utils.HashPassword(input.Password)
	user := models.User{Email: input.Email, Password: hashedPwd}

	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "Email already in use", http.StatusBadRequest)
		return
	}
	w.Write([]byte("User registered"))
}

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var input AuthInput
	json.NewDecoder(r.Body).Decode(&input)

	var user models.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPassword(user.Password, input.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateJWT(user.ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
