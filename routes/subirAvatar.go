package routes

import (
	"api/bd"
	"api/models"
	"io"
	"net/http"
	"os"
	"strings"
)

func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	extension := strings.Split(handler.Filename, ".")[1]
	archivo := "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_RDONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen !"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar la imagen !"+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la BD !"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","applicacion/json")
	w.WriteHeader(http.StatusCreated)

}
