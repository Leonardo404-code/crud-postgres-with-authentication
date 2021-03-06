package controllers

import (
	"crud-postgres/src/auth"
	"crud-postgres/src/models"
	"crud-postgres/src/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strconv"
)

// GetUsers return all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get all the users in the db
	users, err := repository.FindUsersRepository()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(users)
}

// GetUser return user logged
func GetUser(w http.ResponseWriter, r *http.Request) {
	token, err := auth.ExtractUserID(r)

	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	userInDatabase, err := repository.FindUserById(token)

	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(userInDatabase)
}

// CreateUser Create a single user to database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	var user models.User

	err := json.NewDecoder(body).Decode(&user)

	if err != nil {
		log.Fatalf("Error in decode body: %v", err)
	}

	if user.Name == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("field name cannot be empty")
		return
	}

	if user.Email == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("field email cannot be empty")
		return
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

	repository.CreateUserRepository(user)

	json.NewEncoder(w).Encode(user)
}

// UpdatesUser update user database informations
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	query := r.URL.Query()

	r.URL.RawQuery = query.Encode()

	id, err := strconv.Atoi(query.Get("id"))

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(fmt.Sprintf("error in convert string to int: %v", err))
	}

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("error in decode body: %v", err)
	}

	if user.Name == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("field name cannot be empty")
		return
	}

	if user.Email == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("field email cannot be empty")
		return
	}

	_, errEmail := mail.ParseAddress(user.Email)

	if errEmail != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Email invalid")
		return
	}

	repository.UpdateUserRepository(id, user)

	json.NewEncoder(w).Encode(user)
}

// DeleteUser delete user from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	r.URL.RawQuery = query.Encode()

	id, err := strconv.Atoi(query.Get("id"))

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(fmt.Sprintf("error in convert string to int: %v", err))
	}

	repository.DeleteUserRepository(id)

	json.NewEncoder(w).Encode("User delected")
}
