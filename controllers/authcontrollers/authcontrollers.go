package authcontrollers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vernandodev/go-restapi-jwt-mux/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan JSON yang diterima dari frontend

	var userInput models.User

	// ambil json
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		log.Fatal("Gagal mendecode json")
	}

	// close request body
	defer r.Body.Close()

	// hash password menggunakan bcrypt

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert to database
	if err := models.DB.Create(&userInput).Error; err != nil {
		log.Fatal("Gagal menyimpan data")
	}

	response, _ := json.Marshal(map[string]string{"message": "Success"})

	// set header
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
