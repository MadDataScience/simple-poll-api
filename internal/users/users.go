package users

import (
	"database/sql"
	database "github.com/maddatascience/simple-poll-api/internal/pkg/db/sqlite"
	"golang.org/x/crypto/bcrypt"

	"log"
)


type User struct {
	UserID string `json:"userID"`
	Email string `json:"email"`
	Password string `json:"password"`
}


func (user *User) Create() {
	statement, err := database.Db.Prepare("INSERT INTO Users(Email, Password) VALUES(?, ?)")
	print(statement)
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := HashPassword(user.Password)
	_, err = statement.Exec(user.Email, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//GetUserIdByEmail check if a user exists in database by given email address
func GetUserIdByEmail(email string) (int, error) {
	statement, err := database.Db.Prepare("select UserID from Users WHERE Email = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(email)

	var userID int
	err = row.Scan(&userID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return userID, nil
}

func (user *User) Authenticate() bool {
	statement, err := database.Db.Prepare("select Password from Users WHERE Email = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(user.Email)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
