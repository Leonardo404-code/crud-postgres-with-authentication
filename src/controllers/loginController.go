package controllers

import (
	"crud-postgres/src/auth"
	"crud-postgres/src/models"
	"crud-postgres/src/repository"
	"crud-postgres/src/security"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
)

// Login verify if user exists in database and return JWT for user sign in
func Login(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	var user models.User

	err := json.NewDecoder(body).Decode(&user)

	if err != nil {
		log.Fatalf("Error in decode body: %v", err)
	}

	if len(user.Password) <= 8 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Password must contain more than 8 caracters")
		return
	}

	_, errEmail := mail.ParseAddress(user.Email)

	if errEmail != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Email invalid")
		return
	}

	userFind, err := repository.FindUserByEmail(user.Email)

	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("Failed to find user in database")
	}

	if err := security.VerifyPasswordMatch(userFind.Password, user.Password); err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("password invalid")
		return
	}

	token, err := auth.CreateToken(userFind.ID)

	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("error in create token")
		return
	}

	w.Write([]byte(token))
}
