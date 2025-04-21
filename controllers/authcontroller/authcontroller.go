package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MarcellinusAditya/go-jwt/config"
	"github.com/MarcellinusAditya/go-jwt/helper"
	"github.com/MarcellinusAditya/go-jwt/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	
	defer r.Body.Close()

	// user check
	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil{
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Username atau password salah"}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": "Username atau password salah"}
			helper.ResponseJson(w, http.StatusInternalServerError, response)
			return
		}
		
	}

	// password check
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil{
		response := map[string]string{"message": "Username atau password salah"}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	// jwt token generate
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//sign in algoritm
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	// set token to cookie
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Login Berhasil"}
	helper.ResponseJson(w, http.StatusOK, response)

}
func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	
	defer r.Body.Close()

	// hash
	hashPassword,_ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password=string(hashPassword)

	// insert db
	if err := models.DB.Create(&userInput).Error; err != nil{
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"message": "success"}
	helper.ResponseJson(w, http.StatusOK, response)
	
}
func Logout(w http.ResponseWriter, r *http.Request) {

}