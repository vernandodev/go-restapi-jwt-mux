package authcontrollers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vernandodev/go-restapi-jwt-mux/config"
	"github.com/vernandodev/go-restapi-jwt-mux/helper"
	"github.com/vernandodev/go-restapi-jwt-mux/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan JSON yang diterima dari frontend
	var userInput models.User

	// ambil json
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		// set header
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// close request body
	defer r.Body.Close()

	// ambil data user berdasarkan username
	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		// memberikan switch pada error, berkemungkinan mempunyai 2 error (not found, internal server error)
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Username salah"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// verify password yg user input dengan yang ada di database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Password salah"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTclaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-restapi-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarasikan algoritma yang akan digunakan untuk sign in
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// ketika token berhasil di create lalu simpan ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "login berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan JSON yang diterima dari frontend
	var userInput models.User

	// ambil json
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		// set header
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// close request body
	defer r.Body.Close()

	// hash password menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert to database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set header
	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	// hapus token yang ada di cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "logout berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
