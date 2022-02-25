package repository

import (
	"crud-postgres/src/database"
	"crud-postgres/src/models"
	"crud-postgres/src/security"
	"log"
)

func FindUsersRepository() ([]models.User, error) {
	db := database.Connect()

	var users []models.User

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		users = append(users, user)
	}

	return users, err
}

func CreateUserRepository(user models.User) {
	db := database.Connect()

	defer db.Close()

	passwordHash, _ := security.HashPassword(user.Password)

	statement := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	err := db.QueryRow(statement, user.Name, user.Email, passwordHash).Scan(&id)

	if err != nil {
		log.Fatalf("Error in execute query: %v", err)
	}
}

func DeleteUserRepository(id int) {
	db := database.Connect()

	defer db.Close()

	statement := `DELETE FROM users WHERE id=$1`

	res, err := db.Exec(statement, id)

	if err != nil {
		log.Fatalf("error to delete user: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("error while check affected rows: %v", err)
	}

	if rowsAffected > 1 {
		log.Fatalf("more of one user delected, check!")
	}
}

func UpdateUserRepository(id int, user models.User) {
	db := database.Connect()

	defer db.Close()

	statement := `UPDATE users SET name=$2, email=$3 WHERE id=$1`

	res, err := db.Exec(statement, id, user.Name, user.Email)

	if err != nil {
		log.Fatalf("error in execute query: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("error in check affected rows: %v", err)
	}

	if rowsAffected > 1 {
		log.Println("More of one affected rows, check!")
	}
}
