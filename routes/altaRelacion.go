package routes

import (
	"api/bd"
	"api/models"
	"net/http"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var relacion models.Relacion
	relacion.UsuarioId = IDUsuario
	relacion.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(relacion)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion"+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}
