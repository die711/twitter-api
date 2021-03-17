package routes

import (
	"api/bd"
	"api/models"
	"encoding/json"
	"net/http"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(user.Email)
	if encontrado == true {
		http.Error(w, "Ya exist un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(user)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
