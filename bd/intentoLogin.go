package bd

import (
	"api/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {

	usuario, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usuario, false
	}

	passwordBytes := []byte(password)
	paswwordBD := []byte(usuario.Password)

	err := bcrypt.CompareHashAndPassword(paswwordBD, passwordBytes)

	if err != nil {
		return usuario, false
	}

	return usuario, true

}
