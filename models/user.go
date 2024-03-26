package models

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

const UserSchema string = `CREATE TABLE IF NOT EXISTS users(
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50) NOT NULL,
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`