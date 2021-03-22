package routes

import (
	"api/bd"
	"api/models"
	"encoding/json"
	"net/http"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Datos Incorrectos"+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro, Reintente nuevamente"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modigicar el registro del usuario"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
