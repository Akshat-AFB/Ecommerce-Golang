package repositories

import (
    "backend-go/database"
    "backend-go/models"
)

func IsEmailTaken(email string) (bool, error) {
    var count int
    err := database.GetDB().QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
    return count > 0, err
}

func IsUsernameTaken(username string) (bool, error) {
    var count int
    err := database.GetDB().QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", username).Scan(&count)
    return count > 0, err
}

func CreateUser(user models.User) (models.User, error) {
	db := database.GetDB()

	tx, err := db.Begin()
	if err != nil {
		return models.User{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := `
        INSERT INTO users (username, email, password, role)
        VALUES ($1, $2, $3, $4)
        RETURNING id;
    `
	err = tx.QueryRow(query, user.Username, user.Email, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		return models.User{}, err
	}
    
	if commitErr := tx.Commit(); commitErr != nil {
		return models.User{}, commitErr
	}
    user.Password = "" // Clear password before returning
	return user, nil
}


func FindUserByLogin(login string) (models.User, error) {
    var user models.User
    row := database.GetDB().QueryRow(
        "SELECT id, username, email, password, role FROM users WHERE email = $1 OR username = $2",
        login, login,
    )
    err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
    if err != nil {
        return models.User{}, err
    }
    return user,nil
}