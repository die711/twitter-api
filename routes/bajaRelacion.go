package routes

import (
	"api/bd"
	"api/models"
	"net/http"
)

func BajaRelaciono(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	relacion := models.Relacion{IDUsuario, ID}

	status, err := bd.BorroRelacion(relacion)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion"+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)

}
